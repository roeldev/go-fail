go-fail
=======

[![Latest release][latest-release-img]][latest-release-url]
[![Build status][build-status-img]][build-status-url]
[![Go Report Card][report-img]][report-url]
[![Documentation][doc-img]][doc-url]

[latest-release-img]: https://img.shields.io/github/release/roeldev/go-fail.svg?label=latest
[latest-release-url]: https://github.com/roeldev/go-fail/releases
[build-status-img]: https://github.com/roeldev/go-fail/workflows/Go/badge.svg
[build-status-url]: https://github.com/roeldev/go-fail/actions?query=workflow%3AGo
[report-img]: https://goreportcard.com/badge/github.com/roeldev/go-fail
[report-url]: https://goreportcard.com/report/github.com/roeldev/go-fail
[doc-img]: https://godoc.org/github.com/roeldev/go-fail?status.svg
[doc-url]: https://pkg.go.dev/github.com/roeldev/go-fail


Fail is a Go package designed to easily create human-readable failure reports for different kinds of failed tests. It is not a testing or assertion framework.


```sh
go get github.com/roeldev/go-fail
```
```go
import "github.com/roeldev/go-fail"
```


## Show differences
Below examples creates a diff report that shows the difference between the actual and expected values.

```go
func TestAnything(t *testing.T) {
    have := Anything()
    want := "expected value"

    if have != want {
        t.Error(fail.Diff{
            Func: "Anything",
            Msg: "this is a meaningful message",
            Have: have,
            Want: want,
        })
    }
}
```
The output result is something similar like:
```text
Anything() this is a meaningful message
string(
- 	"actual value",
+ 	"expected value",
  )
```

## Unexpected error
```go
func TestDoSomething(t *testing.T) {
    if err := DoSomething(); err != nil {
        // display the error that should not occur
        t.Error(fail.Err{
            Func: "DoSomething",
            Err: err,
        })
    }
}
```
The output result is something similar like:
```text
DoSomething() unexpected error:
some error occurred while doing important stuff...
```


## Documentation
Additional detailed documentation is available at [go.dev][doc-url]


### Created with
<a href="https://www.jetbrains.com/?from=roeldev/go-fail" target="_blank"><img src="https://pbs.twimg.com/profile_images/1206615658638856192/eiS7UWLo_400x400.jpg" width="35" /></a>


## License
[GPL-3.0+](LICENSE) © 2019 [Roel Schut](https://roelschut.nl)
