# gcloud-whoami

This is just a simple whoami tool to show the current Application Default Crendetials that is configured and used by the gcloud cli tool.

It just a simple tool to also understand the GCP IAM authentication better and how to handle different authentications with golang.

> [!CAUTION]
> It will print sensitive information to Stdout. So use it with care or just copy the code.

## Installation

```bash
go install github.com/fr12k/gcloud-whoami
```

## Usage

```bash
gcloud-whoami
```

## Example Output

For an user account the output will look this.

```bash
Application Default Credentials:
{
  "account": "",
  "client_id": "xxxxxxxxxxxxx.apps.googleusercontent.com",
  "client_secret": "xxxxxxxxxxx",
  "quota_project_id": "xxxxxxxxxxx",
  "refresh_token": "xxxxxxxxxxxxx",
  "type": "authorized_user",
  "universe_domain": "googleapis.com"
}

id_token:
&{AccessToken:ya29.a0AcM612z_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAIGJiD4w scope:https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/cloud-platform openid https://www.googleapis.com/auth/sqlservice.login token_type:Bearer] expiryDelta:0}

Running as a user account
Authenticated as: john.doe@example.com
```

For an impersonated service account the output will look like this:

> [!IMPORTANT]
> The user account needs to have the `Service Account Token Creator` granted role on the impersonated service account.

```bash
Application Default Credentials:
{
  "delegates": [],
  "service_account_impersonation_url": "https://iamcredentials.googleapis.com/v1/projects/-/serviceAccounts/xxxx@xxxx.gserviceaccount.com:generateAccessToken",
  "source_credentials": {
    "account": "",
    "client_id": "xxxxxx.apps.googleusercontent.com",
    "client_secret": "xxxxx",
    "refresh_token": "xxxxxx",
    "type": "authorized_user",
    "universe_domain": "googleapis.com"
  },
  "type": "impersonated_service_account"
}

id_token:
&{AccessToken:ya29.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxiagX TokenType:Bearer RefreshToken: Expiry:2024-08-28 15:40:25 +0000 UTC raw:<nil> expiryDelta:0}
Most likely as a service account
Authenticated as: xxxx@[gcp_project].iam.gserviceaccount.com
```

In case the user account don't have the `Service Account Token Creator` role on the service account you will receive the following error.
```bash
2024/08/28 16:49:13 WARN Failed to get token: error="oauth2/google: status code 403: {\n  \"error\": {\n    \"code\": 403,\n    \"message\": \"Permission 'iam.serviceAccounts.getAccessToken' denied on resource (or it may not exist).\",\n    \"status\": \"PERMISSION_DENIED\",\n    \"details\": [\n      {\n        \"@type\": \"type.googleapis.com/google.rpc.ErrorInfo\",\n        \"reason\": \"IAM_PERMISSION_DENIED\",\n        \"domain\": \"iam.googleapis.com\",\n        \"metadata\": {\n          \"permission\": \"iam.serviceAccounts.getAccessToken\"\n        }\n      }\n    ]\n  }\n}\n"
Most likely as a service account
2024/08/28 16:49:13 Failed to get user info: Get "https://www.googleapis.com/userinfo/v2/me?alt=json&prettyPrint=false": credentials: status code 403: {
  "error": {
    "code": 403,
    "message": "Permission 'iam.serviceAccounts.getAccessToken' denied on resource (or it may not exist).",
    "status": "PERMISSION_DENIED",
    "details": [
      {
        "@type": "type.googleapis.com/google.rpc.ErrorInfo",
        "reason": "IAM_PERMISSION_DENIED",
        "domain": "iam.googleapis.com",
        "metadata": {
          "permission": "iam.serviceAccounts.getAccessToken"
        }
      }
    ]
  }
}
```
