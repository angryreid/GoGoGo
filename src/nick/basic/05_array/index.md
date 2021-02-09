# GO Array and Slice

## Defination

```go

var a[3]int;
a[0] = 1;

b := [3]int{1, 2, 3}
c := [2][2]int{{1,2}, {3,4}}

```

## Array slice data

array[startIdx(included), endIndex(excluded)]

```go
a := [...]int{1,2,3,4,5}

a[1:2] // 2
a[1:3] // 2 3
a[1:len(arr)] // 2 3 4 5
a[1:] // 2 3 4 5
a[:3] // 1 2 3

```

### Care

1. Array's inex is not support number lower then 0
2. [:] means slice all data

## Slice

Slice is a data structure of GO

## Array VS Slice

1. Array can compare, Slice can.
2. Array's capacity can't change.
