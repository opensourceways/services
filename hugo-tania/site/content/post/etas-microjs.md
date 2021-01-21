---
title: etas Micro.js
servicename: etas
tags: microjs
---

## Micro.js


### ETAs Calculate
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Schema related to #/components/requestBodies/ETAsCalculateRequest not found
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    # Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/ETAs/Calculate",
        "micro",
        Schema related to #/components/requestBodies/ETAsCalculateRequest not found,
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


