BASE="assets/swagger/"
PROTO_PATH="api/protos/v1/"
NAME="simpleservice.swagger.json"
buf generate
mv "${BASE}${PROTO_PATH}${NAME}" "${BASE}${NAME}" && rm -r "${BASE}api"