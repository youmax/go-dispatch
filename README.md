# Go-Dispatch

Simple dispatcher written by GO.

The dispatcher receive job from Http POST and assign it to worker automatically.

The folder structure is similar to laravel(PHP) and was built based on Gonic-gin framework.

## Install dev package

$ cd src/upay

$ go get github.com/Masterminds/glide

$ glide install

## Build

Add project path to GOPATH then type the below command in your shell

$ go install upay

## Development server

There are many ways to run a local server. I recommand running with gin command tools which
provides hot-reloading features as file changed.The following command lines indicate the basic
usage of gin command tool.

```
$ go get github.com/codegangsta/gin

$ gin -t src/upay -b bin/upay -p 10000 -a 8080 --all run main.go

```

Alternatively, using the Golang built-in command

$ go run src/upay/main.go

Or run the binary file directly

$ ./bin/upay

## Environment setting

GIN_MODE= [release,debug], Default[debug]

DISPATCH_MODE= [development,production], Default[development]

## Testing Your Application

Test all testcases in folder

$ go test -v ./tests

Test one single file

$ go test ./test/sig_test.go 
