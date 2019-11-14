package fail

import (
	"github.com/google/go-cmp/cmp"
)

type RetVal struct {
	Func string
	Msg  string
	Have []interface{}
	Want []interface{}
	Opts []cmp.Option
}

func (rv RetVal) String() string {
	return diff(rv.Func, rv.Msg, rv.Have, rv.Want, rv.Opts)
}
