FROM docker.io/satishsa1107/gin_webserver

CMD go build server.go
CMD go run server.go

EXPOSE 8000
