FROM golang:alpine
RUN apk add build-base
WORKDIR /app
COPY . .

RUN GO111MODULE=on go get github.com/cucumber/godog/cmd/godog@v0.11.0
RUN go build -o /app/fibo main.go
RUN chmod +x /app/run_tests.sh
CMD ["/app/fibo"]