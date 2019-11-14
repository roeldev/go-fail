package fail

import (
	"github.com/google/go-cmp/cmp"
)

type Diff struct {
	Func string
	Msg  string
	Have interface{}
	Want interface{}
	Opts []cmp.Option
}

func (d Diff) String() string {
	return diff(d.Func, d.Msg, d.Have, d.Want, d.Opts)
}
