## This project is written in GO

Used libraries
- gin (a HTTP web framework)
- go-redis (Type-safe Redis client)
- gqlgen (a library for building GraphQL servers)
- prisma-client-go (an auto-generated and fully type-safe database client)
- viper (a complete configuration solution)

How to run the project (read carefully)
1. Copy `.env.example` file then rename the copied to `.env`
2. Install postgresql and redis via docker (Skip if you already have)
```console
user@machine:~/capstone-backend$ chmod +x docker_setup.sh && ./docker_setup.sh
```
3. Edit `.env` file to match your machine environment (Skip if you didn't run the above command)
4. Install required dependencies in the project
```console
user@machine:~/capstone-backend$ go install
```
5. Migrate your database
```console
user@machine:~/capstone-backend$ go run github.com/prisma/prisma-client-go migrate dev
```
6. Start the project
```console
user@machine:~/capstone-backend$ go run .
```
