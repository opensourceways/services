---
title: threads Micro.js
servicename: threads
labels: 
- Micro.js
---

## Micro.js


### Threads CreateConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/CreateConversation",
        "micro",
        {
          "group_id": "string",
          "topic": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads CreateMessage
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/CreateMessage",
        "micro",
        {
          "author_id": "string",
          "conversation_id": "string",
          "id": "string",
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


### Threads DeleteConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/DeleteConversation",
        "micro",
        {
          "id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads ListConversations
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/ListConversations",
        "micro",
        {
          "group_id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads ListMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/ListMessages",
        "micro",
        {
          "conversation_id": "string",
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


### Threads ReadConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/ReadConversation",
        "micro",
        {
          "group_id": {},
          "id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads RecentMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/RecentMessages",
        "micro",
        {
          "conversation_ids": [
                    "string"
          ],
          "limit_per_conversation": {}
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads UpdateConversation
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/UpdateConversation",
        "micro",
        {
          "id": "string",
          "topic": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


