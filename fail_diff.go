package fail

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type Diff struct {
	Func string
	Msg  string
	Have interface{}
	Want interface{}
	Opts []cmp.Option
}

func (diff *Diff) AllowUnexported(types ...interface{}) *Diff {
	diff.Opts = append(diff.Opts, cmp.AllowUnexported(types...))
	return diff
}

func (diff Diff) String() string {
	return fmt.Sprintf(
		"%s() %s\n%s",
		diff.Func,
		diff.Msg,
		cmp.Diff(diff.Have, diff.Want, diff.Opts...),
	)
}
