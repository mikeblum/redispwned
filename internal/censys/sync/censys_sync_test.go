// +build sync

package sync

import (
	"strings"
	"testing"

	"github.com/mikeblum/haveibeenredised/internal/censys"
	"github.com/mikeblum/haveibeenredised/internal/censys/search"
	asserts "github.com/stretchr/testify/assert"
)

// !! WARNING - Censys API Quotas Allow For Only 2X / Month Syncs !!
func TestCensysDataExport(t *testing.T) {
	assert := asserts.New(t)
	client := censys.NewClient()
	response, err := search.RedisQuery(client)
	assert.Nil(err)
	assert.NotNil(response)
	assert.True(strings.EqualFold("ok", response.Status))
}
