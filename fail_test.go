package fail

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

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
{string}:
	-: "foo"
	+: "bar"
`,
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
[1/2] {string}:
	-: "foo"
	+: "bar"
[2/2] {bool}:
	-: true
	+: false`,
		},
	}

	for _, tc := range tests {
		t.Run(reflect.TypeOf(tc.comp).String(), func(t *testing.T) {
			// ignore whitespace at start/end that might not match
			have := strings.TrimSpace(tc.comp.String())
			want := strings.TrimSpace(tc.want)

			if have != want {
				t.Error("\thave:\n", have, "\twant:\n", want)
			}
		})
	}
}
