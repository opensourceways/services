---
title: streams
servicename: streams
labels: 
- Readme
---
# Streams Service

This is the Streams service

Generated with

```
micro new streams
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


### Streams CreateConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Streams/CreateConversation' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "group_id": "string",
  "topic": "string"
};
# Response
{
  "conversation": [
    {
      "created_at": "string",
      "group_id": "string",
      "id": "string",
      "topic": "string"
    }
  ]
}
```


### Streams CreateMessage
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Streams/CreateMessage' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "author_id": "string",
  "conversation_id": "string",
  "text": "string"
};
# Response
{
  "message": [
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


### Streams DeleteConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Streams/DeleteConversation' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{}
```


### Streams ListConversations
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Streams/ListConversations' \
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


### Streams ListMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Streams/ListMessages' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "conversation_id": "string",
  "limit": [
    {}
  ],
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


### Streams ReadConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Streams/ReadConversation' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "group_id": [
    {}
  ],
  "id": "string"
};
# Response
{
  "conversation": [
    {
      "created_at": "string",
      "group_id": "string",
      "id": "string",
      "topic": "string"
    }
  ]
}
```


### Streams RecentMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Streams/RecentMessages' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "limit_per_conversation": [
    {}
  ]
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


### Streams UpdateConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Streams/UpdateConversation' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string",
  "topic": "string"
};
# Response
{
  "conversation": [
    {
      "created_at": "string",
      "group_id": "string",
      "id": "string",
      "topic": "string"
    }
  ]
}
```


