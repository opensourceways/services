---
title: seen
servicename: seen
labels: 
- Readme
---
Seen is a service to keep track of which resources a user has seen (read). For example, it can be used to keep track of what notifications have been seen by a user, or what messages they've read in a chat.


# Seen Service

The seen service is a service to keep track of which resources a user has seen (read).

## cURL


### Seen Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/seen/Seen/Read' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "resource_ids": [
    "string"
  ],
  "resource_type": "string",
  "user_id": "string"
};
# Response
{
  "timestamps": [
    {
      "key": "string",
      "value": "string"
    }
  ]
}
```


### Seen Set
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/seen/Seen/Set' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "resource_id": "string",
  "resource_type": "string",
  "timestamp": "string",
  "user_id": "string"
};
# Response
{}
```


### Seen Unset
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/seen/Seen/Unset' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "resource_id": "string",
  "resource_type": "string",
  "user_id": "string"
};
# Response
{}
```


