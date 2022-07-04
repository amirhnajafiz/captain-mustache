FROM golang:1.18-alpine

WORKDIR /app

COPY * ./

RUN go build -o /docker-gs-ping main.go

EXPOSE ${PORT}

CMD [ "/docker-gs-ping" ]