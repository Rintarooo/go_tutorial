# Go App

* Docker

* Gin

## Usage
### Docker
please see Makefile in more detail. 
```bash
# build
make build

# start local server. port 3000
make up

# curl
make curl

# debug
make logs

# down local server.
make down
```

if you don't use Makefile, run the following commands. 
```bash
# build
docker-compose -f .devcontainer/docker-compose.yml build go-tutorial
# comment out ENTRYPOINT in Dockerfile and get in the container 
docker-compose -f .devcontainer/docker-compose.yml run --rm go-tutorial /bin/sh
# start local server
docker-compose -f .devcontainer/docker-compose.yml up -d

# localhostの3000がコンテナの8080のポートに、PORTSが開かれていることを確認
docker ps

curl http://localhost:3000

docker-compose -f .devcontainer/docker-compose.yml logs
docker-compose -f .devcontainer/docker-compose.yml down

# curl -X 'POST' 'http://localhost:${PORT}/predict' \
# 	    -H 'accept: application/json' \
# 	    -H 'Content-Type: application/json' \
# 	    -d '{"sample1": "f39259", \
# 	         "sample2": "0"}'
```

### Local on M1 Mac
```bash
brew install go
go version

cd ./app/cmd/gin6/
go mod init main
go mod tidy
go run main.go
```

GOPATH was deprecated in Go 1.10. Go 1.11 introduced module, enabling an alternative workflow.
