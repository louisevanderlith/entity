package core

import "github.com/louisevanderlith/husk"

type entityMap func(result interface{}, obj Entity) error

func (m entityMap) Calculate(result interface{}, obj husk.Dataer) error {
	return m(result, obj.(Entity))
}