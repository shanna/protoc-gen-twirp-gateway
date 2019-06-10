# protoc-gen-twirp-gateway

*WARNING: This is alpha quality code, do not use in production.*

Provide RESTful URL structure for your Twirp RPC Go services just as you would with gRPC Go services.

Basic [google.api.HttpRule](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule)
gateway generation for Twirp RPC services.

See Google's [gRPC transcoding](https://cloud.google.com/endpoints/docs/grpc/transcoding) documentation for the gRPC
examples if you haven't used [google.api.HttpRule](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule)
before.

## Motivation

* Keep a grpc and grpc-gateway compatible IDL while developing with Twirp or on plaforms that don't support gRPC such as non GKE Google Cloud Run deployed services.
* Implement REST, Twirp and gRPC interface endpoints all using the same IDL and concrete service implementation.

## TODO

* Tests.
* RAML YAML.
* OpenAPI (Swagger) YAML.
