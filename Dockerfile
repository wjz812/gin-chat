FROM golang:latest
ENV VERSION 1.0
RUN mkdir -p /www/webapp
WORKDIR /www/webapp
COPY . /www/webapp
RUN go build main.go
EXPOSE 8081
RUN chmod +x main
ENTRYPOINT ["./main"]
