package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string
}

type applePhone struct {
	model string
	os    string
}

// Конструктор структуры applePhone принимает в качестве параметров наименование модели (m)
// и название операционной системы (o), возвращает экземпляр структуры с инициализированными полями
func NewApplePhone(m string, o string) applePhone {
	return applePhone{
		model: m,
		os:    o,
	}
}

func (a applePhone) Brand() string {
	return "Apple"
}

func (a applePhone) Model() string {
	return a.model
}

func (a applePhone) Type() string {
	return "smartphone"
}

func (a applePhone) OS() string {
	return a.os
}

/*
func (a applePhone) String() string {
	return fmt.Sprintf("Brand: %s\nModel: %s\nType: %s\nOS: %s\n", a.Brand(), a.Model(), a.Type(), a.OS())
}*/

type androidPhone struct {
	brand string
	model string
	os    string
}

// Конструктор структуры androidPhone принимает в качестве параметров название брэнда (b), наименование модели (m)
// и название операционной системы (o), возвращает экземпляр структуры с инициализированными полями
func NewAndroidPhone(b string, m string, o string) androidPhone {
	return androidPhone{
		brand: b,
		model: m,
		os:    o,
	}
}

func (a androidPhone) Brand() string {
	return a.brand
}

func (a androidPhone) Model() string {
	return a.model
}

func (a androidPhone) Type() string {
	return "smartphone"
}

func (a androidPhone) OS() string {
	return a.os
}

/*
	func (a androidPhone) String() string {
		return fmt.Sprintf("Brand: %s\nModel: %s\nType: %s\nOS: %s\n", a.Brand(), a.Model(), a.Type(), a.OS())
	}
*/
type radioPhone struct {
	brand   string
	model   string
	buttons int
}

// Конструктор структуры radioPhone принимает в качестве параметров название брэнда (b), наименование модели (m)
// и количество кнопок на аппарате (but), возвращает экземпляр структуры с инициализированными полями
func NewRadioPhone(b string, m string, but int) radioPhone {
	return radioPhone{
		brand:   b,
		model:   m,
		buttons: but,
	}
}

func (r radioPhone) Brand() string {
	return r.brand
}

func (r radioPhone) Model() string {
	return r.model
}

func (r radioPhone) Type() string {
	return "station"
}

func (r radioPhone) ButtonsCount() int {
	return r.buttons
}

/*
func (a radioPhone) String() string {
	return fmt.Sprintf("Brand: %s\nModel: %s\nType: %s\nButtons: %d\n", a.Brand(), a.Model(), a.Type(), a.ButtonsCount())
}*/
