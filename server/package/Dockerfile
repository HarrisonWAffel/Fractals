FROM golang

WORKDIR /workspace

COPY ./ ./

RUN go build -o main

EXPOSE 8989

ENTRYPOINT ["./main"]