FROM golang:1.19.1-alpine
# https://zenn.dev/mirai_33/articles/1ea902d2e96b47

# WORKDIR /go/src
ENV WORK_DIR /gin1
ENV SRC_DIR ./app/cmd/gin1
# ./app/cmd/server/
# ./app
WORKDIR ${WORK_DIR}

RUN apk update && apk add git

# EXPOSE 50051
# EXPOSE 8080


COPY ${SRC_DIR} ${WORK_DIR}

RUN go mod init main \
    # && go get -u github.com/gin-gonic/gin \
    && go mod tidy \
    && go build
# RUN go get github.com/gin-gonic/gin@latest

# CMD ["go", "run", "/go/src/cmd/api/main.go"]
# CMD ["go", "run", "/app/src/cmd/chap1/main.go"]
# CMD ["go", "run", "/app/cmd/chap1/main.go"]
# CMD ["go", "run", ${WORK_DIR}"main.go"]
# CMD [".${WORK_DIR}/main"]
CMD ["./main"]
# CMD ["cat" "go.mod"]