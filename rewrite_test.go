package rewrite

import (
	"bytes"
	"testing"

	dmp "github.com/sergi/go-diff/diffmatchpatch"
)

type rewriteTestCase struct {
	in, out []byte
}

type stringTestCase struct {
	in, out string
}

func stringTestCases(in []stringTestCase) (cases []rewriteTestCase) {
	for _, c := range in {
		cases = append(cases, rewriteTestCase{
			in:  []byte(c.in),
			out: []byte(c.out),
		})
	}
	return
}

func testRewriteCases(t *testing.T, rw Rewriter, cases []rewriteTestCase) {
	for i, c := range cases {
		got := rw.Rewrite(c.in)
		if !bytes.Equal(got, c.out) {
			dmp := dmp.New()
			diffs := dmp.DiffMain(string(c.out), string(got), true)
			if len(diffs) == 0 {
				t.Logf("case %d bytes were unequal but computed no difference between results", len(diffs))
				continue
			}

			// t.Errorf("case %d mismatch:\n%s", i, dmp.DiffPrettyHtml(diffs))
			// if len(c.out) < 50 {
			t.Errorf("case %d mismatch: expected: %s, got: %s", i, string(c.out), string(got))
			// }
		}
	}
}
