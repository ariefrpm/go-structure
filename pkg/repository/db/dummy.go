package db

import (
	"log"
)

type Dummy interface {
	GetSomething() Some
}

type dummy struct {

}

func NewDummyRepo() Dummy {
	return &dummy{}
}

func (d *dummy) GetSomething() Some {
	log.Println("getsomethign")
	return Some{Name:"test"}
}
