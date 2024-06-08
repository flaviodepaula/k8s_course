FROM golang:1.22
COPY . .
RUN go mod init exemple/hello
RUN go build -o server .
CMD ["./server"]