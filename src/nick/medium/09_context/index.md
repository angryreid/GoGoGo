# Context

after go 1.9

## Detail

Root Context: using `context.Background()` to create

Child Context: using `context.WithCancel(parentContext)` to create

### Need to learn

1. all sub context will be cancelled when current context cancel.
2. to get cancel msg: `<- ctx.Done()`
