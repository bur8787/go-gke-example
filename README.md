# go-gke-example

## Install

```
$ brew install go
$ echo `export GOPATH=$(go env GOPATH)` >> ~/.bash_profile
$ echo `export PATH=$PATH:$(go env GOPATH)/bin` >> ~/.bash_profile
```

```
$ go get github.com/bur8787/go-gke-example
```

## Generate Code

```
$ brew install protobuf
$ sh protoc-gen-go.sh
```

## Run

```
$ cd $GOPATH/src/github.com/bur8787/go-gke-example/src
$ go run main.go
```

## Test

```
$ curl -L git.io/nodebrew | perl - setup
$ echo `export PATH=$HOME/.nodebrew/current/bin:$PATH` >> ~/.bash_profile
$ nodebrew install v8.11.2
$ nodebrew use v8.11.2
```

```
$ npm install -g grpcc
```

```
$ cd $GOPATH/src/github.com/bur8787/go-gke-example
$ grpcc -i --proto ./protofiles/product.proto --address 127.0.0.1:8080 --eval 'client.PostProduct({},printReply)'
```
