package number

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NumberTestSuite struct {
	suite.Suite
}

func (suite *NumberTestSuite) TestNumber_Add() {
	suite.Run("1 + 1 is 2", func() {
		suite.Equal(int64(2), Int(1).Add(Int(1)).GetInt(), "should be 2")
		suite.Equal(int64(2), Of(1).Add(1).GetInt(), "should be 2")
	})

	suite.Run("1 + 1.1 is 2.1", func() {
		suite.Equal(float64(2.1), Int(1).Add(Float(1.1)).GetFloat(), "should be 2.1")
		suite.Equal(float64(2.1), Of(1).Add(1.1).GetFloat(), "should be 2.1")
	})
}

func (suite *NumberTestSuite) TestNumber_Minue() {
	suite.Run("1 - 2 is -1", func() {
		suite.Equal(int64(-1), Int(1).Minus(Int(2)).GetInt(), "should be -1")
		suite.Equal(int64(-1), Of(1).Minus(2).GetInt(), "should be -1")
	})

	suite.Run("1 - 2.0 is -1.0", func() {
		suite.Equal(float64(-1), Int(1).Minus(Float(2.0)).GetFloat(), "should be -1")
		suite.Equal(float64(-1.0), Of(1).Minus(2.0).GetFloat(), "should be -1.0")
	})
}

func (suite *NumberTestSuite) TestNumber_Divide() {
	suite.Run("1 / 2 is 0.5 in float", func() {
		result, err := Int(1).Divide(Float(2.0))
		suite.Nil(err, "should not have error")
		suite.Equal(float64(0.5), result.GetFloat(), "should be 0.5")

		result, err = Of(1).Divide(2.0)
		suite.Nil(err, "should not have error")
		suite.Equal(float64(0.5), result.GetFloat(), "should be 0.5")
	})

	suite.Run("1 / 0 is illegal", func() {
		result, err := Int(1).Divide(Int(0))
		suite.Equal(err, errors.New("can not divide by 0"))
		suite.Equal(float64(1), result.GetFloat(), "should not be 1")

		result, err = Int(1).Divide(0)
		suite.Equal(err, errors.New("can not divide by 0"))
		suite.Equal(float64(1), result.GetFloat(), "should not be 1")
	})

	suite.Run("1 / 0.0 is illegal", func() {
		result, err := Int(1).Divide(Float(0.0))
		suite.Equal(err, errors.New("can not divide by 0"))
		suite.Equal(float64(1), result.GetFloat(), "should not be 1")

		result, err = Of(1).Divide(0.0)
		suite.Equal(err, errors.New("can not divide by 0"))
		suite.Equal(float64(1), result.GetFloat(), "should not be 1")
	})
}

func (suite *NumberTestSuite) TestNumber_Multiply() {
	suite.Run("1 * 2 is 2", func() {
		suite.Equal(int64(2), Int(1).Multiply(Float(2.0)).GetInt(), "should be 2")
		suite.Equal(int64(2), Of(1).Multiply(2).GetInt(), "should be 2")
	})

	suite.Run("1 * -2.0 is -2.0", func() {
		suite.Equal(int64(-2), Int(1).Multiply(Float(-2.0)).GetInt(), "should be -2.0")
		suite.Equal(int64(2), Of(1).Multiply(2.0).GetInt(), "should be 2")
	})

	suite.Run("1 * -0.0 is 0", func() {
		suite.Equal(int64(0), Int(1).Multiply(Float(-0.0)).GetInt(), "should be 0")
		suite.Equal(int64(0), Int(1).Multiply(-0.0).GetInt(), "should be 0")
	})

}

func (suite *NumberTestSuite) TestNumber_ShiftDecimal() {
	suite.Run("1 should be 0.01 if decimal by -2 digits", func() {
		suite.Equal(float64(0.01), Int(1).ShiftDecimal(-2).GetFloat())
	})

	suite.Run("1 should be 0.0001 if decimal by -4 digits", func() {
		suite.Equal(float64(0.0001), Int(1).ShiftDecimal(-4).GetFloat())
	})

	suite.Run("0.01 should be 1.0 if decimal by 2 digits", func() {
		suite.Equal(float64(1.0), Float(0.01).ShiftDecimal(2).GetFloat())
	})

	suite.Run("0.01 should be 1.0 if shift decimal by 2 digits", func() {
		suite.Equal(float64(1.0), Float(0.01).ShiftDecimal(2).GetFloat())
	})

	suite.Run("1.0 should be 0.0001 if decimal by -4 digits", func() {
		suite.Equal(float64(0.0001), Float(1.0).ShiftDecimal(-4).GetFloat())
	})
}

