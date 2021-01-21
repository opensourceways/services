---
title: tags Micro.js
servicename: tags
tags: microjs
---
Tag any resource by savin a tag associated with their ID in the tag service.

# Tag Service

## Query tags

```
micro tags list --type=post-tag
```

Generated with

```
micro new --namespace=go.micro --type=service tag
```

## Getting Started

- [Tag Service](#tag-service)
  - [Query tags](#query-tags)
  - [Getting Started](#getting-started)
  - [Configuration](#configuration)
  - [Dependencies](#dependencies)
  - [Usage](#usage)

## Configuration

- FQDN: go.micro.service.tag
- Type: service
- Alias: tag

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./tag-service
```

Build a docker image
```
make docker
```
## cURL


### Tags Add
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/tags/Tags/Add",
        "micro",
        {
  "resource_created": 1,
  "resource_id": "string",
  "title": "string",
  "type": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Tags List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
ListRequest: list either by resource id or type.
 Optionally filter by min or max count.
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/tags/Tags/List",
        "micro",
        {
  "max_count": 1,
  "min_count": 1,
  "resource_id": "string",
  "type": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Tags Remove
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/tags/Tags/Remove",
        "micro",
        {
  "resource_id": "string",
  "title": "string",
  "type": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Tags Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    Micro.requireLogin(function () {
      Micro.post(
        "/tags/Tags/Update",
        "micro",
        {
  "description": "string",
  "title": "string",
  "type": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


