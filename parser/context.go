package parser

import (
	"sync"
)

type Context struct {
	readCache sync.Map
}

func NewContext() *Context {
	return &Context{}
}

type ServiceContext struct {
	Context *Context
}

func (self *Context) NewServiceContext() *ServiceContext {
	return &ServiceContext{
		Context: self,
	}
}
