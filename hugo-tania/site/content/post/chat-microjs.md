---
title: chat Micro.js
servicename: chat
tags: microjs
---
Real time messaging API which enables Chat services to be embedded anywhere

# Chat Service

The Chat service is a programmable instant messaging API service which can be used in any application to immediately create conversations. 

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
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/chat/Chat/Connect",
        "micro",
        Schema related to #/components/requestBodies/ChatConnectRequest not found,
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Chat History
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
HistoryRequest contains the id of the chat we want the history for. This RPC will return all 
 historical messages, however in a real life application we'd introduce some form of pagination
 here, only loading the older messages when required.
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/chat/Chat/History",
        "micro",
        {
  "chat_id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Chat New
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
NewRequest contains the infromation needed to create a new chat
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/chat/Chat/New",
        "micro",
        {},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Chat Send
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
SendRequest contains a single message to send to a chat
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/chat/Chat/Send",
        "micro",
        {
  "chat_id": "id of the chat the message is being sent to / from",
  "client_id": "a client side id, should be validated by the server to make the request retry safe",
  "subject": "subject of the message",
  "text": "text of the message",
  "user_id": "id of the user who sent the message"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


