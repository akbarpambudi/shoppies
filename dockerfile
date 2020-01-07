FROM golang:latest
LABEL maintainer="Akbar Pambudi <akbar.pambudi@gmail.com>"
ENV http.address "8080"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]