package core

import (
	"errors"
	"fmt"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/kong/prime"
	"github.com/louisevanderlith/kong/stores"
	"log"
)

type EntityContext interface {
	stores.UserStore
	GetEntities(page, pageSize int) (husk.Collection, error)
	GetEntity(k husk.Key) (husk.Recorder, error)
	Register(obj Registration) error
	RequestReset(email, host string) (string, error)
}

type context struct {
	Entities  husk.Tabler
	Forgotten husk.Tabler
}

var ctx context

func Context() EntityContext {
	return ctx
}

func CreateContext() {
	ctx = context{
		Entities: husk.NewTable(Entity{}),
	}
	seed()
}

func seed() {
	err := ctx.Entities.Seed("db/enities.seed.json")

	if err != nil {
		panic(err)
	}

	err = ctx.Entities.Save()

	if err != nil {
		panic(err)
	}
}

func Shutdown() {
	ctx.Entities.Save()
}

func (c context) GetEntities(page, pageSize int) (husk.Collection, error) {
	return c.Entities.Find(page, pageSize, husk.Everything())
}

func (c context) GetEntity(k husk.Key) (husk.Recorder, error) {
	return c.Entities.FindByKey(k)
}

//GetUser returns the User by ID
func (c context) GetUser(id string) prime.Userer {
	k, err := husk.ParseKey(id)

	if err != nil {
		log.Println("GetUser Parse Error", err)
		return nil
	}

	entity, err := c.Entities.FindByKey(k)

	if err != nil {
		log.Println("GetUser Find Error", err)
		return nil
	}

	data := entity.Data().(Entity)
	return data.User
}

//GetUserByName returns the User & ID by username
func (c context) GetUserByName(username string) (string, prime.Userer) {
	entity, err := c.Entities.FindFirst(byEmail(username))

	if err != nil {
		log.Println("GetUserByName Find Error", err)
		return "", nil
	}

	data := entity.Data().(Entity)

	return entity.GetKey().String(), data.User
}

func (c context) GetUsers(page, size int) []SafeUser {
	var result []SafeUser
	users, err := c.Entities.Find(page, size, husk.Everything())

	if err != nil {
		log.Println(err)
		return nil
	}

	itor := users.GetEnumerator()

	for itor.MoveNext() {
		currEntity := itor.Current().Data().(Entity)
		currUser := currEntity.User

		sfeUser := createSafeUser(itor.Current().GetKey(), currUser)
		result = append(result, sfeUser)
	}

	return result
}

func (c context) Register(obj Registration) error {
	return errors.New("not implemented")
	//c.Entities.Create(obj.)
}

//ResetRequest When users forget their passwords, we create a redeemable 'Reset Request' which can be used to reset their password.
//returns the Request Link or an error
func (c context) RequestReset(email, host string) (string, error) {
	id, usr := c.GetUserByName(email)

	if usr == nil {
		return "", errors.New("user not found")
	}

	k, _ := husk.ParseKey(id)
	forget := Forgot{
		UserKey:  k,
		Redeemed: false,
	}

	forgt, err := ctx.Forgotten.Create(forget)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", host, forgt.GetKey()), nil
}

func (c context) ResetPassword(forgotKey husk.Key, password string) error {
	rec, err := ctx.Forgotten.FindByKey(forgotKey)

	if err != nil {
		return err
	}

	forgetData := rec.Data().(Forgot)

	if forgetData.Redeemed {
		return errors.New("already redeemed")
	}

	if len(password) < 6 {
		return errors.New("password length must be 6 or more characters")
	}

	_, err = ctx.Forgotten.FindByKey(forgetData.UserKey)

	if err != nil {
		return err
	}

	//Change the Users password
	//usrRec.SecurePassword(password)

	//Redeem the Forgot
	forgetData.Redeemed = true

	err = ctx.Entities.Save()

	if err != nil {
		return err
	}

	return ctx.Forgotten.Save()
}
