FROM golang as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /deploy/server/gitm8 ./cmd/gitM8/main.go

FROM scratch

LABEL com.centurylinklabs.watchtower.enable=true

WORKDIR /app
COPY --from=builder ./deploy/server/ .
COPY --from=builder ./config/config.yaml ./config/config.yaml
EXPOSE 8080

ENTRYPOINT ["./gitm8"]