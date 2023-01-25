package main

import (
	"os"
	"path/filepath"
	"nos-gitlab.nmn.io/cloudtech_team/terraform-vault-wrapper/common"
	"nos-gitlab.nmn.io/cloudtech_team/terraform-vault-wrapper/vault"
	"nos-gitlab.nmn.io/cloudtech_team/terraform-vault-wrapper/terraform"
	log "github.com/sirupsen/logrus"
)

func main() {
	configFile, _ := filepath.Abs("wrapper.hcl")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatalf("Not Load Config File: %s", err)
	}
	config := common.ReadConfig(configFile)
	vaultToken := vault.Login(&config)
	token := vault.GetCreds(vaultToken, &config)
	terraform.SetRemoteState(token)
}