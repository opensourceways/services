---
title: geocoding
servicename: geocoding
labels: 
- Readme
- Logistics
---
Geocode an address to gps coordinates and the reverse.

# Geocoding service

The geocoding service provides address to lat lng geocoding as well as the reverse. Useful for building mapping or location 
based services.

## cURL


### Geocoding Geocode
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Schema related to #/components/requestBodies/GeocodingGeocodeRequest not found
```shell
> curl 'https://api.m3o.com/geocoding/Geocoding/Geocode' \
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
> curl 'https://api.m3o.com/geocoding/Geocoding/Reverse' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d Schema related to #/components/requestBodies/GeocodingReverseRequest not found;
# Response
Schema related to #/components/responses/GeocodingReverseResponse not found
```


