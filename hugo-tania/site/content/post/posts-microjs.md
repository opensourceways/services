---
title: posts Micro.js
servicename: posts
tags: microjs
---

## Micro.js


### Posts Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    # Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/posts/Posts/Delete",
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


### Posts Query
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Query posts. Acts as a listing when no id or slug provided.
 Gets a single post by id or slug if any of them provided.
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    # Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/posts/Posts/Query",
        "micro",
        {
          "id": "string",
          "limit": 1,
          "offset": 1,
          "slug": "string",
          "tag": "string"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Posts Save
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    # Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/posts/Posts/Save",
        "micro",
        {
          "content": "string",
          "id": "string",
          "image": "string",
          "metadata": [
                    {}
          ],
          "slug": "string",
          "timestamp": 1,
          "title": "string"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


