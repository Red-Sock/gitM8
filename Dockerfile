FROM golang as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /deploy/server/main ./cmd/gitM8/main.go

FROM scratch

WORKDIR /app
COPY --from=builder ./deploy/server/ .

EXPOSE 8080