package fail

import (
	"fmt"
)

type Err struct {
	Func string
	Err  error
}

func (err Err) String() string {
	return fmt.Sprintf("%s() unexpected error:\n%s", err.Func, err.Err)
}
