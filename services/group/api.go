package group

import "github.com/visit1985/atlasgo/common/request"

// A GetIpWhitelistOutput is the JSON structure a GetIpWhitelist API operation returns.
// https://docs.atlas.mongodb.com/reference/api/whitelist/#id1
type GetIpWhitelistOutput []struct {
    CidrBlock string `json:"cidrBlock"`
    Comment   string `json:"comment"`
    GroupID   string `json:"groupId"`
    IPAddress string `json:"ipAddress,omitempty"`
}

// The GetIpWhitelist API operation for MongoDB Atlas Group Service retrieves a group's IP whitelist,
// which controls client access to the group's MongoDB clusters.
func (g *Group) GetIpWhitelist() (*GetIpWhitelistOutput, error) {
    req, out := g.GetIpWhitelistRequest()
    return out, req.Paginate()
}

// GetIpWhitelistRequest generates a "common/request.Request" representing the client's request
// for the GetIpWhitelist operation. The "out" return value will be populated with the request's
// response once the request completes successfully.
func (g *Group) GetIpWhitelistRequest() (req *request.Request, out *GetIpWhitelistOutput) {
    op := &request.Operation{
        Name:       "GetIpWhitelist",
        HTTPMethod: "GET",
        HTTPPath:   "/groups/" + g.GroupID + "/whitelist",
    }

    out = &GetIpWhitelistOutput{}

    handlers := &request.Handlers {
        ResponseHandler: request.ListResponseHandler,
    }

    req = g.newRequest(op, nil, out, handlers)
    return req, out
}
