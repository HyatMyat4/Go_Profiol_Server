#Dowload Image From Docker
FROM golang:1.20

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . ./

RUN go build -o /go-docker-image

EXPOSE 4000

CMD [ "/go-docker-image" ]