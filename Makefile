run:
	go run main/hw.go

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
	GOOS=linux go build -o ./builds/linuxFile main/hw.go
	GOOS=windows GOARCH=amd64 go build -o ./builds/win.exe main/hw.go

graph:
	go mod graph | modv | dot -T png | open -f -a /System/Applications/Preview.app

