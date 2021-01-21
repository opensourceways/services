---
title: location Micro.js
servicename: location
tags: microjs
---
Real time GPS location tracking and search

# Location Service

Send, store and search real time gps point data and tracking info using the location API. 
Build powerful map rendered views with up to the second updated points on the map.

Generated with

```
micro new location
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```

## cURL


### Location Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/location/Location/Read",
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


### Location Save
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/location/Location/Save",
        "micro",
        {
  "entity": [
    {
      "id": "string",
      "location": [
        {
          "latitude": 1,
          "longitude": 1,
          "timestamp": 1
        }
      ],
      "type": "string"
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


### Location Search
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/location/Location/Search",
        "micro",
        {
  "center": [
    {
      "latitude": 1,
      "longitude": 1,
      "timestamp": 1
    }
  ],
  "num_entities": 1,
  "radius": 1,
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


