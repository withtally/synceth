package parser

import (
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
)

func TestVersionConstraint(t *testing.T) {
	type match struct {
		version string
		err     bool
	}

	for _, tc := range []struct {
		pragma  string
		matches []match
	}{
		{pragma: "^0.5.16", matches: []match{
			{"0.5.16", false},
			{"0.5.104", false},
			{"0.5.15", true},
			{"0.6.0", true},
		}},
		{pragma: ">=0.4.16 <0.9.0", matches: []match{
			{"0.4.16", false},
			{"0.8.104", false},
			{"0.6.0", false},
			{"0.9.9", true},
		}},
		{pragma: "=0.7.6", matches: []match{
			{"0.7.6", false},
			{"0.5.16", true},
			{"0.7.7", true},
		}},
	} {
		c, err := NewVersionConstraint(fmt.Sprintf(`
		pragma solidity %s;

		contract Test {}
		`, tc.pragma))
		if err != nil {
			t.Fatalf("parsing pargma: %v", err)
		}

		for _, m := range tc.matches {
			v, err := semver.NewVersion(m.version)
			if err != nil {
				t.Errorf("parsing match version %s: %v", m.version, err)
				break
			}

			ok := c.Check(v)
			if !m.err && !ok {
				t.Errorf("validating constraints %s for version %s: %v", tc.pragma, m.version, err)
			} else if m.err && ok {
				t.Errorf("constraint %s validation should fail for version %s", tc.pragma, m.version)
			}

			if !m.err && !ok {
				t.Errorf("validating constraints %s for version %s: %v", tc.pragma, m.version, err)
			}
		}
	}
}
