package group

import (
	"github.com/visit1985/atlasgo/common/client"
	"github.com/visit1985/atlasgo/common/request"
)

type Group struct{
	*client.Client
}

func New(gid string) *Group {
	return &Group{
		Client: client.New(gid).Init(),
	}
}

func (c *Group) NewRequest(operation *request.Operation, input interface{}, output interface{}, handlers *request.Handlers) *request.Request {
	return request.New(c.Client, operation, input, output, handlers)
}
