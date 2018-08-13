package model

//Galinha é um metodo que precisa ser comentado
type Galinha interface {
	Cacareja() string
}

//Pata é um metodo que precisa ser comentado
type Pata interface {
	Quack() string
}

//Ave é um metodo que precisa ser comentado
type Ave struct {
	Nome string
}

//Cacareja é um metodo que precisa ser comentado
func (a Ave) Cacareja() string {
	return "Cócóricó..."
}

//Quack é um metodo que precisa ser comentado
func (a Ave) Quack() string {
	return "Quack, Quack..."
}
