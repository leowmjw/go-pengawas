run:
	@go run cmd/cli/main.go

gen:
	@go generate ./...

openapi:
	@oapi-codegen -package vault -generate "types,client,chi-server,spec" ./OASv3/vault/vault.yaml > ./api/vault/vault.gen.go
	 @oapi-codegen -package payments -generate "types,client,chi-server,spec" ./OASv3/payments/payments.yaml > ./api/payments/payments.gen.go
	 @oapi-codegen -package calm -generate "types,client,chi-server,spec" ./OASv3/calm/calm.yaml > ./api/calm/calm.gen.go

tools:
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest # install OpenAPI Generator
