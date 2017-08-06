package main

import "testing"

type test struct {
	pkg      string
	expected map[string]bool
}

func TestList(t *testing.T) {
	tests := []test{
		test{"hash", map[string]bool{"errors": true, "internal/race": true, "io": true, "runtime": true, "runtime/internal/atomic": true, "runtime/internal/sys": true, "sync": true, "sync/atomic": true, "unsafe": true}},
		test{"errors", map[string]bool{"runtime": true, "runtime/internal/sys": true, "runtime/internal/atomic": true, "unsafe": true}},
		test{"math", map[string]bool{"runtime": true, "runtime/internal/atomic": true, "runtime/internal/sys": true, "unsafe": true}},
	}

	for _, tc := range tests {
		deps := list(tc.pkg).Deps
		if len(deps) != len(tc.expected) {
			t.Errorf("dependencies num expects %d but actual %d\n", len(tc.expected), len(deps))
		}
		for _, dep := range deps {
			if _, ok := tc.expected[dep]; !ok {
				t.Errorf("wrong dependency: %s", dep)
			}
		}
	}
}
