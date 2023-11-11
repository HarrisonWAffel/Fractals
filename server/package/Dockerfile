FROM golang:1.21-alpine

WORKDIR /workspace

COPY ./ ./

RUN go build -o main

EXPOSE 8989

RUN apk update
RUN apk upgrade
RUN apk add --no-cache ffmpeg

ENTRYPOINT ["./main"]