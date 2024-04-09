package cars

import "fmt"

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		// сконвертировать value в заданный в параметре UnitType
		switch t {
		case Inch:
			value /= 2.54
		case CM:
			value *= 2.54
		}
	}

	return value
}

func (u Unit) String() string {
	return fmt.Sprintf("%.2f %s", u.Value, u.T)
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type AmericanDimensions struct {
	length Unit
	width  Unit
	height Unit
}

func (a AmericanDimensions) Length() Unit {
	return Unit{
		Value: a.length.Get(Inch),
		T:     Inch,
	}
}

func (a AmericanDimensions) Width() Unit {
	return Unit{
		Value: a.width.Get(Inch),
		T:     Inch,
	}
}

func (a AmericanDimensions) Height() Unit {
	return Unit{
		Value: a.height.Get(Inch),
		T:     Inch,
	}
}

type EuropeanDimensions struct {
	length Unit
	width  Unit
	height Unit
}

func (e EuropeanDimensions) Length() Unit {
	return Unit{
		Value: e.length.Get(CM),
		T:     CM,
	}
}

func (e EuropeanDimensions) Width() Unit {
	return Unit{
		Value: e.width.Get(CM),
		T:     CM,
	}
}

func (e EuropeanDimensions) Height() Unit {
	return Unit{
		Value: e.height.Get(CM),
		T:     CM,
	}
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type auto struct {
	brand       string
	model       string
	dimensions  Dimensions
	maxSpeed    int
	enginePower int
}

//Констуруктор структуры auto принимает в качестве параметров страна происхождения (c), брэнд (b), модель (m),
//размеры в сантиметрах: длина (l), ширина (w), высота (h), макс. скорость (ms) и мощность двигателя (ep)
func NewAuto(c string, b string, m string, l float64, w float64, h float64, ms int, ep int) auto {
	var dimensions Dimensions
	if c == "USA" {
		dimensions = AmericanDimensions{
			length: Unit{Value: l, T: CM},
			width:  Unit{Value: w, T: CM},
			height: Unit{Value: h, T: CM},
		}
	} else {
		dimensions = EuropeanDimensions{
			length: Unit{Value: l, T: CM},
			width:  Unit{Value: w, T: CM},
			height: Unit{Value: h, T: CM},
		}
	}

	return auto{
		brand:       b,
		model:       m,
		dimensions:  dimensions,
		maxSpeed:    ms,
		enginePower: ep,
	}
}

func (a auto) Brand() string {
	return a.brand
}

func (a auto) Model() string {
	return a.model
}

func (a auto) Dimensions() Dimensions {
	return a.dimensions
}

func (a auto) MaxSpeed() int {
	return a.maxSpeed
}

func (a auto) EnginePower() int {
	return a.enginePower
}
