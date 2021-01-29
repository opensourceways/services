---
title: datastore Micro.js
servicename: datastore
labels: 
- Micro.js
- Backend
---

## Micro.js


### Datastore Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/datastore/Datastore/Create",
        "micro",
        {
          "value": "JSON marshalled record to save"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Datastore CreateIndex
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/datastore/Datastore/CreateIndex",
        "micro",
        {
          "index": {
                    "base32encode": true,
                    "field_name": "Field to index on.. eg. email",
                    "float32max": 1,
                    "float64max": 1,
                    "float_format": "string",
                    "order": {
                              "field_name": "Field to order on. eg. age",
                              "order_type": "Type of the ordering. eg. ascending, descending, unordered"
                    },
                    "string_order_pad_length": 1,
                    "type": "Type of index. eg. eq",
                    "unique": true
          }
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Datastore Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/datastore/Datastore/Delete",
        "micro",
        {
          "query": {
                    "index": {
                              "base32encode": true,
                              "field_name": "Field to index on.. eg. email",
                              "float32max": 1,
                              "float64max": 1,
                              "float_format": "string",
                              "order": {
                                        "field_name": "Field to order on. eg. age",
                                        "order_type": "Type of the ordering. eg. ascending, descending, unordered"
                              },
                              "string_order_pad_length": 1,
                              "type": "Type of index. eg. eq",
                              "unique": true
                    },
                    "limit": 1,
                    "offset": 1,
                    "order": {
                              "field_name": "Field to order on. eg. age",
                              "order_type": "Type of the ordering. eg. ascending, descending, unordered"
                    },
                    "value": "string"
          }
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Datastore Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/datastore/Datastore/Read",
        "micro",
        {
          "query": {
                    "index": {
                              "base32encode": true,
                              "field_name": "Field to index on.. eg. email",
                              "float32max": 1,
                              "float64max": 1,
                              "float_format": "string",
                              "order": {
                                        "field_name": "Field to order on. eg. age",
                                        "order_type": "Type of the ordering. eg. ascending, descending, unordered"
                              },
                              "string_order_pad_length": 1,
                              "type": "Type of index. eg. eq",
                              "unique": true
                    },
                    "limit": 1,
                    "offset": 1,
                    "order": {
                              "field_name": "Field to order on. eg. age",
                              "order_type": "Type of the ordering. eg. ascending, descending, unordered"
                    },
                    "value": "string"
          }
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Datastore Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/datastore/Datastore/Update",
        "micro",
        {
          "value": "JSON marshalled record to save"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


