package vault

import (
	"strings"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/zclconf/go-cty/cty/gocty"
	"nos-gitlab.nmn.io/cloudtech_team/terraform-vault-wrapper/common"
)

// Login is Terraform Remote State Backend Authentication
func Login(d *common.VaultBackend) string {
	var vaultCreds common.ParamsConfig
	err := gocty.FromCtyValue(d.Provider.AuthLogin.Params, &vaultCreds)
	log.Printf("role_id: %s, secret_id: %s \n", vaultCreds.RoleID, vaultCreds.SecretID)

	pbytes, err := json.Marshal(vaultCreds)
	if err != nil {
		log.Fatalf("vault-login: %s", err)
	}

	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post(d.Provider.VaultDomain+"/v1/auth/approle/login", "application/json", buff)

	if err != nil {
		log.Fatalf("vault-login: %s", err)
	}

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("vault-login: %s", err)
	}

	if resp.StatusCode != 200 {
	    log.Errorf("Marshaling: %s\n", string(pbytes))
		log.Errorf("Request Header: %s\n", resp.Header)
		log.Errorf("Request URL: %s\n", resp.Request.URL)
		log.Errorf("Request Return Code: %d\n", resp.StatusCode)
		log.Error(string(respBody))
		log.Fatal("Fatal Call Vault API")
	}

	defer resp.Body.Close()

	token := gjson.Get(string([]byte(respBody)), "auth.client_token")
	log.Debugf("Vault token:", token)
	return token.String()
}

// GetCreds is Get HashiCorp Vault Credentials
func GetCreds(token string, d *common.VaultBackend) map[string]gjson.Result {
	urldata := []string {
		d.Provider.VaultDomain+"/v1/",
		d.Data.BackendName+"/",
		d.Data.BackendType+"/",
		d.Data.BackendRole,
	}

	req, err := http.NewRequest("GET", strings.Join(urldata[:], ""), nil)
	if err != nil {
		log.Fatalf("vault-get-creds: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Vault-Token", token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("vault-get-creds: %s", err)
	}

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("vault-get-creds: %s", err)
	}

	if resp.StatusCode != 200 {
		log.Errorf("Request Header: %s\n", resp.Header)
		log.Errorf("Request URL: %s\n", resp.Request.URL)
		log.Errorf("Request Return Code: %d\n", resp.StatusCode)
		log.Error(string(respBody))
		log.Fatal("Fatal Call Vault API")
	}

	defer resp.Body.Close()

	credsData := gjson.Get(string([]byte(respBody)), "data")
	creds := credsData.Map()
	log.Debugf("Access Key", creds["access_key"])
	log.Debugf("Secret Key", creds["secret_key"])
	log.Debugf("Security Token", creds["security_token"])
	log.Debugf("Raw", credsData)
	return creds
}