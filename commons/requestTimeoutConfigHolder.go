package commons

import "time"

var RequestTimeoutConfigMap = make(map[string]*RequestTimeoutConfig)

func GetDefaultRequestTimeoutConfig() *RequestTimeoutConfig {
	return &RequestTimeoutConfig{
		TimeoutInSeconds: time.Second * 15,
	}
}

func GetTimeoutConfig(path string) *RequestTimeoutConfig {
	config := RequestTimeoutConfigMap[path]
	if config == nil {
		config = GetDefaultRequestTimeoutConfig()
	}
	return config
}

func AddRequestTimeoutConfig(path string, config *RequestTimeoutConfig) {
	RequestTimeoutConfigMap[path] = config
}
