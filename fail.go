package fail

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

// Diff is used to create a failure report for any non-matching values.
type Diff struct {
	Func string
	Msg  string
	Have interface{}
	Want interface{}
	Opts []cmp.Option
}

// String returns a human-readable report of the differences between the actual
// (have) and expected (want) values.
// It uses `cmp.Diff()` to do so.
// Options added to the Opts fields are passed to `cmp.Diff()`.
func (fd Diff) String() string {
	return fmt.Sprintf("%s() %s\n%s", fd.Func, fd.Msg, diff(fd.Have, fd.Want, fd.Opts))
}

// Err is used to create a failure report for any unexpected encountered errors.
type Err struct {
	Func string
	Err  error
}

// String returns a human-readable failure report of the unexpected error.
func (fe Err) String() string {
	return fmt.Sprintf("%s() unexpected error:\n%s", fe.Func, fe.Err.Error())
}

// Msg is used to create a simple failure report with only a message.
type Msg struct {
	Func string
	Msg  string
}

// String returns the human-readable failure report.
func (msg Msg) String() string {
	return fmt.Sprintf("%s() %s", msg.Func, msg.Msg)
}

// RetVal is used to create a failure report of non-matching values from a
// function.
type RetVal struct {
	Func string
	Msg  string
	Have []interface{}
	Want []interface{}
	Opts []cmp.Option
}

// String returns a human-readable report of the differences between the
// returned value (have) and expected (want) values.
// It uses `cmp.Diff()` to do so.
// Options added to the Opts fields are passed to `cmp.Diff()`.
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
		b.WriteString(diff(rv.Have[i], rv.Want[i], rv.Opts))
	}

	return fmt.Sprintf("%s() %s\n%s", rv.Func, rv.Msg, b.String())
}

func diff(h, w interface{}, opts []cmp.Option) string {
	return strings.TrimSpace(cmp.Diff(h, w, opts...)) + "\n"
}
