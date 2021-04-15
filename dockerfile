FROM golang:1.16.3

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 5678

CMD ["go", "run", "main.go"]