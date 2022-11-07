FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app/calendar ./cmd/calendar

FROM scratch

COPY --from=builder /app .

EXPOSE 8080
CMD ["./calendar"]