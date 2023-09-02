# Go App

* Docker

* Gin

## Usage
please see Makefile in more detail. 
```bash
make build
make up
make curl
make down
```

if you don't use Makefile, run the following commands. 
```bash
docker-compose -f .devcontainer/docker-compose.yml build go-tutorial
# docker-compose -f .devcontainer/docker-compose.yml run --rm go-tutorial /bin/sh
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
