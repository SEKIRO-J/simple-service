package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/sekiro-j/metapierbackend/configs"
	"github.com/sekiro-j/metapierbackend/internal/app"
	"github.com/sekiro-j/metapierbackend/internal/db"

	api "github.com/sekiro-j/metapierbackend/api/protos/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

var (
	tls         = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile    = flag.String("cert_file", "../../scripts/server_cert.pem", "The TLS cert file")
	keyFile     = flag.String("key_file", "../../scripts/server_key.pem", "The TLS key file")
	gwPort      = flag.Int("gateway-port", 8080, "gateway(http reverse proxy ) port")
	gsPort      = flag.Int("grpc-server-port", 9090, "gRPC server port")
	swaggerPort = flag.Int("swagger-port", 3000, "swagger UI port")
	swaggerPath = flag.String("swagger-assets-path", "./assets/swagger", "swagger UI assests path")
)

type Render struct {
	render  *render.Render
	client  api.MetaPierServiceClient
	context context.Context
	cfg     *configs.EnvConfig
}

type binding struct {
	MainVersionHash    string
	ChunkVersionHash   string
	PriceFetchInterval string
	Env                string
	S3Path             string
}

func runSwaggerUI() error {
	mux := http.NewServeMux()
	swagger := http.FileServer(http.Dir(*swaggerPath))
	mux.Handle("/", swaggerUIHandler(swagger))
	return http.ListenAndServe(fmt.Sprintf(":%d", *swaggerPort), mux)
}

func swaggerUIHandler(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		CORS(w, req)
		fs.ServeHTTP(w, req)
	}
}

func CORS(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if req.Method == "OPTIONS" {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}
}

func (r *Render) renderHTML(w http.ResponseWriter, req *http.Request) {
	CORS(w, req)

	// default versions
	version := "1"
	priceFetchInterval := "10"
	env := "dev"

	// query for latest versions
	femd, err := r.client.GetFEMD(r.context, &api.GetFEMDRequest{})
	if err != nil {
		log.Errorf("failed to fetch latest frontend metadata: %v", err)
	} else {
		version = femd.VersionHash
		priceFetchInterval = femd.PriceFetchInterval
		env = femd.Env
	}

	log.Infof("rendering html with versionHash: %v, priceFetchInterval: %v, env: %v\n", version, priceFetchInterval, env)
	vh := strings.Split(version, "-")
	mainVH := vh[0]
	chunkVH := vh[1]
	s3bucket := "metapier-website"
	if r.cfg.Env == "production" {
		s3bucket += "-prod"
	}
	s3path := s3bucket + ".s3.amazonaws.com/static/js"

	binding := binding{MainVersionHash: mainVH, ChunkVersionHash: chunkVH, PriceFetchInterval: priceFetchInterval, Env: r.cfg.Network, S3Path: s3path}
	r.render.HTML(w, http.StatusOK, "index", binding)
}

func runReverseProxy() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api.RegisterMetaPierServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", *gsPort), opts)
	if err != nil {
		return err
	}

	// rpc client
	conn, err := grpc.Dial(fmt.Sprintf(":%d", *gsPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := api.NewMetaPierServiceClient(conn)

	envConfig, err := configs.LoadEnvConfig()
	if err != nil {
		log.Fatalf("failed to load env config: %v", err)
	}

	r := &Render{render.New(render.Options{
		Directory: "web/templates",
	}),
		c,
		ctx,
		&envConfig,
	}

	http.HandleFunc("/", r.renderHTML)

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.URL.Path, "/v1") {
			mux.ServeHTTP(w, req)
		} else {
			http.DefaultServeMux.ServeHTTP(w, req)
		}
	})

	// Start HTTP server
	log.Info("MetapierService http reverse proxy up")
	return http.ListenAndServe(fmt.Sprintf(":%d", *gwPort), handler)
}

func runServer(dbc *db.Connection) error {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *gsPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	grpcServer := grpc.NewServer(opts...)
	api.RegisterMetaPierServiceServer(grpcServer, app.New(dbc))

	log.Info("MetapierService server up")
	return grpcServer.Serve(lis)
}

func createDBConnection(cfg configs.DBConfig) (*db.Connection, error) {
	return db.New(&db.DatabaseConfig{
		Name:        cfg.Name,
		Host:        cfg.Host,
		Port:        cfg.Port,
		User:        cfg.User,
		Pwd:         cfg.Pwd,
		SSLMode:     cfg.SSLMode,
		SSLRootCert: cfg.SSLRootCert,
	})
}

func main() {

	dbConfig, err := configs.LoadDBConfig()
	if err != nil {
		log.Fatalf("failed to load db config: %v", err)
	}

	dbConnection, err := createDBConnection(dbConfig)
	if err != nil {
		log.Fatalf("failed to connect with db: %v", err)
	}

	go func() {
		if err := runSwaggerUI(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := runServer(dbConnection); err != nil {
			log.Fatal(err)
		}
	}()

	if err := runReverseProxy(); err != nil {
		log.Fatal(err)
	}
}
