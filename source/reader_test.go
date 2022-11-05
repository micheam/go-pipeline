package source_test

import (
	"context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/micheam/go-stream/internal/testinghelper"
	"github.com/micheam/go-stream/source"
)

const testcontent = `73980586-9b76-4550-8a32-9a9e6f253faa
09b74a47-468b-4d15-9058-629c44a7c4df
fa219346-0ccf-4fb8-82f4-105c197a86bd
7a374f2f-e322-43d0-b469-31a8b2c7445d
bf24a188-1fdd-4281-9542-604c0a3b1e43
e6d54de4-c264-4445-8d5f-8b49276261ff
22955c3a-d672-44c8-823d-97450f7b1ea3
e70ded2e-b070-4bcf-a130-150f10a75d0a
690deb26-0d63-48c0-96ac-4d326e114cf5
ac6908f3-f734-415c-a01f-d091b53a4a76`

func TestFromReader(t *testing.T) {
	ctx, cancel := WithDeadline(context.Background())
	defer cancel()

	r := strings.NewReader(testcontent)

	lines := make([]string, 0)
	for line := range source.FromReader(ctx, r) {
		lines = append(lines, line)
	}

	if diff := cmp.Diff(testcontent, strings.Join(lines, "\n")); diff != "" {
		t.Errorf("lines mismatch (-testcontent, +got):%s\n", diff)
	}
}
