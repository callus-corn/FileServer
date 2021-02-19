FROM golang:1.16.0-alpine3.13

WORKDIR /back
COPY . .

RUN go mod tidy
RUN go install

CMD ["back"]
