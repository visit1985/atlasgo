package group

import "github.com/visit1985/atlasgo/common/request"


type GetIpWhitelistOutput []struct {
	CidrBlock string `json:"cidrBlock"`
	Comment   string `json:"comment"`
	GroupID   string `json:"groupId"`
	IPAddress string `json:"ipAddress,omitempty"`
}

func (g *Group) GetIpWhitelist() (*GetIpWhitelistOutput, error) {
	req, out := g.GetIpWhitelistRequest()

	// check for client errors before sending request
	if g.Error != nil {
		return out, g.Error
	}

	if req.Error != nil {
		return out, req.Error
	}

	return out, req.Send()
}

func (g *Group) GetIpWhitelistRequest() (req *request.Request, out *GetIpWhitelistOutput) {
	op := &request.Operation{
		Name:		"GetIpWhitelist",
		HTTPMethod:	"GET",
		HTTPPath:	"/groups/" + g.GroupID + "/whitelist",
	}

	out = &GetIpWhitelistOutput{}

	handlers := &request.Handlers {
		ReponseHandler: request.ListReponseHandler,
	}

	// TODO: add paginator
	req = g.NewRequest(op, nil, out, handlers)
	return req, out
}
