FROM golang:1.20-alpine
RUN mkdir /app 
ADD . /app/ 
ENV ENDPOINT="http://172.217.12.110"
WORKDIR /app 
RUN go build -o main .
CMD ["/app/main"]

EXPOSE 80
