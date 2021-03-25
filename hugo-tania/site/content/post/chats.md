---
title: chats
servicename: chats
labels: 
- Readme
---
Chats is a service for direct messaging

# Chats Service

The chats service enables direct messaging between one or more parties.

## cURL


### Chats CreateChat
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/chats/Chats/CreateChat' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "user_ids": [
    "string"
  ]
};
# Response
{
  "chat": {
    "created_at": "string",
    "id": "string",
    "user_ids": [
      "string"
    ]
  }
}
```


### Chats CreateMessage
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/chats/Chats/CreateMessage' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "author_id": "string",
  "chat_id": "string",
  "id": "string",
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
> curl 'https://api.m3o.com/chats/Chats/ListMessages' \
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
    {
      "author_id": "string",
      "chat_id": "string",
      "id": "string",
      "sent_at": "string",
      "text": "string"
    }
  ]
}
```


