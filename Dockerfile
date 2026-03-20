FROM golang:latest as build

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN GOOS="linux" GOARCH="amd64" CGO_ENABLED=0 go build -o mappsProduct ./cmd/main/main.go

FROM --platform=linux/amd64 alpine

COPY --from=build /app/mappsProduct /app/mappsProduct

WORKDIR /app

EXPOSE 8082

CMD ["/app/mappsProduct"]
