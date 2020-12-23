package result

import (
	"fmt"
	"ginSkill/src/validators"
)

type ErrorResult struct {
	Err  error
	Data interface{}
}

func (this *ErrorResult) Unwrap() interface{} {
	if this.Err != nil {
		validators.CheckErrors(this.Err)
		panic(this.Err.Error())
	}

	return this.Data
}

func Result(vs ...interface{}) *ErrorResult {
	length := len(vs)
	if length == 1 {
		if vs[0] == nil {
			return &ErrorResult{Err: nil, Data: nil}
		}
		if e, ok := vs[0].(error); ok {
			if e != nil {
				return &ErrorResult{Err: e, Data: nil}
			}
		}
	} else if length == 2 {
		if vs[1] == nil {
			return &ErrorResult{Err: nil, Data: vs[0]}
		}
		if e, ok := vs[1].(error); ok {
			return &ErrorResult{Err: e, Data: vs[0]}
		}
	}

	return &ErrorResult{Err: fmt.Errorf("error result format"), Data: nil}
}
