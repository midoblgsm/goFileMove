FROM golang:1.9.1
WORKDIR /go/src/github.com/midoblgsm/goFileMove/
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -a --ldflags '-w -linkmode external -extldflags "-static"'  -o app main.go


FROM alpine:latest
RUN apk --no-cache add ca-certificates=20161130-r2 openssl=1.0.2k-r0
WORKDIR /root/
COPY --from=0 /go/src/github.com/midoblgsm/goFileMove/app .

CMD ["./app"]