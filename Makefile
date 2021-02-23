test:
	go test -v ./...

cover:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out

bench:
	go test -bench=. ./...

doc:
	godoc -http=:6060

build:
	GOOS=linux go build -o ./builds/linuxFile main/hw1.go
	GOOS=windows GOARCH=amd64 go build -o ./builds/win.exe main/hw1.go