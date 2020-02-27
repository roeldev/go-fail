package fail

import (
	"errors"
	"fmt"
	"reflect"
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
	tests := []struct {
		comp fmt.Stringer
		want string
	}{
		{
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
		{
			comp: Err{
				Func: "Test",
				Err:  errors.New("error message"),
			},
			want: "Test() unexpected error:\nerror message",
		},
		{
			comp: Msg{
				Func: "Test",
				Msg:  "some message",
			},
			want: `Test() some message`,
		},
		{
			comp: RetVal{
				Func: "Test",
				Msg:  "some message",
				Have: []interface{}{"foo", true},
				Want: []interface{}{"bar", false},
			},
			want: `
Test() some message
[1/2] string(
- 	"foo",
+ 	"bar",
  )
[2/2] bool(
- 	true,
+ 	false,
  )`,
		},
	}

	for _, tc := range tests {
		t.Run(reflect.TypeOf(tc.comp).String(), func(t *testing.T) {
			have := tc.comp.String()
			// ignore any whitespace that might not match
			if compress(have) != compress(tc.want) {
				t.Error("have:\n", have, "\n\nwant:\n", tc.want)
			}
		})
	}
}
