package varyd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestServicesContract pins the exact gRPC service set of the varyd.api.v1
// wire contract. A regeneration that drops, renames, or adds a service must
// change this assertion deliberately.
func TestServicesContract(t *testing.T) {
	t.Parallel()
	assert.Equal(t, []ServiceName{
		"varyd.api.v1.DocumentService",
		"varyd.api.v1.GraphService",
		"varyd.api.v1.RelationalService",
		"varyd.api.v1.SpatialService",
		"varyd.api.v1.VectorService",
	}, Services())
}
