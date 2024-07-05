```
go-ngodeyuk-core-project/
├── cmd/
│   └── ngodeyuk/
│       └── main.go
├── pkg/
│   └── utils/
│       └── hello.go
├── internal/
│   └── hello/
│       └── handlers/
│       └── router/
│       └── services/
├── scripts/
│   ├── run.sh
│   └── test.sh
├── tests/
│   └── hello_test.go
```

#### Folder Structure Explanation:
1. `cmd/` :
   - This directory is typically used to store the main files of the application. Each subdirectory under `cmd/` represents a different entry point for the application.
   - `ngodeyuk/` : This subdirectory contains the main.go file, which is the main entry point of the application.
   - `main.go` : This file is the main file used to run the application.
2. `pkg/` :
   - This directory is used to store common packages that can be reused by the application or other projects.
   - `utils/` : This subdirectory contains utility or helper functions that can be used throughout the application.
3. `internal/` :
   - This directory is used to store code that is internal to the project and should not be accessible by other projects outside this repository.
   - `hello/` : A subdirectory that may be related to the hello feature or module within the application.
      - `handlers/` : A subdirectory for storing HTTP handlers or other request handlers.
      - `router/` : A subdirectory for storing routing configurations or the application's router.
      - `services/` : A subdirectory for storing business logic or service layer of the application.
4. `scripts/` :
   - This directory is used to store shell scripts that assist in automating tasks such as running or testing the application.
5. `tests` : 
   - This directory is used to store unit test files or other test files.
