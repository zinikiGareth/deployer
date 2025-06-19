package blob_test

import (
	"testing"

	"ziniki.org/deployer/driver/pkg/pluggable"
	"ziniki.org/deployer/testmod/internal/blob"
)

func TestTypes(t *testing.T) {
	var _ pluggable.HasMethods = &blob.Blobber{}
}
