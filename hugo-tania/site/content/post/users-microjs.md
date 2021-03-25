---
title: users Micro.js
servicename: users
labels: 
- Micro.js
- Backend
---

## Micro.js


### Users Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/Create",
        "micro",
        {
          "email": "string",
          "first_name": "string",
          "last_name": "string",
          "password": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Users Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/Delete",
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


### Users List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/List",
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


### Users Login
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/Login",
        "micro",
        {
          "email": "string",
          "password": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Users Logout
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/Logout",
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


### Users Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/Read",
        "micro",
        {
          "ids": [
                    "string"
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


### Users ReadByEmail
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/ReadByEmail",
        "micro",
        {
          "emails": [
                    "string"
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


### Users Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/Update",
        "micro",
        {
          "email": {},
          "first_name": {},
          "id": "string",
          "last_name": {},
          "password": {}
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Users Validate
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/Validate",
        "micro",
        {
          "token": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


