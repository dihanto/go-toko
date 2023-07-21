FROM golang:1.20.5


WORKDIR /app

COPY . .

RUN go mod tidy

# RUN CGO_ENABLED=0 GOOS=linux go build -o /gotoko

# EXPOSE 2000

CMD ["go", "run", "main.go"]