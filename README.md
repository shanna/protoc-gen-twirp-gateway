# protoc-gen-twirp-gateway

**WARNING: This is alpha quality code, do not use in production.**

Provide a RESTful gateway for your [Twirp RPC](https://github.com/twitchtv/twirp) Go services just as you would with
[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway). A basic implementation of
[google.api.HttpRule](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule)
gateway generation for [Twirp RPC](https://github.com/twitchtv/twirp) Go services.

## Usage

See Google's [gRPC transcoding](https://cloud.google.com/endpoints/docs/grpc/transcoding) documentation for the gRPC
examples if you haven't used [google.api.HttpRule](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule)
before.

## Motivation

* Keep a grpc and grpc-gateway compatible IDL while developing with twirp.
* Provide RESTful endpoints on serverless plaforms that don't support gRPC such as non GKE Google Cloud Run deployed services.
* Implement REST, Twirp and gRPC interface endpoints all using the same IDL and concrete service implementation.

## TODO

* Tests.
* RAML YAML.
* OpenAPI (Swagger) YAML.
