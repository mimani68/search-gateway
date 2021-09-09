
FROM golang:1.14.9-stretch

WORKDIR /app
COPY . /app

EXPOSE 3000

CMD [ "flask", "run", "--host=0.0.0.0", "--port=3000" ]
