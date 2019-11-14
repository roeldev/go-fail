package fail

import (
	"github.com/google/go-cmp/cmp"
)

func diff(f string, m string, h interface{}, w interface{}, o []cmp.Option) string {
	return f + "() " + m + "\n" + cmp.Diff(h, w, o...)
}
