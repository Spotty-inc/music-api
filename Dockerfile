FROM golang:1.16-alpine

WORKDIR /app
COPY /src ./

RUN go mod download

RUN go build -o /main

EXPOSE 10000

CMD [ "/main" ]
