---
title: chat
---
The chat service is an example Micro service which leverages bidirectional streaming, the store and events to build a chat backend. There is both a server and client which can be run together to demonstrate the application (see client/main.go for more instructions on running the service).

# Chat Service

The service is documented inline and is designed to act as a reference for the events package.

## Create a chat

### cURL

```bash
> curl 'https://api.m3o.com/chat/New' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d '{"user_ids":["JohnBarry"]}';
{
	"chat_id": "3c9ea66c-d516-45d4-abe8-082089e18b27"
}
```

### CLI

```bash
> micro chat new --user_ids=JohnBarry
{
	"chat_id": "3c9ea66c-d516-45d4-abe8-082089e18b27"
}
```

## Send a message to the chat

### cURL

```bash
> curl 'https://api.m3o.com/chat/Send' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d '{"user_id": "John", "subject": "Hello", "text": "Hey Barry"}';
{}
```

### CLI

```bash
> micro chat send --chat_id=bed4f0f0-da12-46d2-90d2-17ae1714a214 --user_id=John --subject=Hello --text='Hey Barry'
{}
```

## View the chat history

### cURL

```bash
> curl 'https://api.m3o.com/chat/Send' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d '{"chat_id": "bed4f0f0-da12-46d2-90d2-17ae1714a214"}';
{
	"messages": [
		{
			"id": "a61284a8-f471-4734-9192-640d89762e98",
			"client_id": "6ba0d2a6-96fa-47d8-8f6f-7f75b4cc8b3e",
			"chat_id": "bed4f0f0-da12-46d2-90d2-17ae1714a214",
			"user_id": "John",
			"subject": "Hello",
			"text": "Hey Barry"
		}
	]
}
```

### CLI
```bash
> micro chat history --chat_id=bed4f0f0-da12-46d2-90d2-17ae1714a214
{
	"messages": [
		{
			"id": "a61284a8-f471-4734-9192-640d89762e98",
			"client_id": "6ba0d2a6-96fa-47d8-8f6f-7f75b4cc8b3e",
			"chat_id": "bed4f0f0-da12-46d2-90d2-17ae1714a214",
			"user_id": "John",
			"subject": "Hello",
			"text": "Hey Barry"
		}
	]
}
```
## cURL


### Chat Connect
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Schema related to #/components/requestBodies/ChatConnectRequest not found
```shell
> curl 'https://api.m3o.com/chat/Chat/Connect' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d Schema related to #/components/requestBodies/ChatConnectRequest not found;
# Response
Schema related to #/components/responses/ChatConnectResponse not found
```


### Chat History
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
HistoryRequest contains the id of the chat we want the history for. This RPC will return all 
 historical messages, however in a real life application we'd introduce some form of pagination
 here, only loading the older messages when required.
```shell
> curl 'https://api.m3o.com/chat/Chat/History' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "chat_id": "string"
};
# Response
{
  "messages": [
    {
      "chat_id": "id of the chat the message is being sent to / from",
      "client_id": "a client side id, should be validated by the server to make the request retry safe",
      "id": "id of the message, allocated by the server",
      "sent_at": 1,
      "subject": "subject of the message",
      "text": "text of the message",
      "user_id": "id of the user who sent the message"
    }
  ]
}
```


### Chat New
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
NewRequest contains the infromation needed to create a new chat
```shell
> curl 'https://api.m3o.com/chat/Chat/New' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {};
# Response
{
  "chat_id": "string"
}
```


### Chat Send
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
SendRequest contains a single message to send to a chat
```shell
> curl 'https://api.m3o.com/chat/Chat/Send' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "chat_id": "id of the chat the message is being sent to / from",
  "client_id": "a client side id, should be validated by the server to make the request retry safe",
  "subject": "subject of the message",
  "text": "text of the message",
  "user_id": "id of the user who sent the message"
};
# Response
{}
```


