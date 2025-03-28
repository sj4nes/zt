fmt: 
  gofmt -w -s .

test: fmt
  go test -vet=all -tags debug -v ./...

reltest: fmt
  go test -vet=all  ./...

build: test
  go build -tags debug -o zt zt.go

relbuild: test
  go build -o zt zt.go

run: fmt
  go run zt.go

wat:
  watchexec -e go -r just test
