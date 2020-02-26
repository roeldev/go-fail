package fail

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

type Diff struct {
	Func string
	Msg  string
	Have interface{}
	Want interface{}
	Opts []cmp.Option
}

func (fd Diff) String() string {
	return fmt.Sprintf("%s() %s\n%s", fd.Func, fd.Msg, cmp.Diff(fd.Have, fd.Want, fd.Opts...))
}

type Err struct {
	Func string
	Err  error
}

func (fe Err) String() string {
	return fmt.Sprintf("%s() unexpected error:\n%s", fe.Func, fe.Err.Error())
}

type Msg struct {
	Func string
	Msg  string
}

func (msg Msg) String() string {
	return fmt.Sprintf("%s() %s", msg.Func, msg.Msg)
}

type RetVal struct {
	Func string
	Msg  string
	Have []interface{}
	Want []interface{}
	Opts []cmp.Option
}

func (rv RetVal) String() string {
	lh, lw := len(rv.Have), len(rv.Want)
	var l int
	if lh > lw {
		l = lh
	} else {
		l = lw
	}

	b := strings.Builder{}
	for i := 0; i < l; i++ {
		_, _ = fmt.Fprintf(&b, "[%d/%d] ", i+1, l)
		b.WriteString(cmp.Diff(rv.Have[i], rv.Want[i], rv.Opts...))
	}

	return fmt.Sprintf("%s() %s\n%s", rv.Func, rv.Msg, b.String())
}
