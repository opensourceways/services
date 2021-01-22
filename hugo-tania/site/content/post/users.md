---
title: users
servicename: users
labels: 
- Readme
- Backend
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

```shell
> curl 'https://api.m3o.com/users/Users/Create' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "id": "uuid",
  "password": "string",
  "username": "alphanumeric user or org"
};
# Response
{}
```


### Users Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Delete' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{}
```


### Users Login
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Login' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "password": "string",
  "username": "string"
};
# Response
{
  "session": [
    {
      "created": 1,
      "email": "string",
      "expires": 1,
      "id": "string",
      "username": "string"
    }
  ]
}
```


### Users Logout
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Logout' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "session_id": "string"
};
# Response
{}
```


### Users Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Read' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{
  "user": [
    {
      "created": 1,
      "email": "string",
      "id": "uuid",
      "updated": 1,
      "username": "alphanumeric user or org"
    }
  ]
}
```


### Users ReadSession
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/ReadSession' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "session_id": "string"
};
# Response
{
  "session": [
    {
      "created": 1,
      "email": "string",
      "expires": 1,
      "id": "string",
      "username": "string"
    }
  ]
}
```


### Users Search
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Search' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "limit": 1,
  "offset": 1,
  "username": "string"
};
# Response
{
  "users": [
    {
      "created": 1,
      "email": "string",
      "id": "uuid",
      "updated": 1,
      "username": "alphanumeric user or org"
    }
  ]
}
```


### Users Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Update' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "id": "uuid",
  "username": "alphanumeric user or org"
};
# Response
{}
```


### Users UpdatePassword
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/UpdatePassword' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "confirm_password": "string",
  "new_password": "string",
  "old_password": "string",
  "user_id": "string"
};
# Response
{}
```


