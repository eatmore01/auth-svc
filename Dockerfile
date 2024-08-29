FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/main .
COPY --from=build /app/.env .env
RUN mkdir config
COPY --from=build /app/config/config.yml ./config/config.yml

EXPOSE 9091

CMD ["./main"]