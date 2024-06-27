package homepage

import (
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/server/testutil"
)

func TestNewTemplate(t *testing.T) {
	got, err := NewTemplate()
	testutil.AssertNoError(t, err)
	testutil.AssertNotNil(t, got)
}
func TestMustNewTemplate(t *testing.T) {
	got := MustNewTemplate(NewTemplate())
	testutil.AssertNotNil(t, got)
}
