# Go App

* Docker

* Gin

## Usage
```bash
git clone https://github.com/Rintarooo/go_tutorial
```

### Start local server using Docker
please see Makefile in more detail. 
```bash
# build
make build

# start local server.
make up

# check PORTS. make sure the port 3000 of localhost is open to the port 8080 of the container.
docker ps

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
docker-compose -f .devcontainer/docker-compose.local.yml build go-tutorial
# comment out ENTRYPOINT in Dockerfile and get in the container 
docker-compose -f .devcontainer/docker-compose.local.yml run --rm go-tutorial /bin/sh
# start local server
docker-compose -f .devcontainer/docker-compose.local.yml up -d

# check PORTS. make sure the port 3000 of localhost is open to the port 8080 of the container.
docker ps

curl http://localhost:3000

docker-compose -f .devcontainer/docker-compose.local.yml logs
docker-compose -f .devcontainer/docker-compose.local.yml down

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
