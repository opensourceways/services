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
          "id": "uuid",
          "password": "string",
          "username": "alphanumeric user or org"
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
          "password": "string",
          "username": "string"
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
          "session_id": "string"
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


### Users ReadSession
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/ReadSession",
        "micro",
        {
          "session_id": "string"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Users Search
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/Search",
        "micro",
        {
          "email": "string",
          "limit": 1,
          "offset": 1,
          "username": "string"
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
          "email": "string",
          "id": "uuid",
          "username": "alphanumeric user or org"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Users UpdatePassword
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/users/Users/UpdatePassword",
        "micro",
        {
          "confirm_password": "string",
          "new_password": "string",
          "old_password": "string",
          "user_id": "string"
        },
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


