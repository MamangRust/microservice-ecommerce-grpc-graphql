module github.com/MamangRust/monolith-graphql-ecommerce-pb

go 1.25.0

require (
	github.com/MamangRust/microservice-ecommerce-pkg v1.0.18
	google.golang.org/grpc v1.80.0
	google.golang.org/protobuf v1.36.11
)

replace (
	github.com/MamangRust/microservice-ecommerce-pkg => ../pkg
	github.com/MamangRust/microservice-ecommerce-shared => ../shared
)
