FROM golang

RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-gonic/contrib/static


CMD go build server.go
CMD go run server.go



EXPOSE 8000
