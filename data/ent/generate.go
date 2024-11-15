package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ../schema --target . --feature schema/snapshot,intercept,sql/lock,sql/upsert,sql/modifier,sql/versioned-migration --template ./template
