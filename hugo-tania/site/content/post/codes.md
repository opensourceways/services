---
title: codes
servicename: codes
labels: 
- Readme
---
# Codes Service

The codes service generates codes for use with email / sms verification

## cURL


### Codes Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/codes/Codes/Create' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "expires_at": "expiry time for the code, default 5 minutes",
  "identity": "e.g. phone number or email being verified"
};
# Response
{
  "code": "string"
}
```


### Codes Verify
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/codes/Codes/Verify' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "code": "string",
  "identity": "string"
};
# Response
{}
```


