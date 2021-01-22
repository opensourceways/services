---
title: tags Micro.js
servicename: tags
tags: 
- Micro.js
- Headless CMS
---

## Micro.js


### Tags Add
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/tags/Tags/Add",
        "micro",
        {
          "resource_created": 1,
          "resource_id": "string",
          "title": "string",
          "type": "string"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Tags List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
ListRequest: list either by resource id or type.
 Optionally filter by min or max count.
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/tags/Tags/List",
        "micro",
        {
          "max_count": 1,
          "min_count": 1,
          "resource_id": "string",
          "type": "string"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Tags Remove
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/tags/Tags/Remove",
        "micro",
        {
          "resource_id": "string",
          "title": "string",
          "type": "string"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Tags Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/tags/Tags/Update",
        "micro",
        {
          "description": "string",
          "title": "string",
          "type": "string"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


