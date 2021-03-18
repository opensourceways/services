---
title: codes Micro.js
servicename: codes
labels: 
- Micro.js
---

## Micro.js


### Codes Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Codes/Create",
        "micro",
        {
          "expires_at": "expiry time for the code, default 5 minutes",
          "identity": "e.g. phone number or email being verified"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Codes Verify
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Codes/Verify",
        "micro",
        {
          "code": "string",
          "identity": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


