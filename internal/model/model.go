package model

type Model[S any] interface {
	State() S
	UpdateState(state S)
}

type BaseModel[S any] struct {
	state S
}

func (bm *BaseModel[S]) State() S {
	return bm.state
}

func (bm *BaseModel[S]) UpdateState(state S) {
	bm.state = state
}
