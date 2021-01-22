---
title: chat Micro.js
servicename: chat
labels: 
- Micro.js
- Communications
---

## Micro.js


### Chat Connect
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Schema related to #/components/requestBodies/ChatConnectRequest not found
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
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
    // Login is only required for endpoints doing authorization
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
    // Login is only required for endpoints doing authorization
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
    // Login is only required for endpoints doing authorization
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


