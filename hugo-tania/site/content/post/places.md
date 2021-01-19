---
title: places
---
The places service stores places of interest by geolocation.

# Places Service

Generated with

```
micro new places
```

## Usage

Places makes use of postgres. Set the config for the database

```
micro user config set places.database "postgresql://postgres@localhost:5432/locations?sslmode=disable"
```

Run the service

```
micro run .
```

## cURL


### Places Last
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Places/Last' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {};
# Response
Schema related to #/components/responses/PlacesLastResponse not found
```


### Places Near
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Places/Near' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "latitude": [
    {}
  ],
  "longitude": [
    {}
  ],
  "radius": [
    {}
  ]
};
# Response
Schema related to #/components/responses/PlacesNearResponse not found
```


### Places Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Places/Read' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "after": "string",
  "before": "string"
};
# Response
Schema related to #/components/responses/PlacesReadResponse not found
```


### Places Save
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/protobuf/Places/Save' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "places": [
    {
      "id": "string",
      "latitude": [
        {}
      ],
      "longitude": [
        {}
      ],
      "metadata": [
        {}
      ],
      "name": "string",
      "timestamp": "string"
    }
  ]
};
# Response
{}
```


