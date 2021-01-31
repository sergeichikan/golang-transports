package shared

import (
	"math/rand"
	"time"
)

const Var1 = "var1"
const Var2 = "var2"
const Min = 0
const Max = 100

type Data map[string]interface{}

func (d Data) NewSumSub() Data {
	return NewSumSubResult(d[Var1].(float64), d[Var2].(float64))
}

func NewSumSubResult(a float64, b float64) Data {
	return Data{
		Var1: a + b,
		Var2: a - b,
	}
}

func NewSumSubArgs() Data {
	return NewRandomFloatData(Min, Max, Var1, Var2)
}

func NewRandomFloatData(min float64, max float64, keys ...string) Data {
	rand.Seed(time.Now().UnixNano())
	result := Data{}
	for _, key := range keys {
		result[key] = min + rand.Float64()*(max-min)
	}
	return result
}
