# syntax=docker/dockerfile:1
   
FROM alpine:3.17.2

RUN apk add --no-cache git musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN git clone https://github.com/bukalawak/bukalawak-web.git bukalawak-web
WORKDIR /bukalawak-web
RUN go build -o /app-bukalawak-web
CMD ["/app-bukalawak-web"]
EXPOSE 8000