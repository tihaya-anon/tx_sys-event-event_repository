FROM golang:1.24.4

WORKDIR /app

COPY /src /app

RUN cd /app/cmd && go mod tidy && go build -o ../../build/event_repository