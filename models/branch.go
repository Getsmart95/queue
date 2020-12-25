package models

type Branch struct {
	ID int
	Address string		`json:"address"`
	CityID  int  	 	`json:"city_id"`
}