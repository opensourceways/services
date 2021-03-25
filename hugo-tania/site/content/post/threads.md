---
title: threads
servicename: threads
labels: 
- Readme
---
Threaded conversations

# Threads Service

Threads provides threaded conversations as a service grouped by topics.

## Usage

Generated with

```
micro new threads
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```

## cURL


### Threads CreateConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/CreateConversation' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "group_id": "string",
  "topic": "string"
};
# Response
{
  "conversation": {
    "created_at": "string",
    "group_id": "string",
    "id": "string",
    "topic": "string"
  }
}
```


### Threads CreateMessage
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/CreateMessage' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "author_id": "string",
  "conversation_id": "string",
  "id": "string",
  "text": "string"
};
# Response
{
  "message": {
    "author_id": "string",
    "conversation_id": "string",
    "id": "string",
    "sent_at": "string",
    "text": "string"
  }
}
```


### Threads DeleteConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/DeleteConversation' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{}
```


### Threads ListConversations
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/ListConversations' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "group_id": "string"
};
# Response
{
  "conversations": [
    {
      "created_at": "string",
      "group_id": "string",
      "id": "string",
      "topic": "string"
    }
  ]
}
```


### Threads ListMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/ListMessages' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "conversation_id": "string",
  "limit": {},
  "sent_before": "string"
};
# Response
{
  "messages": [
    {
      "author_id": "string",
      "conversation_id": "string",
      "id": "string",
      "sent_at": "string",
      "text": "string"
    }
  ]
}
```


### Threads ReadConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/ReadConversation' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "group_id": {},
  "id": "string"
};
# Response
{
  "conversation": {
    "created_at": "string",
    "group_id": "string",
    "id": "string",
    "topic": "string"
  }
}
```


### Threads RecentMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/RecentMessages' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "conversation_ids": [
    "string"
  ],
  "limit_per_conversation": {}
};
# Response
{
  "messages": [
    {
      "author_id": "string",
      "conversation_id": "string",
      "id": "string",
      "sent_at": "string",
      "text": "string"
    }
  ]
}
```


### Threads UpdateConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/UpdateConversation' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string",
  "topic": "string"
};
# Response
{
  "conversation": {
    "created_at": "string",
    "group_id": "string",
    "id": "string",
    "topic": "string"
  }
}
```


