following guide:
https://www.howtographql.com/graphql-go/1-getting-started/

Added the first schema definitions, but when running "go run github.com/99designs/gqlgen generate" the following errors pop up:

../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.1/main.go:20:2: missing go.sum entry for module providing package github.com/urfave/cli/v2 (imported by github.com/99designs/gqlgen); to add:
	go get github.com/99designs/gqlgen@v0.17.1
../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.1/internal/imports/prune.go:15:2: missing go.sum entry for module providing package golang.org/x/tools/go/ast/astutil (imported by github.com/99designs/gqlgen/internal/imports); to add:
	go get github.com/99designs/gqlgen/internal/imports@v0.17.1
../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.1/internal/code/packages.go:11:2: missing go.sum entry for module providing package golang.org/x/tools/go/packages (imported by github.com/99designs/gqlgen/codegen/config); to add:
	go get github.com/99designs/gqlgen/codegen/config@v0.17.1
../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.1/internal/imports/prune.go:16:2: missing go.sum entry for module providing package golang.org/x/tools/imports (imported by github.com/99designs/gqlgen/internal/imports); to add:
	go get github.com/99designs/gqlgen/internal/imports@v0.17.1
