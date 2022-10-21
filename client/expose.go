package client

import pb "github.com/casbin/casbin-server/proto"

func NewEnforcer(handler int32, remoteClient pb.CasbinClient) *Enforcer {
	return &Enforcer{
		handler: handler,
		client: &Client{
			remoteClient: remoteClient,
		},
	}
}
