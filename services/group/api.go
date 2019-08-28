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
// The "input" value needs to be a string representation of an IP address or CIDR block.
// https://tools.ietf.org/html/rfc4632
func (g *Group) GetWhitelistEntry(input string) (*GetWhitelistEntryOutput, error) {
    req, out := g.GetWhitelistEntryRequest(input)
    return out, req.Send()
}

// GetWhitelistEntryRequest generates a "common/request.Request" representing the client's request
// for the GetWhitelistEntry operation.
//
// The "out" return value will be populated with the request's response once the request
// completes successfully.
//
// The "input" value needs to be a string representation of an IP address or CIDR block.
// https://tools.ietf.org/html/rfc4632
func (g *Group) GetWhitelistEntryRequest(input string) (req *request.Request, out *GetWhitelistEntryOutput) {
    op := &request.Operation{
        Name:       "GetWhitelistEntry",
        HTTPMethod: "GET",
        HTTPPath:   "/groups/" + g.GroupID + "/whitelist/" + url.QueryEscape(input),
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
func (g *Group) SetWhitelistEntry(input *SetWhitelistEntryInput) (error) {
    req := g.SetWhitelistEntryRequest(input)
    return req.Send()
}

// GetWhitelistEntryRequest generates a "common/request.Request" representing the client's request
// for the GetWhitelistEntry operation.
//
// The "input" value needs to be an array of whitelist entries.
func (g *Group) SetWhitelistEntryRequest(input *SetWhitelistEntryInput) (req *request.Request) {
    op := &request.Operation{
        Name:       "SetWhitelistEntry",
        HTTPMethod: "POST",
        HTTPPath:   "/groups/" + g.GroupID + "/whitelist",
    }

    handlers := &request.Handlers{
        RequestHandler: request.RequestHandler,
    }

    req = g.newRequest(op, input, nil, handlers)

    return req
}
