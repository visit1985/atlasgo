package group

import (
    "github.com/visit1985/atlasgo/common/client"
    "github.com/visit1985/atlasgo/common/request"
)

// Group provides the API operation methods for making requests to MongoDB Atlas in context of a Group (aka Project).
type Group struct{
    *client.Client
    GroupID string
}

// New creates a new instance of the Group client.
func New(gid string) *Group {
    return &Group{
        Client: client.New().Init(),
        GroupID: gid,
    }
}

// newRequest creates a new request for a Group operation.
func (g *Group) newRequest(operation *request.Operation, input interface{}, output interface{}, handlers *request.Handlers) *request.Request {
    return request.New(g.Client, operation, input, output, handlers)
}
