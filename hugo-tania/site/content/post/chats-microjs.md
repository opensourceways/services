---
title: chats Micro.js
servicename: chats
labels: 
- Micro.js
---

## Micro.js


### Chats CreateChat
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Chats/CreateChat",
        "micro",
        {
          "user_ids": [
                    "string"
          ]
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Chats CreateMessage
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Chats/CreateMessage",
        "micro",
        {
          "author_id": "string",
          "chat_id": "string",
          "text": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Chats ListMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Chats/ListMessages",
        "micro",
        {
          "chat_id": "string",
          "limit": {},
          "sent_before": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


