FROM golang:1.15-alpine3.12 as builder
COPY main.go .
RUN go build -o /app main.go

FROM alpine:3.12.3
ENV GOTRACEBACK=single
CMD ["./app"]
COPY --from=builder /app .
