package context

import (
	"context"
	"github.com/labstack/echo/v4"
)

type ServerContext struct {
	echo.Context
	ctx context.Context
}

func New(echoCtx echo.Context, ctx context.Context) *ServerContext  {
	return &ServerContext{Context: echoCtx, ctx: ctx}
}

func (c *ServerContext) Ctx() context.Context {
	return c.ctx
}

func (c *ServerContext) SetCtxValue(name string, value interface{})  {
	c.ctx = context.WithValue(c.ctx, name, value)
}