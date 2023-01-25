# Tines(Terraform-Vault-Wrapper)
## Arichtecture
["Architecture"](architecture.png)

## Tines 제공기능 
  - Terraform Remote State 저장시 Vault 연동 제공.
  - AWS - S3, DynamoDB 연동만 제공. 
  - Vault 연동설정은 HCL로 작성하며 Terraform Vault Provider 설정과 동일

### Requirement
Terraform 12.29+
아래 Example과 같이 Remote State 설정이 적용되어 있어야 함.

#### Example
```
terraform {
  backend "s3" {
    bucket         = "test-terraform-state-kr"
    key            = "test/terraform.tfstate"
    region         = "ap-northeast-2"
    encrypt        = true
    dynamodb_table = "TestTerraformStateLock"
    acl            = "bucket-owner-full-control"
  }
}
```

### Vault Configuration
tines는 working dicrecotory 에서 "wrapper.hcl" 파일 참조해서 동작.

#### Provider Document
https://registry.terraform.io/providers/hashicorp/vault/latest/docs#provider-arguments

#### Datastrore Document
https://registry.terraform.io/providers/hashicorp/vault/latest/docs/data-sources/aws_access_credentials

#### Example
```
provider "vault" {
  address   = "http://1.1.1.1:8200"
  auth_login {
    path       = "auth/approle/login"
    parameters = {
      role_id   = "xxxxxx"
      secret_id = "xxxxxx"
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
