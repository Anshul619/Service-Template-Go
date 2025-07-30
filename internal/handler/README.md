# internal/handler
- Implements your server endpoints.
- Often maps directly to HTTP routes or gRPC methods.
- Usually depends on your business logic (internal/service), but not vice versa.
- Can be HTTP (net/http, chi, echo) or gRPC (grpc.Server).