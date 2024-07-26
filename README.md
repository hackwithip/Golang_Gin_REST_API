### REST

```
cms-app/
│
├── cmd/                  # Main application entry points
│   ├── server/           # Main server entry point
│   │   └── main.go       # Server main file
│   └── worker/           # Temporal worker entry point
│       └── main.go       # Worker main file
│
├── configs/              # Configuration files
│   ├── app.yaml          # Application config (e.g., ports, keys)
│   └── db.yaml           # Database config (e.g., credentials)
│
├── pkg/                  # Core packages
│   ├── api/              # API routes and handlers
│   │   ├── middleware/   # Middleware (e.g., auth, logging)
│   │   ├── v1/           # API version 1
│   │   │   ├── auth/     # Auth-related handlers
│   │   │   ├── category/ # Category handlers
│   │   │   ├── author/   # Author handlers
│   │   │   └── blog/     # Blog handlers
│   │   └── router.go     # Router setup
│   │
│   ├── models/           # Database models
│   │   ├── category.go   # Category model
│   │   ├── author.go     # Author model
│   │   ├── blog.go       # Blog model
│   │   └── user.go       # User model
│   │
│   ├── repositories/     # Database interaction
│   │   ├── category.go   # Category repository
│   │   ├── author.go     # Author repository
│   │   ├── blog.go       # Blog repository
│   │   └── user.go       # User repository
│   │
│   ├── services/         # Business logic
│   │   ├── auth.go       # Authentication service
│   │   ├── category.go   # Category service
│   │   ├── author.go     # Author service
│   │   ├── blog.go       # Blog service
│   │   └── email.go      # Email service (using Temporal)
│   │
│   ├── temporal/         # Temporal workflows and activities
│   │   ├── workflows.go  # Workflow definitions
│   │   └── activities.go # Activity definitions
│   │
│   ├── utils/            # Utility functions and helpers
│   │   ├── config.go     # Config loader
│   │   └── logger.go     # Logging utility
│   │
│   └── docs/             # Documentation (Swagger, etc.)
│       └── swagger.yaml  # API documentation
│
├── static/               # Static files (images, CSS, JavaScript)
│
├── scripts/              # Deployment and setup scripts
│   ├── setup.sh          # Setup script for local environment
│   └── deploy.sh         # Deployment script
│
├── Dockerfile            # Dockerfile for containerizing the app
├── docker-compose.yml    # Docker Compose file for local setup
├── go.mod                # Go modules dependencies
├── go.sum                # Go modules checksum file
└── README.md             # Project overview and setup instructions
```

### KEYCLOAK

- Signup : -> Endpont : admin token -> 
- Keycloak : signup upser => keycloak entry
- Login -> table -> keycloak -> Token -> User