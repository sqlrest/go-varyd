// Package varyd is the Go binding for the varyd wire API: the protobuf
// definitions under api/proto and the buf-generated gRPC contract under
// src/proto, shared by the vaql client, the varyd server, and any other
// consumer. Importing this package registers every contract descriptor.
package varyd

import (
	"slices"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	// Blank imports register every contract descriptor with the global
	// protobuf registry, which is this package's stated purpose and what
	// Services enumerates.
	_ "github.com/sqlrest/go-varyd/src/proto/document"
	_ "github.com/sqlrest/go-varyd/src/proto/graph"
	_ "github.com/sqlrest/go-varyd/src/proto/relational"
	_ "github.com/sqlrest/go-varyd/src/proto/spatial"
	_ "github.com/sqlrest/go-varyd/src/proto/types"
	_ "github.com/sqlrest/go-varyd/src/proto/varyd"
	_ "github.com/sqlrest/go-varyd/src/proto/vector"
)

// apiPackage is the protobuf package every contract descriptor lives in.
const apiPackage protoreflect.FullName = "varyd.api.v1"

// ServiceName is the full protobuf name of a varyd gRPC service.
type ServiceName string

// Services returns the sorted full names of every gRPC service the varyd wire
// contract registers. It is the library's contract surface: a regeneration
// that drops or renames a service changes this set and fails the contract
// test.
func Services() []ServiceName {
	var names []ServiceName
	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if fd.Package() != apiPackage {
			return true
		}
		for i, services := 0, fd.Services(); i < services.Len(); i++ {
			names = append(names, ServiceName(services.Get(i).FullName()))
		}
		return true
	})
	slices.Sort(names)
	return names
}
