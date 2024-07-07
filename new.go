package mapfactory

import (
	"github.com/nazarifard/bigtype"
	"github.com/nazarifard/gomap"
)

func NewMap[K StringKey, V any](engine MapEngine, options ...any) Map[K, V] {
	switch engine {
	case GoMap:
		// if len(options)>0 {
		// 	hintSize, ok :=options[0].(int)
		// 	if ok {
		// 		return gomap.New[K,V]()
		// 	}
		// }
		return gomap.New[K, V]()
	case BigMap:
		return bigtype.NewMap[K, V](options...)
	default:
		panic("invalid map engine")
	}
}
