package router

import "context"

type RouterCtx interface {
	SetDependencies(key, value any) RouterCtx
	GetDependencies(key any) any
}

type routerCtx struct {
	context.Context
}

func NewRouterCtx() RouterCtx {
	return &routerCtx{
		context.Background(),
	}
}

func (r *routerCtx) SetDependencies(key, value any) RouterCtx {
	r.Context = context.WithValue(r.Context, key, value)

	return r
}

func (r *routerCtx) GetDependencies(key any) any {
	return r.Context.Value(key)
}
