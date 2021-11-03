# COMPILING
FROM golang:1.16-alpine3.14 AS compiler 

ADD . /go/src/server

WORKDIR /go/src/server

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://goproxy.cn 
RUN go mod tidy 
RUN go build -o build/server github.com/JEDIAC/server

# RUNNING 
FROM alpine:3.14 
COPY --from=compiler /go/src/server/build/server .
COPY --from=compiler /go/src/server/setting/*.yml ./setting/
ENTRYPOINT [ "./server" ]