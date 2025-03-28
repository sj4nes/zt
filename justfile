fmt: 
  gofmt -w -s .

test: fmt
  DEBUG=1 go test -vet=all -tags debug -cover -v ./... -coverprofile=coverage.out

cover: test
  go tool cover -html=coverage.out -o coverage.html
  go tool cover -func=coverage.out | egrep -v -e "(100.0%|element|datum)"

reltest: fmt
  go test -vet=all  ./...

build: test
  go build -tags debug -o zt zt.go

relbuild: test
  go build -o zt zt.go

run: fmt
  DEBUG=1 go run zt.go

relrun: fmt
  go run zt.go

wat:
  watchexec -e go -r 'clear;just cover'
 
relwat:
  watchexec -e go -r 'clear;just reltest'

snaps:
 DEBUG=1 UPDATE_SNAPS=true go test ./... 

dox:
  killall godoc || true
  ~/go/bin/godoc -http=:6060 -index &
  open http://localhost:6060/pkg/zt/?m=all

