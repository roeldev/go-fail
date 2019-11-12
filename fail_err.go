package fail

type Err struct {
	Func string
	Err  error
}

func (err Err) String() string {
	return err.Func + "() unexpected error:\n" + err.Err.Error()
}
