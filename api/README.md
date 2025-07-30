# `/api`
- OpenAPI/Swagger specs, JSON schema files, protocol definition files.
- This folder is optional, and its purpose can vary depending on your project. It may include:

🔸 Option A: Generated Code
- gRPC: *.pb.go files generated from .proto.
- REST: Go files generated from OpenAPI specs using tools like oapi-codegen.

🔸 Option B: API Layer Abstractions
- Request/response DTO structs.
- Interface definitions for handler layers.
- API versioning (v1, v2, etc.).
- Common error types or status codes.

Examples:
- <https://github.com/kubernetes/kubernetes/tree/master/api>
- <https://github.com/moby/moby/tree/master/api>