module github.com/opentdf/platform/lib/fory

go 1.25.0

require (
	github.com/apache/fory/go/fory v0.15.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.30.0 // indirect
	golang.org/x/sync v0.18.0 // indirect
	golang.org/x/tools v0.39.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// NOTE: Apache Fory Go module is not yet published to the Go module proxy.
// To build locally, clone https://github.com/apache/fory and add a replace directive:
//   replace github.com/apache/fory/go/fory => /path/to/apache-fory/go/fory
// This will be removed once the module is officially published.
