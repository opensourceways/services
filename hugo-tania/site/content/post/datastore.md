---
title: datastore
servicename: datastore
labels: 
- Readme
- Backend
---
# Datastore Service

This is the Datastore service

Generated with

```
micro new datastore
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


### Datastore Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/datastore/Datastore/Create' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "value": "JSON marshalled record to save"
};
# Response
{}
```


### Datastore CreateIndex
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/datastore/Datastore/CreateIndex' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "index": {
    "base32encode": true,
    "field_name": "Field to index on.. eg. email",
    "float32max": 1,
    "float64max": 1,
    "float_format": "string",
    "order": {
      "field_name": "Field to order on. eg. age",
      "order_type": "Type of the ordering. eg. ascending, descending, unordered"
    },
    "string_order_pad_length": 1,
    "type": "Type of index. eg. eq",
    "unique": true
  }
};
# Response
{}
```


### Datastore Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/datastore/Datastore/Delete' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "query": {
    "index": {
      "base32encode": true,
      "field_name": "Field to index on.. eg. email",
      "float32max": 1,
      "float64max": 1,
      "float_format": "string",
      "order": {
        "field_name": "Field to order on. eg. age",
        "order_type": "Type of the ordering. eg. ascending, descending, unordered"
      },
      "string_order_pad_length": 1,
      "type": "Type of index. eg. eq",
      "unique": true
    },
    "limit": 1,
    "offset": 1,
    "order": {
      "field_name": "Field to order on. eg. age",
      "order_type": "Type of the ordering. eg. ascending, descending, unordered"
    },
    "value": "string"
  }
};
# Response
{}
```


### Datastore Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/datastore/Datastore/Read' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "query": {
    "index": {
      "base32encode": true,
      "field_name": "Field to index on.. eg. email",
      "float32max": 1,
      "float64max": 1,
      "float_format": "string",
      "order": {
        "field_name": "Field to order on. eg. age",
        "order_type": "Type of the ordering. eg. ascending, descending, unordered"
      },
      "string_order_pad_length": 1,
      "type": "Type of index. eg. eq",
      "unique": true
    },
    "limit": 1,
    "offset": 1,
    "order": {
      "field_name": "Field to order on. eg. age",
      "order_type": "Type of the ordering. eg. ascending, descending, unordered"
    },
    "value": "string"
  }
};
# Response
{
  "value": "JSON marshalled record found"
}
```


### Datastore Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/datastore/Datastore/Update' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "value": "JSON marshalled record to save"
};
# Response
{}
```


