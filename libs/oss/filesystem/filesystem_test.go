package filesystem

import (
	"testing"

	"GoTenancy/libs/oss/tests"
)

func TestAll(t *testing.T) {
	fileSystem := New("/tmp")
	tests.TestAll(fileSystem, t)
}
