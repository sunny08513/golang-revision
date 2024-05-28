# errors handling in golang

The error interface is built into the language.

```
http://golang.org/pkg/builtin/#error
type error interface {
    Error() string
}
```

This is why it appears to be an unexported identifier. Any concrete value that implements this interface can be used as an error value.

# errors package in golang

```
// http://golang.org/src/pkg/errors/errors.go
type errorString struct {
    s string
}

// http://golang.org/src/pkg/errors/errors.go
func (e *errorString) Error() string {
    return e.s
}
```
