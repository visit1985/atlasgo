package group

import "github.com/visit1985/atlasgo/common/client"

type Group struct{
	*client.Client
}

func New(gid string) *Group {
	return &Group{
		Client: client.New(gid).Init(),
	}
}
