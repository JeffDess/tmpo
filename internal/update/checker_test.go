package update

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		name     string
		current  string
		latest   string
		expected int
	}{
		{
			name:     "update available - patch version",
			current:  "1.2.3",
			latest:   "1.2.4",
			expected: -1,
		},
		{
			name:     "update available - minor version",
			current:  "1.2.3",
			latest:   "1.3.0",
			expected: -1,
		},
		{
			name:     "update available - major version",
			current:  "1.2.3",
			latest:   "2.0.0",
			expected: -1,
		},
		{
			name:     "same version",
			current:  "1.2.3",
			latest:   "1.2.3",
			expected: 0,
		},
		{
			name:     "ahead of latest (dev build)",
			current:  "1.3.0",
			latest:   "1.2.3",
			expected: 1,
		},
		{
			name:     "with v prefix - current",
			current:  "v1.2.3",
			latest:   "1.2.4",
			expected: -1,
		},
		{
			name:     "with v prefix - latest",
			current:  "1.2.3",
			latest:   "v1.2.4",
			expected: -1,
		},
		{
			name:     "with v prefix - both",
			current:  "v1.2.3",
			latest:   "v1.2.4",
			expected: -1,
		},
		{
			name:     "large version numbers",
			current:  "1.20.3",
			latest:   "1.21.0",
			expected: -1,
		},
		{
			name:     "different lengths - current shorter",
			current:  "1.2",
			latest:   "1.2.1",
			expected: -1,
		},
		{
			name:     "different lengths - latest shorter",
			current:  "1.2.1",
			latest:   "1.2",
			expected: 1,
		},
		{
			name:     "prerelease is less than stable",
			current:  "1.0.0-rc.1",
			latest:   "1.0.0",
			expected: -1,
		},
		{
			name:     "stable is greater than prerelease",
			current:  "1.0.0",
			latest:   "1.0.0-rc.1",
			expected: 1,
		},
		{
			name:     "prerelease alpha less than beta",
			current:  "1.0.0-alpha",
			latest:   "1.0.0-beta",
			expected: -1,
		},
		{
			name:     "prerelease beta less than rc",
			current:  "1.0.0-beta",
			latest:   "1.0.0-rc.1",
			expected: -1,
		},
		{
			name:     "same prerelease version",
			current:  "1.0.0-rc.1",
			latest:   "1.0.0-rc.1",
			expected: 0,
		},
		{
			name:     "prerelease with v prefix",
			current:  "v1.0.0-rc.1",
			latest:   "v1.0.0",
			expected: -1,
		},
		{
			name:     "older prerelease than newer stable",
			current:  "1.0.0-beta",
			latest:   "1.1.0",
			expected: -1,
		},
		{
			name:     "newer stable than older prerelease",
			current:  "1.1.0",
			latest:   "1.0.0-rc.1",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareVersions(tt.current, tt.latest)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsConnectedToInternet(t *testing.T) {
	// This test will pass if there's internet connectivity
	// It's more of an integration test than a unit test
	result := IsConnectedToInternet()
	// We can't assert true/false since it depends on actual connectivity
	// Just verify it returns without panicking
	t.Logf("Internet connectivity: %v", result)
}
