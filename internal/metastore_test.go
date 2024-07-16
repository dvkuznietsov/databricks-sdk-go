package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/stretchr/testify/assert"
)

func TestUcAccMetastoreAssignments(t *testing.T) {
	ctx, a := ucacctTest(t)
	ws, err := a.MetastoreAssignments.ListAll(ctx, catalog.ListAccountMetastoreAssignmentsRequest{
		MetastoreId: GetEnvOrSkipTest(t, "TEST_METASTORE_ID"),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, ws)
}
