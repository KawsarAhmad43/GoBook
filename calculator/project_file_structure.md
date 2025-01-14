# Project folder structure
```go
calculator/
├── go.mod
├── go.sum
├── main.go
├── pkg/
│   └── utils.go
├── cmd/
│   └── server.go
├── internal/
│   └── config/
│       └── config.go
├── README.md
└── .gitignore
```

# Folder explanation
```go
calculator/
├── go.mod            # Module definition
├── go.sum            # Dependency checksums (auto-generated)
├── main.go           # Entry point of the application
├── pkg/              # Reusable packages (optional)
│   └── utils.go      # Utility functions or helper packages
├── cmd/              # Subcommands for your application (optional)
│   └── server.go     # Example command
├── internal/         # Private packages (optional)
│   └── config/       # Configuration handling
│       └── config.go
├── README.md         # Project documentation
└── .gitignore        # Git ignore rules 
```