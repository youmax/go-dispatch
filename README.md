# Go-Dispatch

Simple dispatcher written by GO.

The dispatcher receive job from Http POST and assign it to worker automatically.

## install dev package

cd src/upay

go get github.com/tools/godep

godep restore

## build

go install upay

## Testing Your Application

$ go test -v ./tests
