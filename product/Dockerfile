FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o micro-product main.go

FROM alpine
WORKDIR /app
COPY .env .
COPY goql .
COPY db db
COPY --from=builder /app/micro-product /app/micro-product
EXPOSE 9001
CMD ["/app/micro-product"]