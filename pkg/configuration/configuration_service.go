package configuration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type ConfigurationService struct {
	versionControlClearCachePath string

	services.Service
}

func NewConfigurationService(sling *sling.Sling, uriTemplate string, versionControlClearCachePath string) *ConfigurationService {
	return &ConfigurationService{
		versionControlClearCachePath: versionControlClearCachePath,
		Service:                      services.NewService(constants.ServiceConfigurationService, sling, uriTemplate),
	}
}

// GetByID returns a ConfigurationSection that matches the input ID. If one cannot be found, it returns nil and an error.
func (s *ConfigurationService) GetByID(id string) (*ConfigurationSection, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(ConfigurationSection), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ConfigurationSection), nil
}
