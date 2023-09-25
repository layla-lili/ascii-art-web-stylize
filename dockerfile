FROM golang:1.20 

WORKDIR /app

COPY . /app

RUN go build -o bin .  

EXPOSE 8081

ENTRYPOINT ["/app/bin"]
