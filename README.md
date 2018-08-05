# Go-Dispatch

Simple dispatcher written by GO.

The dispatcher receive job from Http POST and assign it to worker automatically.

## Install dev package

go get github.com/joho/godotenv

cd src/upay

go get github.com/Masterminds/glide

glide install

## build

Add project path to GOPATH 

go install upay

## Development server

go get github.com/codegangsta/gin

gin -t src/upay -b bin/upay -p 10000 -a 8080 --all run main.go

## Testing Your Application

$ go test -v ./tests
