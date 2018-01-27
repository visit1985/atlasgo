package group

import (
    "github.com/visit1985/atlasgo/common/client"
    "github.com/visit1985/atlasgo/common/request"
)

// Group provides the API operation methods for making requests to MongoDB Atlas Group Service.
type Group struct{
    *client.Client
}

// New creates a new instance of the Group client.
func New(gid string) *Group {
    return &Group{
        Client: client.New(gid).Init(),
    }
}

// newRequest creates a new request for a Group operation.
func (c *Group) newRequest(operation *request.Operation, input interface{}, output interface{}, handlers *request.Handlers) *request.Request {
    return request.New(c.Client, operation, input, output, handlers)
}
