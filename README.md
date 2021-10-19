# capstone-backend
This project is written in GO [Link to the frontend repo](https://github.com/NapatJamjan/capstone-frontend)

## Used libraries
| Name | Description |
| ---- | ------------ |
| [gin](https://github.com/gin-gonic/gin) | a HTTP web framework |
| [go-redis](https://github.com/go-redis/redis) | a type-safe Redis client |
| [gqlgen](https://github.com/99designs/gqlgen) | a library for building GraphQL servers |
| [golang-jwt](https://github.com/golang-jwt/jwt) | an implementation of JSON Web Tokens |
| [google-uuid](https://github.com/google/uuid) | a uuid generator |
| [prisma-client-go](https://github.com/prisma/prisma-client-go) | an auto-generated and fully type-safe database client |
| [viper](https://github.com/spf13/viper) | a complete configuration solution |

## How to run the project
1. Clone the project to your machine
```console
git clone https://github.com/Nuttawut503/capstone-backend && cd capstone-backend
```
2. Copy `.env.example` file then rename the copied to `.env`
3. Install postgresql and redis via docker (Skip if you already have)
```console
chmod +x docker_setup.sh && ./docker_setup.sh
```
4. Edit `.env` file to match your machine environment (Skip if you didn't run the above command)
5. Install required dependencies in the project
```console
go install
```
6. Migrate your database
```console
go run github.com/prisma/prisma-client-go migrate dev
```
7. Start the project
```console
go run .
```
