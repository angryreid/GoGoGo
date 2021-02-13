# Error

## Spec with go error

1. no error system
2. error implement the error interface
3. go can use `error.New` to create a error

```go
type error interface {
  Error() string
}

errors.New("...")

```
