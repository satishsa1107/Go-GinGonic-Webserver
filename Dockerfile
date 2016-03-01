FROM satishsa1107/gin_webserver/master:latest

CMD go build server.go
CMD go run server.go

EXPOSE 8000
