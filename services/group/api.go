package group

import (
    "github.com/visit1985/atlasgo/common/request"
    "net/url"
)

// A GetWhitelistOutput is the JSON structure a GetWhitelist API operation returns.
// https://docs.atlas.mongodb.com/reference/api/whitelist/#id1
type GetWhitelistOutput []struct {
    CidrBlock string `json:"cidrBlock"`
    Comment   string `json:"comment"`
    GroupId   string `json:"groupId"`
    IpAddress string `json:"ipAddress,omitempty"`
}

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

// A GetAddressOutput is the JSON structure a GetAddress API operation returns.
// https://docs.atlas.mongodb.com/reference/api/whitelist/#id2
type GetAddressOutput struct {
    CidrBlock string `json:"cidrBlock"`
    Comment   string `json:"comment"`
    GroupId   string `json:"groupId"`
    IpAddress string `json:"ipAddress,omitempty"`
}

// The GetAddress API operation for MongoDB Atlas Group Service retrieves an IP address from
// a group's whitelist.
//
// The "input" value needs to be a string representation of an IP address or CIDR block.
// https://tools.ietf.org/html/rfc4632
func (g *Group) GetAddress(input string) (*GetAddressOutput, error) {
    req, out := g.GetAddressRequest(input)
    return out, req.Send()
}

// GetAddressRequest generates a "common/request.Request" representing the client's request
// for the GetAddress operation.
//
// The "out" return value will be populated with the request's response once the request
// completes successfully.
//
// The "input" value needs to be a string representation of an IP address or CIDR block.
// https://tools.ietf.org/html/rfc4632
func (g *Group) GetAddressRequest(input string) (req *request.Request, out *GetAddressOutput) {
    op := &request.Operation{
        Name:       "GetAddress",
        HTTPMethod: "GET",
        HTTPPath:   "/groups/" + g.GroupID + "/whitelist/" + url.QueryEscape(input),
    }

    out = &GetAddressOutput{}

    handlers := &request.Handlers{
        ResponseHandler: request.ResponseHandler,
    }

    req = g.newRequest(op, nil, out, handlers)

    return req, out
}
