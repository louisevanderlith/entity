package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type entityMap func(result interface{}, obj Entity) error

func (m entityMap) Calculate(result interface{}, obj hsk.Record) error {
	return m(result, obj.GetValue().(Entity))
}
