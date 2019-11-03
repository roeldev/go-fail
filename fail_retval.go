package fail

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type RetVal struct {
	Func string
	Msg  string
	Have []interface{}
	Want []interface{}
}

func (rv RetVal) String() string {
	return fmt.Sprintf("%s() %s\n%s", rv.Func, rv.Msg, cmp.Diff(rv.Have, rv.Want))
}
