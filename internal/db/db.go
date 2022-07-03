package db

import (
	"fmt"

	"github.com/sekiro-j/metapierbackend/internal/pkg/frontend"
	"github.com/sekiro-j/metapierbackend/internal/pkg/shipyard"
	"github.com/sekiro-j/metapierbackend/internal/pkg/token"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

const delimiter string = "/"

type Connection struct {
	orm *gorm.DB
}

type DatabaseConfig struct {
	Name        string
	Host        string
	Port        int
	User        string
	Pwd         string
	SSLMode     string
	SSLRootCert string
}

func New(cfg *DatabaseConfig) (*Connection, error) {
	connection := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s sslrootcert=%s",
		cfg.Host, cfg.Port, cfg.Name, cfg.User, cfg.Pwd, cfg.SSLMode, cfg.SSLRootCert)
	log.Infof("connection string: %s", connection)

	if cfg.Host == "" {
		log.Info("database config not provided, starting without db connection")
		return nil, nil
	}

	orm, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	orm.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	orm.AutoMigrate(shipyard.Event{}, shipyard.Auction{}, shipyard.Project{}, token.Token{}, frontend.Metadata{})

	return &Connection{
		orm: orm,
	}, nil
}
