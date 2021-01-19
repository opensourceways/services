---
title: geocoding
---
The geocoding service provides address to lat lng geocoding as well as the reverse.

# Geocoding service

Examples coming soon. Check the [proto](https://github.com/micro/services/blob/master/geocoding/proto/geocoding.proto) for more details.
## cURL


### Geocoding Geocode
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Schema related to #/components/requestBodies/GeocodingGeocodeRequest not found
```shell
> curl 'https://api.m3o.com/protobuf/Geocoding/Geocode' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d Schema related to #/components/requestBodies/GeocodingGeocodeRequest not found;
# Response
Schema related to #/components/responses/GeocodingGeocodeResponse not found
```


### Geocoding Reverse
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Schema related to #/components/requestBodies/GeocodingReverseRequest not found
```shell
> curl 'https://api.m3o.com/protobuf/Geocoding/Reverse' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d Schema related to #/components/requestBodies/GeocodingReverseRequest not found;
# Response
Schema related to #/components/responses/GeocodingReverseResponse not found
```


