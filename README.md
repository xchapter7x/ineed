#I Need

### Written because:
* all i need is a clean way to invert control
* I dont want DI
* I dont want be required to have a 'New' constructor functions in my packages with 100 arguments
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
