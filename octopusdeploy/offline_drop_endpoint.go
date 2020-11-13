package octopusdeploy

type OfflineDropEndpoint struct {
	ApplicationsDirectory                string                  `json:"ApplicationsDirectory,omitempty"`
	Destination                          *OfflineDropDestination `json:"Destination"`
	SensitiveVariablesEncryptionPassword *SensitiveValue         `json:"SensitiveVariablesEncryptionPassword"`
	WorkingDirectory                     string                  `json:"OctopusWorkingDirectory,omitempty"`

	endpoint
}

func NewOfflineDropEndpoint() *OfflineDropEndpoint {
	return &OfflineDropEndpoint{
		endpoint: *newEndpoint("OfflineDrop"),
	}
}
