# Build
FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY /src/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /main


# Run
FROM alpine
WORKDIR /opt
COPY --from=builder /main .

CMD [ "./main" ]
