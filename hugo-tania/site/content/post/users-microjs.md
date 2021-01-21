---
title: users Micro.js
servicename: users
tags: microjs
---
A user service for storing accounts and simple auth.

# Users Service

The users service provides user management and authentication so you can easily add them to your own apps 
without having to build the entire thing from scratch.

## Getting started

```
micro run github.com/micro/services/users
```

## Usage

User server implements the following RPC Methods

Users
- Create
- Read
- Update
- Delete
- Search
- UpdatePassword
- Login
- Logout
- ReadSession


### Create

```shell
micro call users Users.Create '{"id": "ff3c06de-9e43-41c7-9bab-578f6b4ad32b", "username": "asim", "email": "asim@example.com", "password": "password1"}'
```

### Read

```shell
micro call users Users.Read '{"id": "ff3c06de-9e43-41c7-9bab-578f6b4ad32b"}'
```

### Update

```shell
micro call users Users.Update '{"id": "ff3c06de-9e43-41c7-9bab-578f6b4ad32b", "username": "asim", "email": "asim+update@example.com"}'
```

### Update Password

```shell
micro call users Users.UpdatePassword '{"userId": "ff3c06de-9e43-41c7-9bab-578f6b4ad32b", "oldPassword": "password1", "newPassword": "newpassword1", "confirmPassword": "newpassword1" }'
```

### Delete

```shell
micro call users Users.Delete '{"id": "ff3c06de-9e43-41c7-9bab-578f6b4ad32b"}'
```

### Login

```shell
micro call users Users.Login '{"username": "asim", "password": "password1"}'
```

### Read Session

```shell
micro call users Users.ReadSession '{"sessionId": "sr7UEBmIMg5hYOgiljnhrd4XLsnalNewBV9KzpZ9aD8w37b3jRmEujGtKGcGlXPg1yYoSHR3RLy66ugglw0tofTNGm57NrNYUHsFxfwuGC6pvCn8BecB7aEF6UxTyVFq"}'
```

### Logout

```shell
micro call users Users.Logout '{"sessionId": "sr7UEBmIMg5hYOgiljnhrd4XLsnalNewBV9KzpZ9aD8w37b3jRmEujGtKGcGlXPg1yYoSHR3RLy66ugglw0tofTNGm57NrNYUHsFxfwuGC6pvCn8BecB7aEF6UxTyVFq"}'
```

## cURL


### Users Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
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


