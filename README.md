# Call Me Maybe

This repo contains code that was used to generate benchmarks for the Call Me Maybe talk.

Create Humans REST:

```bash
goos: darwin
goarch: arm64
pkg: CallMeMaybe/rest
BenchmarkCreateHumans-10            9297            189943 ns/op
PASS
ok      CallMeMaybe/rest        3.158s
```

Get Humans REST:

```bash
goos: darwin
goarch: arm64
pkg: CallMeMaybe/rest
BenchmarkGetHumans-10               9762            169877 ns/op
PASS
ok      CallMeMaybe/rest        2.803s
```
