# letsGO

## Go api starter


### Ingredients

- Go
- gin ( https://github.com/gin-gonic/gin )
- mongodb ( https://www.mongodb.com/ )
- mongo-go-driver ( https://github.com/mongodb/mongo-go-driver )
- oauth2 ( https://github.com/golang/oauth2 )
- check ( https://godoc.org/gopkg.in/check.v1 )
- godotenv ( https://github.com/joho/godotenv )

### Directory Structure

- The `controllers` directory contains the core code of your application.
- The `helpers` directory contains helpers functions of your application.
- The `middlewares` directory contains middlewares of your application.
- The `routes` directory contains RESTful api routes of your application.
- The `tests` directory contains tests of your application.
- The `types` directory contains the types/structures of your application.

### Environment Configuration

letsGo uses `godotenv` for setting environment variables. The root directory of your application will contain a `.env.example` file.
copy and rename it to `.env` to set your environment variables.

You need to create a `.env.testing` file from `.env.example` for running tests.

### Setting up

- clone letsGo
- change package name in `glide.yaml` to your package name
- change the internal package (controllers, tests, helpers etc.) paths as per your requirement
- setup `.env` and `.env.testing`
- run `glide install` to install dependencies

### Run : ```go run main.go```

### Build : ```go build```

### Test : ```go test tests/main_test.go```

### Authentication

letsGo uses Go OAuth2 (https://godoc.org/golang.org/x/oauth2) for authentication.

#### Client Credential Grant
- `/api/v1/credentials` : returns `CLIENT_ID` and `CLIENT_SECRET`
- `/api/v1/token` : Send `CLIENT_ID`, `CLIENT_SECRET`, `grant_type` and `scope` to generate `access_token`
- Any route starting with `/api/v1/auth/` needs `Authorization` Header with the `Bearer access_token`



