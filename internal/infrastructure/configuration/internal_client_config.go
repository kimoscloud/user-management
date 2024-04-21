package configuration

import "os"

type InternalClientConfig struct {
	internalClientId     string
	internalClientSecret string
}

var internalClientConfig *InternalClientConfig

func GetInternalClientConfig() *InternalClientConfig {
	if internalClientConfig == nil {
		initClientConfig()
	}
	return internalClientConfig
}

func initClientConfig() {
	internalClientConfig = &InternalClientConfig{}
	internalClientConfig.internalClientId = os.Getenv("INTERNAL_CLIENT_ID")
	internalClientConfig.internalClientSecret = os.Getenv("INTERNAL_CLIENT_SECRET")

}

// Getters for ClientConfig

func (internalClientConfig *InternalClientConfig) GetInternalClientId() string {
	return internalClientConfig.internalClientId
}
func (interalClientConfig *InternalClientConfig) GetInternalClientSecret() string {
	return interalClientConfig.internalClientSecret
}
