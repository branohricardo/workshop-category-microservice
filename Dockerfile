FROM golang:latest

ADD . /go/src/github.com/branohricardo/workshop-category-microservice

RUN go install github.com/branohricardo/workshop-category-microservice

ENTRYPOINT /go/bin/workshop-category-microservice

EXPOSE 8080