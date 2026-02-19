FROM golang : alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main cmd/grpcapi/server.go
EXPOSE 50051
CMD ["./server"]