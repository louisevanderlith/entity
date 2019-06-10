package core

import "github.com/louisevanderlith/husk"

type Entity struct {
	Name         string   `hsk:"size(30)"`
	ProfileKey   husk.Key `hsk:"null"`
	Registration string   `hsk:"size(30)"` //This can be a Company Registration, RSA ID, Passport or another interal identifier
	//Addresses  []Address: Future development
	Contact Contact
}

func (e Entity) Valid() (bool, error) {
	return husk.ValidateStruct(&e)
}

func GetEntities(page, pagesize int) husk.Collection {
	return ctx.Entities.Find(page, pagesize, husk.Everything())
}

func GetEntity(key husk.Key) (husk.Recorder, error) {
	return ctx.Entities.FindByKey(key)
}

func (e Entity) Create() (husk.Recorder, error) {
	rec := ctx.Entities.Create(e)

	if rec.Error != nil {
		return nil, rec.Error
	}

	ctx.Entities.Save()

	return rec.Record, nil
}
