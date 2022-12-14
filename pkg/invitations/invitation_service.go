package invitations

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type InvitationService struct {
	services.Service
}

func NewInvitationService(sling *sling.Sling, uriTemplate string) *InvitationService {
	return &InvitationService{
		Service: services.NewService(constants.ServiceInvitationService, sling, uriTemplate),
	}
}
