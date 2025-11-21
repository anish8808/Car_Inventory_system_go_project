package models

import (
	"car/config"
	"fmt"
)

type Car struct {
	Id    int
	Name  string
	Model string
	Brand string
	Year  int
	Price float64
}

var Cars = make(map[int]Car)

func (c *Car) Insert() {
	query := `INSERT INTO cars (name , model , brand , year , price) VALUES ($1 , $2 , $3 , $4 , $5) RETURNING id`
	if err := config.DB.QueryRow(query, c.Name, c.Model, c.Brand, c.Year, c.Price).Scan(&c.Id); err != nil {
		fmt.Errorf("Error insterting car: %v\n", err)
	}

}
