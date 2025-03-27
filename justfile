fmt: 
  gofmt -w -s .

test: fmt
  go test -v ./...

build: test
  go build -o zt zt.go

run: fmt
  go run zt.go

