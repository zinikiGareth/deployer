package blob_test

import (
	"testing"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/testmod/internal/blob"
)

func TestTypes(t *testing.T) {
	var _ driverbottom.HasMethods = &blob.Blobber{}
}
