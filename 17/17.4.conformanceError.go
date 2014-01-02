package example

type Public interface {
	Name() string
}

type Example struct {
	name string
}

//func (e Example) Nme() string {
func (e Example) Name() string {
	return e.name
}

// func NewExample() Example {
// 	return Example{"No Name"}
// }

func NewExample() Public {
	return Example{"No Name"}
}

func NewExample2() Public {
	e := Example{"No Name"}
	e.(Public)

	return e
}


//# command-line-arguments
//.\17.4.conformanceError.go:26: invalid type assertion: e.(Public) (non-interface type Example on left)