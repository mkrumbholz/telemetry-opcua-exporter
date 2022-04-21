FROM golang:latest

COPY . /app

WORKDIR /app

RUN go get github.com/gopcua/opcua@962060b

RUN go build -o ./opcua-exporter

CMD ["./opcua-exporter"]
