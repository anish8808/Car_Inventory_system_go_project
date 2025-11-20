package models

type Car struct {
	Id    int
	Name  string
	Model string
	Brand string
	Year  int
	Price float64
}

var Cars = make(map[int]Car)
