# Service-Template-Go

.
├── cmd/                    # Entrypoints for your application (main packages)
│   └── main.go            # `main()` lives here
├── internal/               # Private application and library code
│   ├── config/             # Configuration loading (env, files)
│   ├── handler/            # HTTP / gRPC / CLI handlers
│   ├── service/            # Business logic
│   ├── repository/         # Data access layer (DB, APIs)
│   ├── model/              # Domain models and DTOs
│   ├── middleware/         # HTTP middleware
│   └── utils/              # Utility functions/helpers
├── api/                    # OpenAPI/Protobuf definitions and generated code
│   └── openapi.yaml
├── migrations/             # Database schema migrations (e.g., SQL or tools like goose, migrate)
├── scripts/                # Shell or CI/CD scripts
├── infrastructure/         # Deployment configs (Docker, Kubernetes, etc.)
├── test/                   # Integration tests and test data
├── go.mod
├── go.sum
└── README.md

# References
- [Go Project Layout](https://github.com/golang-standards/project-layout)