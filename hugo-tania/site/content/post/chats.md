---
title: chats
servicename: chats
labels: 
- Readme
---
# Chats Service

This is the Chats service

## cURL


### Chats CreateChat
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Chats/CreateChat' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "user_ids": [
    {}
  ]
};
# Response
{
  "chat": {
    "created_at": "string",
    "id": "string",
    "user_ids": [
      {}
    ]
  }
}
```


### Chats CreateMessage
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Chats/CreateMessage' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "author_id": "string",
  "chat_id": "string",
  "text": "string"
};
# Response
{
  "message": {
    "author_id": "string",
    "chat_id": "string",
    "id": "string",
    "sent_at": "string",
    "text": "string"
  }
}
```


### Chats ListMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Chats/ListMessages' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "chat_id": "string",
  "limit": {},
  "sent_before": "string"
};
# Response
{
  "messages": [
    {}
  ]
}
```


