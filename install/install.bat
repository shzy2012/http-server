

#env GOOS=windows GOARCH=amd64 go build -mod=vendor -o http-server.exe main.go

#env GOOS=darwin GOARCH=amd64 go build -mod=vendor -o http-server-mac main.go