package globalConfigHolders

import "compose/commons"

var EndpointSecurityConfigMap = make(map[string]*commons.EndpointSecurityConfig)

func GetDefaultEndpointSecurityConfig() *commons.EndpointSecurityConfig {
	return &commons.EndpointSecurityConfig{
		CheckAccessToken: true,
		CheckUserId:      true,
		CheckUserEmail:   true,
	}
}

func AddEndpointSecurityConfig(path string, config *commons.EndpointSecurityConfig) {
	EndpointSecurityConfigMap[path] = config
}
