package calc

import (
	"errors"
	"fmt"
)

type calculator struct {
}

func NewCalculator() calculator {
	return calculator{}
}
func (c calculator) add(a, b float64) float64 {
	return a + b
}
func (c calculator) sub(a, b float64) float64 {
	return a - b
}
func (c calculator) mul(a, b float64) float64 {
	return a * b
}
func (c calculator) div(a, b float64) float64 {
	return a / b
}
func (c calculator) Calculate(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return c.add(a, b), nil
	case "-":
		return c.sub(a, b), nil
	case "*":
		return c.mul(a, b), nil
	case "/":
		if b == 0 {
			return 0, errors.New("ошибка деления на ноль")
		}
		return c.div(a, b), nil
	default:
		return 0, errors.New(fmt.Sprintf("несуществующий арифметический оператор %s", operator))
	}
}
