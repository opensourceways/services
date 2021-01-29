---
title: places Micro.js
servicename: places
labels: 
- Micro.js
- Logistics
---

## Micro.js


### Places Last
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Places/Last",
        "micro",
        {
          "ids": [
                    {}
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


### Places Near
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Places/Near",
        "micro",
        {
          "latitude": {},
          "longitude": {},
          "radius": {}
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Places Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Places/Read",
        "micro",
        {
          "after": "string",
          "before": "string",
          "ids": [
                    {}
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


### Places Save
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Places/Save",
        "micro",
        {
          "places": [
                    {}
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


