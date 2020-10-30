package testdata

// FuncCaller form of function's call which includes expected input and output
type FuncCaller struct {
	IsCalled bool
	Input    []interface{}
	Output   []interface{}
}
