package number

import (
	"errors"
	"math"
)

type Number interface {
	Abs() Number
	Add(any) Number
	Divide(any) (Number, error)
	GetFloat() float64
	GetInt() int64
	IsEqualTo(any) bool
	IsNegative() bool
	IsPositive() bool
	IsLessThan(any) bool
	IsLessThanOrEqualTo(any) bool
	IsGreaterThan(any) bool
	IsGreaterThanOrEqualTo(any) bool
	IsZero() bool
	Max(any) Number
	Min(any) Number
	Minus(any) Number
	Multiply(any) Number
	Round(digits ...uint) Number
	RoundDown(digits ...uint) Number
	RoundUp(digits ...uint) Number
	ShiftDecimal(digits int) Number
}

type floatNumber float64

type intNumber int64

var (
	DivideByZeroError = errors.New("can not divide by 0")
)

func Zero() Number {
	return intNumber(0)
}

func Of(num any) Number {
	switch num.(type) {
	case int8:
		return intNumber(num.(int8))
	case int16:
		return intNumber(num.(int16))
	case int:
		return intNumber(num.(int))
	case int32:
		return intNumber(num.(int32))
	case int64:
		return intNumber(num.(int64))
	case float32:
		return floatNumber(num.(float32))
	case float64:
		return floatNumber(num.(float64))
	case intNumber:
		return num.(intNumber)
	case floatNumber:
		return num.(floatNumber)
	default:
		return intNumber(0)
	}
}

func Float[T float32 | float64](num T) Number {
	return floatNumber(num)
}

func Int[T int8 | int16 | int | int32 | int64](num T) Number {
	return intNumber(num)
}

func (f floatNumber) GetFloat() float64 {
	return float64(f)
}

func (f floatNumber) GetInt() int64 {
	return int64(f)
}

func (f floatNumber) Add(another any) Number {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return f + floatNumber(num.(intNumber))
	case floatNumber:
		return f + num.(floatNumber)
	default:
		return f
	}
}

func (f floatNumber) Minus(another any) Number {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return f - floatNumber(num.(intNumber))
	case floatNumber:
		return f - num.(floatNumber)
	default:
		return f
	}
}

func (f floatNumber) Divide(another any) (Number, error) {
	num := Of(another)
	if num.IsZero() {
		return f, DivideByZeroError
	}
	switch num.(type) {
	case floatNumber:
		return f / num.(floatNumber), nil
	case intNumber:
		return f / floatNumber(num.(intNumber)), nil
	}
	return f, nil
}

func (f floatNumber) Multiply(another any) Number {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return f * floatNumber(num.(intNumber))
	case floatNumber:
		return f * num.(floatNumber)
	}
	return f
}

// Round rounds the number to the specified digits
// if no digits specified, it will round to the nearest integer
// ie: 1.2345 => 123.45 => 123 => 1.23
func (f floatNumber) Round(digits ...uint) Number {
	if len(digits) > 0 {
		digit := int(digits[0])
		return floatNumber(math.Round(f.ShiftDecimal(digit).GetFloat())).ShiftDecimal(digit * -1)
	}
	return floatNumber(math.Round(float64(f)))
}

const DefaultRoundPrecision = 2

func (f floatNumber) RoundUp(digits ...uint) Number {
	digit := DefaultRoundPrecision
	if len(digits) > 0 {
		offDigit := int(digits[0])
		return f.ShiftDecimal(offDigit).RoundUp().ShiftDecimal(offDigit * -1)
	}
	return Of(math.Ceil(Of(f.ShiftDecimal(digit).GetInt()).ShiftDecimal(digit * -1).GetFloat()))
}

// RoundDown TODO: handle 0.9999999 represents 0.1 issue
func (f floatNumber) RoundDown(digits ...uint) Number {
	digit := DefaultRoundPrecision
	if len(digits) > 0 {
		offDigit := int(digits[0])
		return f.ShiftDecimal(offDigit).RoundDown().ShiftDecimal(offDigit * -1)
	}
	return Of(math.Floor(Of(f.ShiftDecimal(digit).GetInt()).ShiftDecimal(digit * -1).GetFloat()))
}

// ShiftDecimal positive digits means shift to right, negative digits means shift to left
// ie: digits =  2 : 123.45 => 12345
// digits = -2 : 123.45 => 1.2345
func (f floatNumber) ShiftDecimal(digits int) Number {
	return floatNumber(float64(f) * math.Pow10(digits))
}

func (f floatNumber) Abs() Number {
	return floatNumber(math.Abs(float64(f)))
}

func (f floatNumber) IsZero() bool {
	return float64(f) == 0.0
}

func (f floatNumber) IsNegative() bool {
	return float64(f) < 0.0
}

func (f floatNumber) IsPositive() bool {
	return float64(f) > 0.0
}

