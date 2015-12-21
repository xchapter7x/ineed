#I Need

### Written because:
* all i need is a clean way to invert control
* i want to keep as close to idiomatic go struct initialization as possible
* i dont want struct element tags (thats way too much like annotations for me)
* I dont want a over the top DI framework
* I dont want be required to have a 100 argument 'New' constructor functions in my packages
* I want to just be able to pass in fakes or real objects to initialize my structs with

### Examples:

```

package main

import (
	"fmt"

	"github.com/xchapter7x/ineed"
)

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
}

func (s Something) PrintAll() {
	fmt.Println(s.randomPriv)
	fmt.Println(s.RandomPub)
}

func main() {
	deps := i.Want().
		ToMap("RandomPub", "i am public").
		ToMap("randomPriv", "i am private")

	something := New(deps)
	something.PrintAll()
	somethingPrivate := NewWithUnexported(deps)
	somethingPrivate.PrintAll()
}
```
