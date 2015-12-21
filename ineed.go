package i

import (
	"github.com/oleiade/reflections"
	"github.com/pborman/uuid"
)

type (
	want map[string]interface{}
	Need interface {
		MapInto(interface{})
		Get(string) interface{}
		CastInto(interface{})
	}
)

func (s want) Get(name string) interface{} {
	return s[name]
}

func (s want) ToMap(n string, d interface{}) want {
	s[n] = d
	return s
}

func (s want) ToUse(d interface{}) want {
	s.ToMap(uuid.New(), d)
	return s
}

func (s want) CastInto(obj interface{}) {
	fields, _ := reflections.Fields(obj)
	for _, fieldname := range fields {

		for _, dependency := range s {
			reflections.SetField(obj, fieldname, dependency)
		}
	}
}

func (s want) MapInto(obj interface{}) {
	for n, v := range s {
		reflections.SetField(obj, n, v)
	}
}

func Want() want {
	return make(want)
}
