module github.com/opentdf/platform/lib/fory

go 1.26.0

require (
	github.com/apache/fory/go/fory v0.0.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/apache/fory/go/fory => /home/actions/surfworks/apache-fory/go/fory
