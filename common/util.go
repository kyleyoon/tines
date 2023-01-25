package common

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
	log "github.com/sirupsen/logrus"
	"github.com/zclconf/go-cty/cty"
)

// VaultBackend is hcl parse
type VaultBackend struct {
	Provider ProviderConfig `hcl:"provider,block"`
	Data     DataConfig     `hcl:"data,block"`
}

// ProviderConfig is hcl parse
type ProviderConfig struct {
	ProviderName string          `hcl:"provider_name,label"`
	VaultDomain  string          `hcl:"address"`
	AuthLogin    AuthLoginConfig `hcl:"auth_login,block"`
}

// AuthLoginConfig is hcl parse
type AuthLoginConfig struct {
	ReqPath string    `hcl:"path"`
	Params  cty.Value `hcl:"parameters"`
}

// DataConfig is hcl parse
type DataConfig struct {
	DataMethod  string `hcl:"data_method,label"`
	DataType    string `hcl:"data_type,label"`
	BackendName string `hcl:"backend"`
	BackendRole string `hcl:"role"`
	BackendType string `hcl:"type"`
}

// ParamsConfig is hcl parse
type ParamsConfig struct {
	RoleID   string `cty:"role_id" json:"role_id"`
	SecretID string `cty:"secret_id" json:"secret_id"`
}

// ReadConfig is hcl parse
func ReadConfig(file string) VaultBackend {
	var config VaultBackend
	err := hclsimple.DecodeFile(file, nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Debugf("Configuration is %#v", config)
	return config
}