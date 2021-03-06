package group

import (
    "github.com/visit1985/atlasgo/common/request"
    "net/url"
    "time"
)

// A WhitelistEntry is a generic JSON structure used in API operations on group IP whitelists.
// https://docs.atlas.mongodb.com/reference/api/whitelist/#sample-entity
type WhitelistEntry struct {
    CidrBlock       string     `json:"cidrBlock,omitempty"`
    Comment         string     `json:"comment,omitempty"`
    GroupId         string     `json:"groupId,omitempty"`
    IpAddress       string     `json:"ipAddress,omitempty"`
    DeleteAfterDate *time.Time `json:"deleteAfterDate,omitempty"`
}

// A GetWhitelistOutput is the JSON structure a GetWhitelist API operation returns.
// https://docs.atlas.mongodb.com/reference/api/whitelist/#id1
type GetWhitelistOutput []WhitelistEntry

// The GetWhitelist API operation for MongoDB Atlas Group Service retrieves a group's IP address
// whitelist, which controls client access to the group's MongoDB clusters.
func (g *Group) GetWhitelist() (*GetWhitelistOutput, error) {
    req, out := g.GetWhitelistRequest()
    return out, req.Paginate()
}

// GetWhitelistRequest generates a "common/request.Request" representing the client's request
// for the GetWhitelist operation.
//
// The "out" return value will be populated with the request's response once the request
// completes successfully.
func (g *Group) GetWhitelistRequest() (req *request.Request, out *GetWhitelistOutput) {
    op := &request.Operation{
        Name:       "GetWhitelist",
        HTTPMethod: "GET",
        HTTPPath:   "/groups/" + g.GroupID + "/whitelist",
    }

    out = &GetWhitelistOutput{}

    handlers := &request.Handlers {
        ResponseHandler: request.ListResponseHandler,
    }

    req = g.newRequest(op, nil, out, handlers)
    return req, out
}

// A GetWhitelistEntryOutput is the JSON structure a GetWhitelistEntry API operation returns.
// https://docs.atlas.mongodb.com/reference/api/whitelist/#id2
type GetWhitelistEntryOutput WhitelistEntry

// The GetWhitelistEntry API operation for MongoDB Atlas Group Service retrieves an IP address from
// a group's whitelist.
//
// The "in" value needs to be a string representation of an IP address or CIDR block.
// https://tools.ietf.org/html/rfc4632
func (g *Group) GetWhitelistEntry(in string) (*GetWhitelistEntryOutput, error) {
    req, out := g.GetWhitelistEntryRequest(in)
    return out, req.Send()
}

// GetWhitelistEntryRequest generates a "common/request.Request" representing the client's request
// for the GetWhitelistEntry operation.
//
// The "out" return value will be populated with the request's response once the request
// completes successfully.
//
// The "in" value needs to be a string representation of an IP address or CIDR block.
// https://tools.ietf.org/html/rfc4632
func (g *Group) GetWhitelistEntryRequest(in string) (req *request.Request, out *GetWhitelistEntryOutput) {
    op := &request.Operation{
        Name:       "GetWhitelistEntry",
        HTTPMethod: "GET",
        HTTPPath:   "/groups/" + g.GroupID + "/whitelist/" + url.QueryEscape(in),
    }

    out = &GetWhitelistEntryOutput{}

    handlers := &request.Handlers{
        ResponseHandler: request.ResponseHandler,
    }

    req = g.newRequest(op, nil, out, handlers)

    return req, out
}

// A SetWhitelistEntryInput is the input JSON structure for a SetWhitelistEntry API operation.
// https://docs.atlas.mongodb.com/reference/api/whitelist/#id5
type SetWhitelistEntryInput []WhitelistEntry

// The SetWhitelistEntry API operation for MongoDB Atlas Group Service adds an IP address or
// CIDR block to a group's whitelist.
func (g *Group) SetWhitelistEntry(in *SetWhitelistEntryInput) error {
    req := g.SetWhitelistEntryRequest(in)
    return req.Send()
}

// GetWhitelistEntryRequest generates a "common/request.Request" representing the client's request
// for the GetWhitelistEntry operation.
//
// The "out" return value will be populated with the request's response once the request
// completes successfully.
//
// The "in" value needs to be an array of whitelist entries.
func (g *Group) SetWhitelistEntryRequest(in *SetWhitelistEntryInput) (req *request.Request) {
    op := &request.Operation{
        Name:       "SetWhitelistEntry",
        HTTPMethod: "POST",
        HTTPPath:   "/groups/" + g.GroupID + "/whitelist",
    }

    handlers := &request.Handlers{
        RequestHandler:  request.RequestHandler,
    }

    req = g.newRequest(op, in, nil, handlers)

    return req
}

// The DeleteWhitelistEntry API operation for MongoDB Atlas Group Service deletes an IP address from
// a group's whitelist.
//
// The "in" value needs to be a string representation of an IP address or CIDR block.
// https://tools.ietf.org/html/rfc4632
func (g *Group) DeleteWhitelistEntry(in string) error {
    req := g.DeleteWhitelistEntryRequest(in)
    return req.Send()
}

// DeleteWhitelistEntryRequest generates a "common/request.Request" representing the client's request
// for the DeleteWhitelistEntry operation.
//
// The "in" value needs to be a string representation of an IP address or CIDR block.
// https://tools.ietf.org/html/rfc4632
func (g *Group) DeleteWhitelistEntryRequest(in string) (req *request.Request) {
    op := &request.Operation{
        Name:       "DeleteWhitelistEntry",
        HTTPMethod: "DELETE",
        HTTPPath:   "/groups/" + g.GroupID + "/whitelist/" + url.QueryEscape(in),
    }

    handlers := &request.Handlers{}

    req = g.newRequest(op, nil, nil, handlers)

    return req
}