func (suite *NumberTestSuite) TestNumber_Round() {
	suite.Run("should round up if it is 0.1", func() {
		suite.Equal(float64(0), Float(0.1).Round().GetFloat())
	})

	suite.Run("should round up if 0.5", func() {
		suite.Equal(float64(1), Float(0.5).Round().GetFloat())
	})
}

func (suite *NumberTestSuite) TestNumber_RoundUp() {
	suite.Run("should round up if it is 0.1", func() {
		suite.Equal(float64(1), Float(0.1).RoundUp().GetFloat())
	})

	suite.Run("should round up if 0.5", func() {
		suite.Equal(float64(1), Float(0.5).RoundUp().GetFloat())
	})

	suite.Run("should round up to given decimal", func() {
		suite.Equal(float64(0.67), Float(0.664).RoundUp(2).GetFloat())
	})

	suite.Run("should round to one place if given digits value is 0", func() {
		suite.Equal(float64(1.0), Float(0.664).Round(0).GetFloat())
	})

	suite.Run("should handle float number precision issue properly", func() {
		suite.Equal(float64(15), Float(float32(0.1)).Multiply(150).RoundUp().GetFloat())
	})
}

func (suite *NumberTestSuite) TestNumber_RoundDown() {
	suite.Run("should round down if it is 0.1", func() {
		suite.Equal(float64(0), Float(0.1).RoundDown().GetFloat())
	})

	suite.Run("should round down if 0.5", func() {
		suite.Equal(float64(0), Float(0.5).RoundDown().GetFloat())
	})

	suite.Run("should round down to given decimal", func() {
		suite.Equal(float64(0.66), Float(0.668).RoundDown(2).GetFloat())
	})
}

func (suite *NumberTestSuite) TestNumber_IsZero() {
	suite.Run("should be true if it is 0", func() {
		suite.True(Of(0).IsZero())
		suite.True(Of(0).Add(1).Minus(1).IsZero())
	})

	suite.Run("should be true if it is 0.0", func() {
		suite.True(Of(0.0).IsZero())
		suite.True(Of(0).Add(1.0).Minus(1.0).IsZero())
	})
}

func (suite *NumberTestSuite) TestNumber_IsPositive() {
	suite.Run("should be true if it is 1", func() {
		suite.True(Of(1).IsPositive())
		suite.True(Of(-1).Multiply(-1).IsPositive())
	})

	suite.Run("should be true if it is 1.0", func() {
		suite.True(Of(1.0).IsPositive())
		suite.True(Of(-1.0).Multiply(-1.0).IsPositive())
	})

	suite.Run("should be false if it is -1.0", func() {
		suite.False(Of(-1.0).IsPositive())
	})
}

func (suite *NumberTestSuite) TestNumber_IsGreaterThan() {
	suite.Run("2 is greater than 1", func() {
		suite.True(Of(2).IsGreaterThan(Of(1)))
		suite.True(Of(2).IsGreaterThan(1))
	})

	suite.Run("2.5 is greater than 1", func() {
		suite.True(Of(2.5).IsGreaterThan(Of(1)))
		suite.True(Of(2.5).IsGreaterThan(1.0))
	})
}

func (suite *NumberTestSuite) TestNumber_IsLessThan() {
	suite.Run("1 is less than 2", func() {
		suite.True(Of(1).IsLessThan(Of(2)))
		suite.True(Of(1).IsLessThan(2))
	})

	suite.Run("1.0 is less than 2.5", func() {
		suite.True(Of(1.0).IsLessThan(Of(2.5)))
		suite.True(Of(1).IsLessThan(2.5))
	})
}

