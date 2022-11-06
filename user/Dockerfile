FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o micro-user main.go


FROM alpine
WORKDIR /app
COPY .env .
COPY goql .
COPY db db
COPY --from=builder /app/micro-user /app/micro-user
EXPOSE 9000

CMD ["/app/micro-user"]
