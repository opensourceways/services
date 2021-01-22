---
title: feeds Micro.js
servicename: feeds
tags: 
- Micro.js
- Headless CMS
---

## Micro.js


### Feeds Entries
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/feeds/Feeds/Entries",
        "micro",
        {
          "url": "rss feed url. eg. http://a16z.com/feed/"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Feeds New
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/feeds/Feeds/New",
        "micro",
        {
          "name": "rss feed name. eg. a16z",
          "url": "rss feed url. eg. http://a16z.com/feed/"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


