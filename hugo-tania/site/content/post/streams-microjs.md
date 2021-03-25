---
title: streams Micro.js
servicename: streams
labels: 
- Micro.js
---

## Micro.js


### Streams Publish
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Schema related to #/components/requestBodies/StreamsPublishRequest not found
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/streams/Streams/Publish",
        "micro",
        Schema related to #/components/requestBodies/StreamsPublishRequest not found,
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Streams Subscribe
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/streams/Streams/Subscribe",
        "micro",
        {
          "token": "tokens should be provided if the user is not proving an API key on the request (e.g. in cases. where the stream is being consumed directly from the frontend via websockets). tokens can be. generated using the Token RPC",
          "topic": "topic the user wishes to subscribe to, required"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Streams Token
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/streams/Streams/Token",
        "micro",
        {
          "topic": "the topic the token should be restricted to, if no topic is required the token can be used to . subscribe to any topic"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


