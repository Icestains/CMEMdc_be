FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/icestains/CMEMdc_be
COPY . $GOPATH/src/github.com/icestains/CMEMdc_be
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./CMEMdc_be"]