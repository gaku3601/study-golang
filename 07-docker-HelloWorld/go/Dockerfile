FROM golang

WORKDIR /workdir
ADD $PWD/workdir/* /workdir
RUN go get github.com/ant0ine/go-json-rest/rest
CMD go run Service.go

