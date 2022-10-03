FROM golang:1.19.1-alpine As build_base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN cd cmd && go build -o ./out

EXPOSE 3000

CMD ["./cmd/out"]