FROM ubuntu:latest
# Add Maintainer Info
LABEL maintainer="Aditya Malpani <aditya.malpani84@gmail.com>"
RUN mkdir /app
COPY . /app
WORKDIR /app
#RUN go mod download
COPY . .
RUN echo "changes"
RUN ls -alrth
RUN go build -o main . 
CMD ["/app/main"]