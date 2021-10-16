## This project is written in GO

Used libraries
- gin (a HTTP web framework)
- go-redis (Type-safe Redis client)
- gqlgen (a library for building GraphQL servers)
- prisma-client-go (an auto-generated and fully type-safe database client)
- viper (a complete configuration solution)

How to run the project (read carefully)
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
