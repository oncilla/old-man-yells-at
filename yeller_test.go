// Copyright 2021 oncilla
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yeller_test

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"testing"

	yeller "github.com/oncilla/old-man-yells-at"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// update is a cmd line flag that enables golden file updates. To update the
// golden files simply run 'go test -update ./...'.
var update = flag.Bool("update", false, "set to true to regenerate golden files")

func TestYellAt(t *testing.T) {
	tests := map[string]struct {
		Input  string
		Output string
	}{
		"bazel": {
			Input:  "testdata/bazel.png",
			Output: "testdata/old-man-yells-at-bazel.png",
		},
		"google": {
			Input:  "testdata/google.png",
			Output: "testdata/old-man-yells-at-google.png",
		},
		"slack": {
			Input:  "testdata/slack.png",
			Output: "testdata/old-man-yells-at-slack.png",
		},
		"vscode": {
			Input:  "testdata/vscode.png",
			Output: "testdata/old-man-yells-at-vscode.png",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			raw, err := ioutil.ReadFile(tc.Input)
			require.NoError(t, err)
			target, err := png.Decode(bytes.NewReader(raw))
			require.NoError(t, err)
			yelledAt := yeller.YellAt(target)

			var buf bytes.Buffer
			err = png.Encode(&buf, yelledAt)
			require.NoError(t, err)

			if *update {
				err := ioutil.WriteFile(tc.Output, buf.Bytes(), 0666)
				require.NoError(t, err)
			}

			expected, err := ioutil.ReadFile(tc.Output)
			require.NoError(t, err)
			assert.Equal(t, expected, buf.Bytes())
		})
	}
}

// Hack to do embedding, should be replaced!
func TestAbeBase64(t *testing.T) {
	if !*update {
		return
	}
	raw, err := ioutil.ReadFile("fig/old_man_yells_at.png")
	require.NoError(t, err)
	ioutil.WriteFile("packed.go", []byte(fmt.Sprintf(
		`// Copyright 2021 oncilla
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yeller

var rawOldMan = []byte(%q)
`, base64.StdEncoding.EncodeToString(raw))), 0666)
}
