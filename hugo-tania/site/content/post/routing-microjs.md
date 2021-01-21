---
title: routing Micro.js
servicename: routing
tags: microjs
---

## Micro.js


### Routing Route
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    # Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/protobuf/Routing/Route",
        "micro",
        {
          "destination": [
                    {
                              "latitude": [
                                        {}
                              ],
                              "longitude": [
                                        {}
                              ]
                    }
          ],
          "origin": [
                    {
                              "latitude": [
                                        {}
                              ],
                              "longitude": [
                                        {}
                              ]
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


