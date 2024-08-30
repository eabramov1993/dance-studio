FROM golang:1.23-alpine 

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . . 

# Переходим в директорию cmd, где находится main.go
WORKDIR /app/cmd

# Собираем проект в директории cmd
RUN go build -o /app/main . 

# Возвращаемся в корневую директорию /app
WORKDIR /app

RUN chmod +x main 

EXPOSE 8080

CMD ["./main"]