FROM golang:1.17

WORKDIR /app
COPY . .

ENV STATE=PROD

RUN go build

CMD ["./axie-infinity-back"]

#docker build -t moskv/axie-infinity-back:0.1.0 .
#docker run -p 8080:8080 moskv/axie-infinity-back:0.1.0 