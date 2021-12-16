FROM golang:latest

RUN mkdir /totality-corp-assignments
ADD . /totality-corp-assignments
WORKDIR /totality-corp-assignments
RUN go build -o main .
EXPOSE 8000
CMD ["/totality-corp-assignments/main"]