
FROM golang:1.14.9-stretch

WORKDIR /app
COPY . /app

EXPOSE 3000

CMD [ "go", "run", "main.go" ]
