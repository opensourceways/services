---
title: tags
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

```shell
> curl 'https://api.m3o.com/tags/Tags/Add' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "resource_created": 1,
  "resource_id": "string",
  "title": "string",
  "type": "string"
};
# Response
{}
```


### Tags List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
ListRequest: list either by resource id or type.
 Optionally filter by min or max count.
```shell
> curl 'https://api.m3o.com/tags/Tags/List' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "max_count": 1,
  "min_count": 1,
  "resource_id": "string",
  "type": "string"
};
# Response
{
  "tags": [
    {
      "count": 1,
      "description": "string",
      "slug": "string",
      "title": "string",
      "type": "Type is useful for namespacing and listing across resources,. ie. list tags for posts, customers etc."
    }
  ]
}
```


### Tags Remove
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/tags/Tags/Remove' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "resource_id": "string",
  "title": "string",
  "type": "string"
};
# Response
{}
```


### Tags Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/tags/Tags/Update' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "description": "string",
  "title": "string",
  "type": "string"
};
# Response
{}
```


