# How to detect web service performance with pprof

1. import pprof package

```go
import _ "net/http/pprof"
```

2. To access `/debug/pprof` after the server started

3. You can also use `go tool pprof ` to analyze