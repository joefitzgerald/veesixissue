Reproduction steps:

1. Run `GO111MODULE=off go get -u github.com/myitcv/gobin`
1. Run `go generate` in the same directory as `myinterface.go`
1. Observe the following error:

```shell
$ go generate
failed to run /Users/joe/projects/veesixissue/.gobincache/github.com/maxbrunsfeld/counterfeiter/v6/@v/v6.0.2/github.com/maxbrunsfeld/counterfeiter/v6/v6: fork/exec /Users/joe/projects/veesixissue/.gobincache/github.com/maxbrunsfeld/counterfeiter/v6/@v/v6.0.2/github.com/maxbrunsfeld/counterfeiter/v6/v6: no such file or directory
myinterface.go:3: running "gobin": exit status 1
```

I have tried changing the `go generate` directive in `myinterface.go` to be:

```go
//go:generate gobin -m -run github.com/maxbrunsfeld/counterfeiter . MySpecialInterface`
```

This works, but it causes an extraneous dependency to be registered in `go.mod`:

```
github.com/maxbrunsfeld/counterfeiter v0.0.0-20190111051351-d7285d544cff // indirect
github.com/maxbrunsfeld/counterfeiter/v6 v6.0.2
```

This makes me believe that either `gobin` is failing to correctly use the tagged `v6.0.2` dependency, or I've messed something up when releasing the `v6.0.2` `counterfeiter` dependency. Which is it?
