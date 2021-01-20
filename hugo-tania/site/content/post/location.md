---
title: location
---
Real time GPS location tracking and search

# Location Service

Send, store and search real time gps point data and tracking info using the location API. 
Build powerful map rendered views with up to the second updated points on the map.

Generated with

```
micro new location
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```

## cURL


### Location Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/location/Location/Read' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{
  "entity": [
    {
      "id": "string",
      "location": [
        {
          "latitude": 1,
          "longitude": 1,
          "timestamp": 1
        }
      ],
      "type": "string"
    }
  ]
}
```


### Location Save
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/location/Location/Save' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "entity": [
    {
      "id": "string",
      "location": [
        {
          "latitude": 1,
          "longitude": 1,
          "timestamp": 1
        }
      ],
      "type": "string"
    }
  ]
};
# Response
{}
```


### Location Search
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/location/Location/Search' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "center": [
    {
      "latitude": 1,
      "longitude": 1,
      "timestamp": 1
    }
  ],
  "num_entities": 1,
  "radius": 1,
  "type": "string"
};
# Response
{
  "entities": [
    {
      "id": "string",
      "location": [
        {
          "latitude": 1,
          "longitude": 1,
          "timestamp": 1
        }
      ],
      "type": "string"
    }
  ]
}
```


