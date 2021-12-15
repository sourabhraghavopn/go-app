FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./go-app  src/*.go
RUN echo $(ls src/ )

CMD ["./go-app"]