func (f floatNumber) IsEqualTo(another any) bool {
	num := Of(another)
	switch num.(type) {
	case floatNumber:
		return f == num.(floatNumber)
	case intNumber:
		return f == floatNumber(num.(intNumber))
	}
	return false
}

func (f floatNumber) IsLessThan(another any) bool {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return f < floatNumber(num.(intNumber))
	case floatNumber:
		return f < num.(floatNumber)
	}
	return false
}

func (f floatNumber) IsLessThanOrEqualTo(another any) bool {
	return !f.IsGreaterThan(another)
}

func (f floatNumber) IsGreaterThan(another any) bool {
	num := Of(another)
	switch num.(type) {
	case floatNumber:
		return f > num.(floatNumber)
	case intNumber:
		return f > floatNumber(num.(intNumber))
	}
	return false
}

func (f floatNumber) IsGreaterThanOrEqualTo(another any) bool {
	return !f.IsLessThan(another)
}

func (f floatNumber) Max(another any) Number {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return compare(f, floatNumber(num.(intNumber)), true)
	case floatNumber:
		return compare(f, num.(floatNumber), true)
	}
	return f
}

func (f floatNumber) Min(another any) Number {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return compare(f, floatNumber(num.(intNumber)), false)
	case floatNumber:
		return compare(f, num.(floatNumber), false)
	}
	return f
}

func (i intNumber) Add(another any) Number {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return i + num.(intNumber)
	case floatNumber:
		return floatNumber(i) + num.(floatNumber)
	default:
		return i
	}
}

func (i intNumber) Minus(another any) Number {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return i - num.(intNumber)
	case floatNumber:
		return floatNumber(i) - num.(floatNumber)
	default:
		return i
	}
}

func (i intNumber) Divide(another any) (Number, error) {
	num := Of(another)
	if num.IsZero() {
		return i, DivideByZeroError
	}
	switch num.(type) {
	case floatNumber:
		return floatNumber(i) / num.(floatNumber), nil
	case intNumber:
		return floatNumber(i) / floatNumber(num.(intNumber)), nil
	}
	return i, nil
}

func (i intNumber) Multiply(another any) Number {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return i * num.(intNumber)
	case floatNumber:
		return floatNumber(i) * num.(floatNumber)
	}
	return i
}

// positive digits means shift to right, negative digits means shift to left
func (i intNumber) ShiftDecimal(digits int) Number {
	return floatNumber(float64(i) * math.Pow10(digits))
}

func (i intNumber) Abs() Number {
	if i < 0 {
		return -i
	}
	return i
}

func (i intNumber) Round(_digits ...uint) Number {
	return i
}

func (i intNumber) RoundUp(_digits ...uint) Number {
	return i
}

func (i intNumber) RoundDown(_digits ...uint) Number {
	return i
}

func (i intNumber) IsZero() bool {
	return i == 0
}

func (i intNumber) IsPositive() bool {
	return i > 0
}

func (i intNumber) IsNegative() bool {
	return i < 0
}

func (i intNumber) IsEqualTo(another any) bool {
	num := Of(another)
	switch num.(type) {
	case floatNumber:
		return floatNumber(i) == num.(floatNumber)
	case intNumber:
		return i == num.(intNumber)
	}
	return false
}

func (i intNumber) IsGreaterThan(another any) bool {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return i > num.(intNumber)
	case floatNumber:
		return floatNumber(i) > num.(floatNumber)
	}
	return false
}

func (i intNumber) IsGreaterThanOrEqualTo(another any) bool {
	return !i.IsLessThan(another)
}

func (i intNumber) IsLessThan(another any) bool {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return i < num.(intNumber)
	case floatNumber:
		return floatNumber(i) < num.(floatNumber)
	}
	return false
}

func (i intNumber) IsLessThanOrEqualTo(another any) bool {
	return !i.IsGreaterThan(another)
}

func (i intNumber) GetFloat() float64 {
	return float64(i)
}

func (i intNumber) GetInt() int64 {
	return int64(i)
}

func (i intNumber) Max(another any) Number {
	num := Of(another)
	switch num.(type) {
	case intNumber:
		return compare(i, num.(intNumber), true)
	case floatNumber:
		return compare(floatNumber(i), num.(floatNumber), true)
	}
	return i
}

func (i intNumber) Min(another any) Number {
	num := Of(another)
	switch num.(type) {
	case floatNumber:
		return compare(floatNumber(i), num.(floatNumber), false)
	case intNumber:
		return compare(i, num.(intNumber), false)
	}
	return i
}

func compare(num1 Number, num2 Number, larger bool) Number {
	if num1.IsGreaterThan(num2) {
		if larger {
			return num1
		}
		return num2
	}
	if larger {
		return num2
	}
	return num1
}

var _ Number = intNumber(0)
var _ Number = floatNumber(0.0)
