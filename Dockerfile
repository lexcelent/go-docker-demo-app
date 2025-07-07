FROM golang:1-alpine3.20

# don't do this
ENV MONGO_DB_USERNAME=admin \
    MONGO_DB_PWD=password

RUN mkdir -p /home/app

COPY ./app /home/app

WORKDIR /home/app

RUN go build ./main.go

CMD ["./main"]
