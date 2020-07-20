FROM golang:alpine

LABEL maintainer="Aditya Malpani <aditya.malpani@gmail.com>"
RUN mkdir /app
COPY myprogram.go /app
WORKDIR /app
RUN go build -o main .
RUN ls  -alrth 
CMD ["/app/main"]
#RUN go run .