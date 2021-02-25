test:
	go test -v ./...

cover:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out

bench:
	go test -bench=. ./...

graph:
	go mod graph | modv | dot -T png | open -f -a /System/Applications/Preview.app