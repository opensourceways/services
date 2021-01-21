---
title: groups Micro.js
servicename: groups
tags: microjs
---
# Groups Service

The group serivce is a basic CRUD service for groups. You can use it to create groups, add members and lookup which groups a user is a member of.

Example usage:

```bash
$ micro groups create --name=Micro
{
	"group": {
		"id": "e35562c9-b6f6-459a-b52d-7e6159465fd6",
		"name": "Micro"
	}
}
$ micro groups addMember --group_id=e35562c9-b6f6-459a-b52d-7e6159465fd6 --member_id=Asim
{}
$ micro groups list --member_id=Asim
{
	"groups": [
		{
			"id": "e35562c9-b6f6-459a-b52d-7e6159465fd6",
			"name": "Micro",
			"member_ids": [
				"Asim"
			]
		}
	]
}
$ micro groups list --member_id=Boris
{}
```

## cURL


### Groups AddMember
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/groups/Groups/AddMember",
        "micro",
        {
  "group_id": "string",
  "member_id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Groups Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/groups/Groups/Create",
        "micro",
        {
  "name": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Groups Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/groups/Groups/Delete",
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


### Groups List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/groups/Groups/List",
        "micro",
        {
  "member_id": "passing a member id will restrict the groups to that which the member is part of"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Groups Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/groups/Groups/Read",
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


### Groups RemoveMember
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/groups/Groups/RemoveMember",
        "micro",
        {
  "group_id": "string",
  "member_id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Groups Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/groups/Groups/Update",
        "micro",
        {
  "id": "string",
  "name": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


