FROM golang:1.26-alpine as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o couik ./cmd/couik

FROM scratch 

COPY --from=builder /app/couik /usr/local/bin/couik

ENTRYPOINT [ "couik" ]
