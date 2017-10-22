package mathutils

import "errors"

// Pow returns the nth power of x if n is non-negative
// Otherwise the return value is undefined
func Pow(x int, n int) (int, error) {
	if n < 0 {
		return 0, errors.New("bad input")
	} else if n == 0 {
		return 1, nil
	} else {
		y, err := Pow(x, n-1)
		return x * y, err
	}
}

func Pow_Imperative(x int, n int)(int, error) {
    if n < 0 {
        return 0, errors.New("bad input")
    } else if n == 0 {
        return 1, nil
    } else {
        result := 1
        for n >= 1 {
            result*=x
            n-=1
            
        }
        return result
}
	
