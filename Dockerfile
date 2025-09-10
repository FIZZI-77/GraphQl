FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

FROM alpine:latest
WORKDIR /root/

# Копируем бинарник
COPY --from=build /app/server .

# Копируем .env из корня проекта
COPY .env ./

CMD ["./server"]
