mkdir pkg
export GOPATH=$(pwd)/pkg
go get github.com/go-cmd/cmd
go build syncvmix.go
./syncvmix
