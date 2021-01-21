---
title: places Micro.js
servicename: places
tags: microjs
---
Store and search for points of interest

# Places Service

The places API stores points of interest and enables you to search for places nearby or last visited.


## Usage

Places makes use of postgres. Set the config for the database

```
micro user config set places.database "postgresql://postgres@localhost:5432/locations?sslmode=disable"
```

Run the service

```
micro run .
```

## cURL


### Places Last
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Places/Last",
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


### Places Near
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Places/Near",
        "micro",
        {
  "latitude": [
    {}
  ],
  "longitude": [
    {}
  ],
  "radius": [
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


### Places Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Places/Read",
        "micro",
        {
  "after": "string",
  "before": "string"
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
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Places/Save",
        "micro",
        {
  "places": [
    {
      "id": "string",
      "latitude": [
        {}
      ],
      "longitude": [
        {}
      ],
      "metadata": [
        {}
      ],
      "name": "string",
      "timestamp": "string"
    }
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