func (suite *NumberTestSuite) TestNumber_IsLessThanEqualTo() {
	suite.Run("1 is less than or equal to 2", func() {
		suite.True(Of(1).IsLessThanOrEqualTo(Of(2)))
		suite.True(Of(1).IsLessThanOrEqualTo(2))
	})

	suite.Run("1.0 is less than or equal to 2.5", func() {
		suite.True(Of(1.0).IsLessThanOrEqualTo(Of(2.5)))
		suite.True(Of(1).IsLessThanOrEqualTo(2.5))
	})

	suite.Run("1.0 is less than or equal to 1.0", func() {
		suite.True(Of(1.0).IsLessThanOrEqualTo(Of(1.0)))
		suite.True(Of(1).IsLessThanOrEqualTo(1.0))
	})
}

func (suite *NumberTestSuite) TestNumber_Max() {
	suite.Run("2 is greater than 1", func() {
		suite.Equal(intNumber(2), Of(2).Max(Of(1)))
		suite.Equal(intNumber(2), Of(2).Max(1))
	})

	suite.Run("2.5 is greater than 1", func() {
		suite.Equal(floatNumber(2.5), Of(2.5).Max(Of(1)))
		suite.Equal(floatNumber(2.5), Of(2.5).Max(1))
	})
}

func (suite *NumberTestSuite) TestNumber_Min() {
	suite.Run("2 is greater than 1", func() {
		suite.Equal(intNumber(1), Of(2).Min(Of(1)))
		suite.Equal(intNumber(1), Of(2).Min(1))
	})

	suite.Run("2.5 is greater than 1.1", func() {
		suite.Equal(floatNumber(1.1), Of(2.5).Min(Of(1.1)))
		suite.Equal(floatNumber(1.1), Of(2.5).Min(1.1))
	})
}

func (suite *NumberTestSuite) TestNumber_IsNegative() {
	suite.Run("should be true if it is -1", func() {
		suite.True(Of(-1).IsNegative())
	})

	suite.Run("should be true if it is -1.0", func() {
		suite.True(Of(-1.0).IsNegative())
	})

	suite.Run("should be false if it is 1.0", func() {
		suite.False(Of(1.0).IsNegative(), "1.0 is not negative")
	})
}

func (suite *NumberTestSuite) TestNumber_Abs() {
	suite.Run("should be 1 if it is -1", func() {
		suite.Equal(int64(1), Of(-1).Abs().GetInt(), "should be 1")
	})

	suite.Run("should be 1.0 if it is -1.0", func() {
		suite.Equal(int64(1), Of(-1.0).Abs().GetInt(), "should be 1.0")
	})
}

func (suite *NumberTestSuite) TestNumber_GetInt() {
	suite.Run("should be 1 if it is 1", func() {
		suite.Equal(int64(1), Of(1).GetInt(), "should be 1")
	})

	suite.Run("should be -1 if it is -1.0", func() {
		suite.Equal(int64(-1), Of(-1.0).GetInt(), "should be 1")
	})
}

func (suite *NumberTestSuite) TestNumber_GetFloat() {
	suite.Run("should be 1.0 if it is 1", func() {
		suite.Equal(1.0, Of(1).GetFloat(), "should be 1")
	})

	suite.Run("should be -1.0 if it is -1.0", func() {
		suite.Equal(-1.0, Of(-1.0).GetFloat(), "should be -1.0")
	})
}

func (suite *NumberTestSuite) TestNumber_Integrated() {
	suite.Run("(3 + 5) / 2 == 4", func() {
		result, err := Of(3).Add(Of(5)).Divide(Of(2))
		suite.Nil(err, "should have no error")
		suite.Equal(int64(4), result.GetInt())
	})

	suite.Run("3 + 5 * 2 == 13", func() {
		result := Of(3).Add(Of(5).Multiply(Of(2)))
		suite.Equal(int64(13), result.GetInt())
	})

	suite.Run("100 * 0.03 == 3.0", func() {
		result := Of(100).Multiply(Of(3).ShiftDecimal(2))
		suite.Equal(float64(30000), result.GetFloat())
	})
}

func TestNumberTestSuite(t *testing.T) {
	suite.Run(t, new(NumberTestSuite))
}
