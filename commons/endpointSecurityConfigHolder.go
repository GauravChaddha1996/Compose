package commons

var EndpointSecurityConfigMap = make(map[string]*EndpointSecurityConfig)

func GetDefaultEndpointSecurityConfig() *EndpointSecurityConfig {
	return &EndpointSecurityConfig{
		CheckAccessToken: true,
		CheckUserId:      true,
		CheckUserEmail:   true,
	}
}

func AddEndpointSecurityConfig(path string, config *EndpointSecurityConfig) {
	EndpointSecurityConfigMap[path] = config
}
