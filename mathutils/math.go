package mathutils

import (
	"math"
)

func isOddInt(x float64) bool {
	xi, xf := math.Modf(x)
	return xf == 0 && int64(xi)&1 == 1
}

// Pow returns the nth power of x

func Pow(x, y float64) float64 {
	switch {
	case y == 0 || x == 1:
		return 1
	case y == 1:
		return x
	case y == 0.5:
		return math.Sqrt(x)
	case y == -0.5:
		return 1 / math.Sqrt(x)
	case math.IsNaN(x) || math.IsNaN(y):
		return math.NaN()
	case x == 0:
		switch {
		case y < 0:
			if isOddInt(y) {
				return math.Copysign(math.Inf(1), x)
			}
			return math.Inf(1)
		case y > 0:
			if isOddInt(y) {
				return x
			}
			return 0
		}
	case math.IsInf(y, 0):
		switch {
		case x == -1:
			return 1
		case (math.Abs(x) < 1) == math.IsInf(y, 1):
			return 0
		default:
			return math.Inf(1)
		}
	case math.IsInf(x, 0):
		if math.IsInf(x, -1) {
			return Pow(1/x, -y) // Pow(-0, -y)
		}
		switch {
		case y < 0:
			return 0
		case y > 0:
			return math.Inf(1)
		}
	}

	absy := y
	flip := false
	if absy < 0 {
		absy = -absy
		flip = true
	}
	yi, yf := math.Modf(absy)
	if yf != 0 && x < 0 {
		return math.NaN()
	}
	if yi >= 1<<63 {
		return math.Exp(y * math.Log(x))
	}

	a1 := 1.0
	ae := 0

	if yf != 0 {
		if yf > 0.5 {
			yf--
			yi++
		}
		a1 = math.Exp(yf * math.Log(x))
	}

	x1, xe := math.Frexp(x)
	for i := int64(yi); i != 0; i >>= 1 {
		if i&1 == 1 {
			a1 *= x1
			ae += xe
		}
		x1 *= x1
		xe <<= 1
		if x1 < .5 {
			x1 += x1
			xe--
		}
	}

	if flip {
		a1 = 1 / a1
		ae = -ae
	}
	return math.Ldexp(a1, ae)
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
        return result, nil
    }
}
