package fail

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
)

type Diff struct {
	Func string
	Msg  string
	Have interface{}
	Want interface{}
	Opts []cmp.Option
}

// AllowUnexported is a wrapper around cmp.AllowUnexported and add it to the Opts slice.
func (diff *Diff) AllowUnexported(types ...interface{}) *Diff {
	diff.Opts = append(diff.Opts, cmp.AllowUnexported(types...))
	return diff
}

// AllowUnexportedPtrs converts pointer types to their original values. It then passes those types
// to AllowUnexported().
func (diff *Diff) AllowUnexportedPtrs(types ...interface{}) *Diff {
	for i, typ := range types {
		rVal := reflect.ValueOf(typ)
		if rVal.Kind() == reflect.Ptr {
			types[i] = reflect.Indirect(rVal).Interface()
		}
	}

	return diff.AllowUnexported(types...)
}

func (diff Diff) String() string {
	return diff.Func + "() " + diff.Msg + "\n" + cmp.Diff(diff.Have, diff.Want, diff.Opts...)
}
