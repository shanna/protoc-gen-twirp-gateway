# protoc-gen-twirp-gateway

**WARNING: This is alpha quality code, do not use in production.**

Generate [google.api.HttpRule](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule)
defined gateways for your [Twirp RPC](https://github.com/twitchtv/twirp) Go services.

## Usage

See Google's [gRPC transcoding](https://cloud.google.com/endpoints/docs/grpc/transcoding) documentation for gRPC
examples if you haven't used [google.api.HttpRule](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule)
before.

```sh
protoc --go_out="." --twirp_out="." --twirp-gateway_out="." $proto
```

## Motivation

* Keep your [twirp](https://github.com/twitchtv/twirp) protobuf service definitions [grpc](https://grpc.io) and [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) compatible.
* Provide RESTful endpoints on serverless plaforms that don't support gRPC such as non GKE Google Cloud Run deployed services.
* Implement REST, Twirp and gRPC interface endpoints all using the same IDL and concrete service implementation.

## TODO

* Tests.
* RAML YAML.
* OpenAPI (Swagger) YAML.
