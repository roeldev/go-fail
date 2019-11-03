package fail

import (
	"fmt"
)

type Msg struct {
	Func string
	Msg  string
}

func (msg Msg) String() string {
	return fmt.Sprintf("%s() %s", msg.Func, msg.Msg)
}
