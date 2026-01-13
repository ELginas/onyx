FROM golang:1.25 AS development

RUN go install github.com/air-verse/air@latest
WORKDIR /app

CMD ["air", "-c", "air.toml"]