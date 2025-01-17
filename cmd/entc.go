package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"log"
)

func main() {
	// Create the EntGQL extension
	log.Printf("Creating EntGQL extension")
	ex, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithWhereFilters(true),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./internal/ent/auto-schema.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	log.Printf("Generating Ent code")
	opts := []entc.Option{
		entc.Extensions(ex),
		entc.FeatureNames("privacy", "schema/snapshot", "sql/upsert"),
	}

	// Generate Ent code from the root schema directory (it will pick subdirectories automatically)
	if err := entc.Generate("./ent/schema", &gen.Config{
		Target:  "./internal/ent", // Output directory for the generated code
		Package: "go-ent-project/internal/ent",
		// Package name for the generated code
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
