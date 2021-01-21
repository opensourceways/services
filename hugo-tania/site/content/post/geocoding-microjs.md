---
title: geocoding Micro.js
servicename: geocoding
tags: microjs
---
Geocode an address to gps coordinates and the reverse.

# Geocoding service

The geocoding service provides address to lat lng geocoding as well as the reverse. Useful for building mapping or location 
based services.

## cURL


### Geocoding Geocode
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Schema related to #/components/requestBodies/GeocodingGeocodeRequest not found
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Geocoding/Geocode",
        "micro",
        Schema related to #/components/requestBodies/GeocodingGeocodeRequest not found,
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Geocoding Reverse
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Schema related to #/components/requestBodies/GeocodingReverseRequest not found
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Geocoding/Reverse",
        "micro",
        Schema related to #/components/requestBodies/GeocodingReverseRequest not found,
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


