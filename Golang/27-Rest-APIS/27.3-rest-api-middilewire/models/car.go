package models

import (
	"car/config"
	"database/sql"
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

func (c *Car) Get() error {
	query := `SELECT name, model, brand, year, price from cars where id=$1`
	if err := config.DB.QueryRow(query, c.Id).Scan(&c.Name, &c.Model, &c.Brand, &c.Price); err != nil {

		if err == sql.ErrNoRows {
			fmt.Printf("Error getting car: %v\n", err)
			return err
		}

	}

	return nil
}

func (c *Car) Update() error {
	query := `UPDATE cars SET name = $1, model = $2, brand=$3, year=$4, price=$5 WHERE id = $6`
	_, err := config.DB.Exec(query, c.Name, c.Model, c.Brand, c.Year, c.Price, c.Id)
	if err != nil {
		fmt.Errorf("Error updating car: %v\n", err)
		return err
	}

	return nil
}

func (c *Car) Delete() {
	query := `DELETE FROM cars WHERE id = $1`
	_, err := config.DB.Exec(query, c.Id)

	if err != nil {
		fmt.Errorf("Error while delteing the Car: %v\n", err)
	}
}
