---
title: streams
servicename: streams
labels: 
- Readme
---
# Streams Service

The streams service provides an event stream, designed for sending messages from a server to mutliple
clients connecting via Websockets. The Token RPC should be called to generate a token for each client,
the clients should then subscribe using the Subscribe RPC.

## cURL


### Streams Publish
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Schema related to #/components/requestBodies/StreamsPublishRequest not found
```shell
> curl 'https://api.m3o.com/streams/Streams/Publish' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d Schema related to #/components/requestBodies/StreamsPublishRequest not found;
# Response
{}
```


### Streams Subscribe
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/streams/Streams/Subscribe' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "token": "tokens should be provided if the user is not proving an API key on the request (e.g. in cases. where the stream is being consumed directly from the frontend via websockets). tokens can be. generated using the Token RPC",
  "topic": "topic the user wishes to subscribe to, required"
};
# Response
Schema related to #/components/responses/StreamsSubscribeResponse not found
```


### Streams Token
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/streams/Streams/Token' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "topic": "the topic the token should be restricted to, if no topic is required the token can be used to . subscribe to any topic"
};
# Response
{
  "token": "string"
}
```


