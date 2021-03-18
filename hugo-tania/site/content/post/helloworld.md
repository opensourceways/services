---
title: helloworld
servicename: helloworld
labels: 
- Readme
- Backend
---
# Helloworld Service

This is the Helloworld service

## Overview

An example of how to write a simple helloworld service. This can also be generated using `micro new helloworld`.

## Usage

```
# run the server
micro server

# run the service
micro run github.com/micro/services/helloworld

## call the service
micro call helloworld Helloworld.Call '{"name": "Alice"}'
```

## cURL


### Helloworld Call
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/helloworld/Helloworld/Call' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "name": "string"
};
# Response
{
  "msg": "string"
}
```


