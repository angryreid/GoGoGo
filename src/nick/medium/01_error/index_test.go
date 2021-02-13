package error_test

import "testing"


var lessThanTwoError = errors.New("n should not be less than 2")

var largeThanHundredError = errors.New("n should not be larger than 100")
func getFibonacci(n int)(int[], error)  {
	if n < 2 {
		return nil, lessThanTwoError
	}

	if n > 100 {
		return nil, largeThanHundredError
	}

	fibList := []int{1,1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[n-2] + fibList[n-1])
	}

	return fibList, nil
}

func TestErrorHandle(t *testing.T) {
	t.Log(getFibonacci(3))
	if v, err := getFibonacci(5); err == nil {
		t.Log(v)
	} else {
		t.Error(err)
	}
}

func TestError(t *testing.T)  {
	t.Log("hello")
}
