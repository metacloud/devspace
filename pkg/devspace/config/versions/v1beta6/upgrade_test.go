package v1beta6

import (
	"reflect"
	"testing"

	next "github.com/devspace-cloud/devspace/pkg/devspace/config/versions/v1beta7"
	"github.com/devspace-cloud/devspace/pkg/util/log"
	"github.com/devspace-cloud/devspace/pkg/util/ptr"
	yaml "gopkg.in/yaml.v2"
)

type testCase struct {
	in       *Config
	expected *next.Config
}

func TestSimple(t *testing.T) {
	testCases := []*testCase{
		{
			in:       &Config{},
			expected: &next.Config{},
		},
		{
			in: &Config{
				Dev: &DevConfig{
					Sync: []*SyncConfig{
						&SyncConfig{},
						&SyncConfig{
							DownloadOnInitialSync: ptr.Bool(true),
						},
						&SyncConfig{
							DownloadOnInitialSync: ptr.Bool(false),
						},
					},
				},
			},
			expected: &next.Config{
				Dev: &next.DevConfig{
					Sync: []*next.SyncConfig{
						&next.SyncConfig{},
						&next.SyncConfig{
							InitialSync: next.InitialSyncStrategyPreferLocal,
						},
						&next.SyncConfig{},
					},
				},
			},
		},
		{
			in: &Config{
				Images: map[string]*ImageConfig{
					"test": &ImageConfig{
						Image: "test",
						Tag:   "ttt",
					},
					"test2": &ImageConfig{
						Image: "test2",
					},
				},
			},
			expected: &next.Config{
				Images: map[string]*next.ImageConfig{
					"test": &next.ImageConfig{
						Image: "test",
						Tags:  []string{"ttt"},
					},
					"test2": &next.ImageConfig{
						Image: "test2",
					},
				},
			},
		},
	}

	// Run test cases
	for index, testCase := range testCases {
		newConfig, err := testCase.in.Upgrade(log.Discard)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}

		isEqual := reflect.DeepEqual(newConfig, testCase.expected)
		if !isEqual {
			newConfigYaml, _ := yaml.Marshal(newConfig)
			expectedYaml, _ := yaml.Marshal(testCase.expected)

			t.Fatalf("TestCase %d: Got %s, but expected %s", index, newConfigYaml, expectedYaml)
		}
	}
}
