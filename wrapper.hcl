provider "vault" {
  address   = "http://183.110.89.27:8200"
  auth_login {
    path       = "auth/approle/login"
    parameters = {
      role_id   = "5c3fd53a-d5cf-5f73-20b2-fc8d0e79632b"
      secret_id = "7e29aa92-7bca-6985-191f-a288bec084ed"
    }
  }
}

data "vault_aws_access_credentials" "service" {
  backend = "remote-state"
  role    = "assumed"
  type    = "sts"
}