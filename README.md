# hello

## https://go.dev/doc/tutorial/call-module-code

go mod edit -replace example.com/greetings=../greetings

github.com/dirani/greetings

## on hello:
go mod edit -replace github.com/dirani/greetings=../greetings