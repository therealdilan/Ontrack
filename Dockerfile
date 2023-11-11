FROM golang:1.16-alpine

WORKDIR /gotrack

COPY go.mod .
COPY main.go .
COPY index.html .

RUN go build -o bin . 

EXPOSE 8080

ENTRYPOINT [ "/gotrack/bin" ]
