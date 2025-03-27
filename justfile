fmt: 
  gofmt -w -s .

run: fmt
  go run zt.go

