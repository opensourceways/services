---
title: groups
servicename: groups
labels: 
- Readme
---
# Groups Service

The group serivce is a basic CRUD service for groups. You can use it to create groups, add members and lookup which groups a user is a member of.

Example usage:

```bash
$ micro groups create --name=Micro
{
	"group": {
		"id": "e35562c9-b6f6-459a-b52d-7e6159465fd6",
		"name": "Micro"
	}
}
$ micro groups addMember --group_id=e35562c9-b6f6-459a-b52d-7e6159465fd6 --member_id=Asim
{}
$ micro groups list --member_id=Asim
{
	"groups": [
		{
			"id": "e35562c9-b6f6-459a-b52d-7e6159465fd6",
			"name": "Micro",
			"member_ids": [
				"Asim"
			]
		}
	]
}
$ micro groups list --member_id=Boris
{}
```

## cURL


### Groups AddMember
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/groups/Groups/AddMember' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "group_id": "string",
  "member_id": "string"
};
# Response
{}
```


### Groups Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/groups/Groups/Create' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "name": "string"
};
# Response
{
  "group": {
    "id": "string",
    "member_ids": [
      "string"
    ],
    "name": "string"
  }
}
```


### Groups Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/groups/Groups/Delete' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{}
```


### Groups List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/groups/Groups/List' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "member_id": "passing a member id will restrict the groups to that which the member is part of"
};
# Response
{
  "groups": [
    {
      "id": "string",
      "member_ids": [
        "string"
      ],
      "name": "string"
    }
  ]
}
```


### Groups Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/groups/Groups/Read' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "ids": [
    "string"
  ]
};
# Response
{
  "groups": [
    {
      "key": "string",
      "value": {
        "id": "string",
        "member_ids": [
          "string"
        ],
        "name": "string"
      }
    }
  ]
}
```


### Groups RemoveMember
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/groups/Groups/RemoveMember' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "group_id": "string",
  "member_id": "string"
};
# Response
{}
```


### Groups Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/groups/Groups/Update' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string",
  "name": "string"
};
# Response
{
  "group": {
    "id": "string",
    "member_ids": [
      "string"
    ],
    "name": "string"
  }
}
```


