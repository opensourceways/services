---
title: seen Micro.js
servicename: seen
labels: 
- Micro.js
---

## Micro.js


### Seen Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/seen/Seen/Read",
        "micro",
        {
          "resource_ids": [
                    "string"
          ],
          "resource_type": "string",
          "user_id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Seen Set
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/seen/Seen/Set",
        "micro",
        {
          "resource_id": "string",
          "resource_type": "string",
          "timestamp": "string",
          "user_id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Seen Unset
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/seen/Seen/Unset",
        "micro",
        {
          "resource_id": "string",
          "resource_type": "string",
          "user_id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


