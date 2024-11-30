FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM scratch AS prod

COPY --from=builder /app/main /main
COPY --from=builder /app/internal/search/data/data.json /internal/search/data/data.json

CMD [ "/main" ]