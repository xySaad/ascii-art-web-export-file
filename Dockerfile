FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o main . 

LABEL maintainer="Aymen <aymen@example.com>"
LABEL maintainer="abderahman <abderahman@example.com>"
LABEL maintainer="mohamed <mohamed@example.com>"
LABEL version="1.0"
LABEL description="This image is for ascii art web project"

EXPOSE 8080

CMD ["./main"]
