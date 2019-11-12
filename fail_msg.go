package fail

type Msg struct {
	Func string
	Msg  string
}

func (msg Msg) String() string {
	return msg.Func + "() " + msg.Msg
}
