package src_learn

import "encoding/json"

var i int = 1 << 3

type Sprinter[T json.Marshaler] interface {
	Sprint() string
}

type mySprinter struct {
	s string
}

func (this *mySprinter) Sprint() string {
	return this.s
}
