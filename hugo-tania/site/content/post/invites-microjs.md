---
title: invites Micro.js
servicename: invites
labels: 
- Micro.js
---

## Micro.js


### Invites Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Invites/Create",
        "micro",
        {
          "email": "string",
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


### Invites Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Invites/Delete",
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


### Invites List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Invites/List",
        "micro",
        {
          "email": {},
          "group_id": {}
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Invites Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Invites/Read",
        "micro",
        {
          "code": {},
          "id": {}
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


