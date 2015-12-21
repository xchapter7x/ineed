package main

import (
	"fmt"

	i "github.com/xchapter7x/ineed"
)

func NewFromType(deps i.Need) *Something {
	s := new(Something)
	deps.CastInto(s)
	return s
}

func NewWithUnexported(deps i.Need) *Something {
	s := &Something{
		randomPriv: deps.Get("randomPriv").(string),
	}
	deps.MapInto(s)
	return s
}

func New(deps i.Need) *Something {
	s := new(Something)
	deps.MapInto(s)
	return s
}

type Something struct {
	randomPriv string
	RandomPub  string
	Cool       CoolObject
}

type CoolObject struct {
	AField       string
	AnotherField string
}

func (s Something) PrintAll() {
	fmt.Println(s.randomPriv)
	fmt.Println(s.RandomPub)
	fmt.Println(s.Cool)
}

func main() {
	deps := i.Want().
		ToMap("RandomPub", "i am public").
		ToMap("randomPriv", "i am private")

	coolDep := CoolObject{
		AField:       "inject me",
		AnotherField: "Don't forget me too",
	}
	blindDeps := i.Want().
		ToUse(coolDep)

	something := New(deps)
	something.PrintAll()

	somethingPrivate := NewWithUnexported(deps)
	somethingPrivate.PrintAll()

	somethingWithRandomNamedFields := NewFromType(blindDeps)
	somethingWithRandomNamedFields.PrintAll()
}
