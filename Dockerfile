FROM golang:1.24.1-alpine3.21

COPY . /app
WORKDIR /app
EXPOSE 9000

CMD [ "go","run","main.go" ]