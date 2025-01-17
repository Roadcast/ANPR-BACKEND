//go:generate go get entgo.io/ent/cmd/ent@v0.14.1
//go:generate go run entgo.io/ent/cmd/ent generate --feature schema/snapshot,privacy,entgql --target ./internal/ent ./ent/schema
//go:generate go run -mod=mod ./cmd/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
package main
