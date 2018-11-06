FROM golang:1.11.2-alpine3.8

ENV CPU_LOAD_TO 10

COPY main.go /app/

WORKDIR /app

RUN go build -o cpuburn \
    && rm main.go

ENTRYPOINT /app/cpuburn
