FROM golang:1.19

WORKDIR /app

COPY . ./src

RUN go mod init mathematics

#RUN go build -o math

#CMD ["./math"]