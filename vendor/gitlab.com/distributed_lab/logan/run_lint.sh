go get -u github.com/alecthomas/gometalinter
gometalinter --install &> /dev/null
go install
gometalinter --vendor ./... | grep $1
