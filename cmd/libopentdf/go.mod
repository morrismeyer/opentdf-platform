module github.com/opentdf/platform/cmd/libopentdf

go 1.26.0

require github.com/opentdf/platform/lib/fory v0.0.0

require github.com/apache/fory/go/fory v0.0.0 // indirect

replace github.com/opentdf/platform/lib/fory => ../../lib/fory

replace github.com/apache/fory/go/fory => /home/actions/surfworks/apache-fory/go/fory
