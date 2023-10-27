FROM golang:1.19 as builder
ENV WORK_DIR /gin9
ENV LOCAL_DIR ./app/cmd${WORK_DIR}
WORKDIR ${WORK_DIR}
COPY ${LOCAL_DIR}/go.mod \
     ${LOCAL_DIR}/go.sum \
     ${LOCAL_DIR}/main.go .
RUN CGO_ENABLED=0 go build -o /bin/main main.go
# ENTRYPOINT ["/bin/main"]

FROM alpine:latest
ENV WORK_DIR /gin9
ENV LOCAL_DIR ./app/cmd${WORK_DIR}
WORKDIR ${WORK_DIR}
COPY ${LOCAL_DIR} .
COPY --from=builder /bin/main .
ENTRYPOINT ["./main"]

# https://zenn.dev/mirai_33/articles/1ea902d2e96b47
# https://mixi-developers.mixi.co.jp/go-remote-debug-5c8dda2d7064