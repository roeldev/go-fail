package fail

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func compress(str string) string {
	b := strings.Builder{}
	b.Grow(len(str))

	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}

	return b.String()
}

func Test(t *testing.T) {
	tests := map[string]struct {
		comp fmt.Stringer
		want string
	}{
		"diff": {
			comp: Diff{
				Func: "Test",
				Msg:  "some message",
				Have: "foo",
				Want: "bar",
			},
			want: `
Test() some message
string(
- 	"foo",
+ 	"bar",
  )`,
		},
		"error": {
			comp: Err{
				Func: "Test",
				Err:  errors.New("error message"),
			},
			want: "Test() unexpected error:\nerror message",
		},
		"message": {
			comp: Msg{
				Func: "Test",
				Msg:  "some message",
			},
			want: `Test() some message`,
		},
		"retval": {
			comp: RetVal{
				Func: "Test",
				Msg:  "diff",
				Have: []interface{}{"foo", true},
				Want: []interface{}{"bar", false},
			},
			want: `
Test() diff
[1/2] string(
- 	"foo",
+ 	"bar",
  )
[2/2] bool(
- 	true,
+ 	false,
  )`,
		},
		"retval with too few values": {
			comp: RetVal{
				Func: "Test",
				Msg:  "diff, too few values",
				Have: []interface{}{"foo", true},
				Want: []interface{}{"bar"},
			},
			want: `
Test() diff, too few values
[1/2] string(
- 	"foo",
+ 	"bar",
  )
[2/2] interface{}(
- 	bool(true),
  )`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			have := tc.comp.String()
			// ignore any whitespace that might not match
			if compress(have) != compress(tc.want) {
				t.Error("have:\n", have, "\n\nwant:\n", tc.want)
			}
		})
	}
}
