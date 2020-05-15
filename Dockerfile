FROM golang:latest 

RUN mkdir /usr/src/app
ADD . /usr/src/app/ 
WORKDIR /usr/src/app
RUN go build -o main . 

CMD ["/app/main"]