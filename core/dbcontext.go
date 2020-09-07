package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/kong/prime"
	"github.com/louisevanderlith/kong/stores"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"reflect"
)

type EntityContext interface {
	stores.UserStore
	CreateEntity(obj Entity) (hsk.Key, error)
	GetEntities(page, pageSize int) (records.Page, error)
	GetEntity(k hsk.Key) (hsk.Record, error)
	Register(obj Registration) error
	RequestReset(email, host string) (string, error)
	ResetPassword(forgotKey hsk.Key, password string) error
}

type context struct {
	Entities  husk.Table
	Forgotten husk.Table
}

func (c context) CreateEntity(obj Entity) (hsk.Key, error) {
	return c.Entities.Create(obj)
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
	entities, err := entitySeeds()

	if err != nil {
		panic(err)
	}

	err = ctx.Entities.Seed(entities)

	if err != nil {
		panic(err)
	}

	err = ctx.Entities.Save()

	if err != nil {
		panic(err)
	}
}

func entitySeeds() (collections.Enumerable, error) {
	f, err := os.Open("db/entities.seed.json")

	if err != nil {
		return nil, err
	}

	var items []Entity
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		return nil, err
	}

	return collections.ReadOnlyList(reflect.ValueOf(items)), nil
}

func Shutdown() {
	ctx.Entities.Save()
}

func (c context) GetEntities(page, pageSize int) (records.Page, error) {
	return c.Entities.Find(page, pageSize, op.Everything())
}

func (c context) GetEntity(k hsk.Key) (hsk.Record, error) {
	return c.Entities.FindByKey(k)
}

//GetUser returns the User by ID
func (c context) GetUser(id string) prime.Userer {
	k, err := keys.ParseKey(id)

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
	users, err := c.Entities.Find(page, size, op.Everything())

	if err != nil {
		log.Println(err)
		return nil
	}

	itor := users.GetEnumerator()

	for itor.MoveNext() {
		rec := itor.Current().(hsk.Record)
		currEntity := rec.Data().(Entity)
		currUser := currEntity.User

		sfeUser := createSafeUser(rec.GetKey(), currUser)
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

	k, _ := keys.ParseKey(id)
	forget := Forgot{
		EntityKey: k,
		Redeemed:  false,
	}

	k, err := ctx.Forgotten.Create(forget)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", host, k), nil
}

func (c context) ResetPassword(forgotKey hsk.Key, password string) error {
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

	entity, err := ctx.Entities.FindByKey(forgetData.EntityKey)

	if err != nil {
		return err
	}

	pss, err := bcrypt.GenerateFromPassword([]byte(password), 11)
	if err != nil {
		return err
	}

	obj := entity.Data().(Entity)
	obj.User.Password = string(pss)

	err = ctx.Entities.Save()

	if err != nil {
		return err
	}

	//Redeem the Forgot
	forgetData.Redeemed = true

	return ctx.Forgotten.Update(forgotKey, forgetData)
}
