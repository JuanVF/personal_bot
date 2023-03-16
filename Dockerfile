FROM golang:1.20

WORKDIR /personal_bot

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

ENV ENVIRONMENT container

CMD ["./main"]