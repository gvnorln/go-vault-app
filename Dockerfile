# Stage 1: Build
FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Stage 2: Minimal runtime
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
RUN chmod +x ./app
EXPOSE 8080
CMD ["./app"]
