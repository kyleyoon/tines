# Tines(Terraform-Vault-Wrapper)
## Arichtecture
!["Architecture"](https://wiki.nmn.io/download/attachments/189404549/image2021-2-15_11-22-45.png)

## Tines 제공기능 
  - Terraform Remote State 저장시 Vault 연동 제공.
  -- AWS - S3, DynamoDB 연동만 제공. 
  - Vault 연동설정은 HCL로 작성하며 Terraform Vault Provider 설정과 동일

## Installation
### Download
- linux version
https://drive.google.com/u/0/uc?id=1-OGmYdB9T1aTBaNugUyVB6OTz79dYYXz&export=download
- window version
https://drive.google.com/u/0/uc?id=1-Oh54jD3Wg2v9z9nynoXUsyWiGO006Hx&export=download

### Requirement
Terraform 12.29+
아래 Example과 같이 Remote State 설정이 적용되어 있어야 함.

#### Example
```
terraform {
  backend "s3" {
    bucket         = "test-terraform-state-kr"
    key            = "btsw-gb/game/terraform.tfstate"
    region         = "ap-northeast-2"
    encrypt        = true
    dynamodb_table = "TestTerraformStateLock"
    acl            = "bucket-owner-full-control"
  }
}
```

### Vault Configuration
tines는 working dicrecotory 에서 "wrapper.hcl" 파일 참조해서 동작.
-- Provider Document
https://registry.terraform.io/providers/hashicorp/vault/latest/docs#provider-arguments
-- Datastrore Document
https://registry.terraform.io/providers/hashicorp/vault/latest/docs/data-sources/aws_access_credentials
#### Example
```
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
```



### Execute
Basic
```
# linux version
tines-linux

# windows version
tines-win.exe
```

Debug
```
# linux version
TF_LOG=TRACE tines-linux
```

#### Log Level
Terraform에서 사용하는 TF_LOG Environment 설정시 Tines도 Log Level이 적용됨.
-- Default: INFO

|TF_LOG|Tines Log Level|
|------|---------------|
|TRACE|TRACE|
|INFO|INFO|