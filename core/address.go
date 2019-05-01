package core

import "github.com/louisevanderlith/husk"

type Address struct {
	StreeNo     int
	Street      string `hsk:"size(80)"`
	UnitNo      string `hsk:"size(15)"`
	EstateName  string `hsk:"size(50)"`
	Suburb      string `hsk:"size(40)"`
	City        string `hsk:"size(40)"`
	Province    string `hsk:"size(30)"`
	PostalCode  string `hsk:"size(5)"`
	Coordinates string `hsk:"size(100)"`
	IsDelivery  bool
}

func (a Address) Valid() (bool, error) {
	return husk.ValidateStruct(&a)
}
