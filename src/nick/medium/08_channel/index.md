# Channel Close

## Close needs care

1. it will occur `panic` when close channel which is passing data
2. `v, ok := <- ch; ok` if `ok` is `true`, which means receiving data rightly, `false` means channel closed alredy
3. all channel recevier will be get an message when channel closed.