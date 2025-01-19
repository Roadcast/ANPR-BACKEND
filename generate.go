//go:generate go get -u entgo.io/ent/cmd/ent
//go:generate go run entgo.io/ent/cmd/ent generate --feature schema/snapshot,privacy,entgql --target ./internal/ent ./ent/schema
//go:generate go run -mod=mod ./cmd/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
package main
