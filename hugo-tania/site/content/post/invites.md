---
title: invites
servicename: invites
labels: 
- Readme
---
# Invites Service

The invites services allows you to create and manage invites. Example usage:

```bash
> micro invites create --group_id=myawesomegroup --email=john@doe.com
{
	"invite": {
		"id": "fb3a3552-3c7b-4a18-a1f8-08ab56940862",
		"group_id": "myawesomegroup",
		"email": "john@doe.com",
		"code": "86285587"
	}
}

> micro invites list --group_id=fb3a3552-3c7b-4a18-a1f8-08ab56940862
{
	"invites": [
    {
    "id": "fb3a3552-3c7b-4a18-a1f8-08ab56940862",
    "group_id": "myawesomegroup",
    "email": "john@doe.com",
    "code": "86285587"
    }
  ]
}

> micro invites read --code=86285587
{
	"invite": {
		"id": "fb3a3552-3c7b-4a18-a1f8-08ab56940862",
		"group_id": "myawesomegroup",
		"email": "john@doe.com",
		"code": "86285587"
	}
}

> micro invites delete --id=fb3a3552-3c7b-4a18-a1f8-08ab56940862
{}
```

## cURL


### Invites Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Invites/Create' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "group_id": "string"
};
# Response
{
  "invite": {
    "code": "string",
    "email": "string",
    "group_id": "string",
    "id": "string"
  }
}
```


### Invites Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Invites/Delete' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{}
```


### Invites List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Invites/List' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": {},
  "group_id": {}
};
# Response
{
  "invites": [
    {
      "code": "string",
      "email": "string",
      "group_id": "string",
      "id": "string"
    }
  ]
}
```


### Invites Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Invites/Read' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "code": {},
  "id": {}
};
# Response
{
  "invite": {
    "code": "string",
    "email": "string",
    "group_id": "string",
    "id": "string"
  }
}
```


