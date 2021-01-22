---
title: routing
servicename: routing
tags: 
- Readme
- Logistics
---
Point to point routing directions

# Routing Service

The routing service provides point to point directions

## cURL


### Routing Route
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Routing/Route' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "destination": [
    {
      "latitude": [
        {}
      ],
      "longitude": [
        {}
      ]
    }
  ],
  "origin": [
    {
      "latitude": [
        {}
      ],
      "longitude": [
        {}
      ]
    }
  ]
};
# Response
{
  "waypoints": [
    {
      "latitude": [
        {}
      ],
      "longitude": [
        {}
      ]
    }
  ]
}
```


