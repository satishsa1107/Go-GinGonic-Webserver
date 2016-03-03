FROM golang

ADD . /go/src/sagar.satish2/go_web_server
ADD webfiles /go/src/sagar.satish2/go_web_server/webfiles

RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-gonic/contrib/static
RUN go install sagar.satish2/go_web_server

ADD webfiles /go/bin/go_web_server/webfiles

ENTRYPOINT /go/bin/go_web_server

EXPOSE 8000
