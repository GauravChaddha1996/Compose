package globalConfigHolders

import (
	"compose/commons"
	"time"
)

var RequestTimeoutConfigMap = make(map[string]*commons.RequestTimeoutConfig)

func GetDefaultRequestTimeoutConfig() *commons.RequestTimeoutConfig {
	return &commons.RequestTimeoutConfig{
		TimeoutInSeconds: time.Second * 15,
	}
}

func GetTimeoutConfig(path string) *commons.RequestTimeoutConfig {
	config := RequestTimeoutConfigMap[path]
	if config == nil {
		config = GetDefaultRequestTimeoutConfig()
	}
	return config
}

func AddRequestTimeoutConfig(path string, config *commons.RequestTimeoutConfig) {
	RequestTimeoutConfigMap[path] = config
}
