# GO data type

## Data Type

**bool**
**string**
**int** **int8** **int16** **int32** **int64**
**uint** **uint8** **uint16** **uint32** **uint64** **uintptr**
**byte** //alias for uint8
**rune** // alias for int32, represent a unicode point
**float32** **float64**
**complex64** **complex128**

## Diff with other lang

1. Go can't change type under `Implicit conversion`
2. Alias also can't change type `Implicit conversion`

## Predefined Value

1. math.MaxInt64
2. math.MaxFloat64
3. math.MaxUint32

## Pointer type

1. Go doesn't support `pointer` calc
2. string is vaule type, it's not refrece type. default value is blank string.

## nil

nil means null
