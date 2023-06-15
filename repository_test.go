// Copyright 2023 oncilla
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
	"encoding/json"
	"flag"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	yeller "github.com/oncilla/old-man-yells-at"
	"github.com/stretchr/testify/require"
)

// update is a cmd line flag that enables golden file updates. To update the
// golden files simply run 'go test -update-repository ./...'.
var updateRepository = flag.Bool("update-repository", false, "set to true to regenerate repository")

func TestCreateRepository(t *testing.T) {
	if !*updateRepository {
		t.SkipNow()
	}

	// Clone the emojis repository
	cmd := exec.Command("git", "clone", "git@github.com:buildkite/emojis")
	cmd.Dir = t.TempDir()
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, string(out))

	type Description struct {
		Name  string `json:"name"`
		Image string `json:"image"`
	}

	raw, err := os.ReadFile(filepath.Join(cmd.Dir, "emojis", "img-buildkite-64.json"))
	require.NoError(t, err)

	var emojis []Description
	json.Unmarshal(raw, &emojis)

	for _, emoji := range emojis {
		emoji := emoji
		t.Run(emoji.Name, func(t *testing.T) {
			t.Parallel()

			raw, err := os.ReadFile(filepath.Join(cmd.Dir, "emojis", emoji.Image))
			require.NoError(t, err)
			target, err := png.Decode(bytes.NewReader(raw))
			require.NoError(t, err)
			yelledAt := yeller.YellAt(target)

			var buf bytes.Buffer
			require.NoError(t, png.Encode(&buf, yelledAt))

			out := filepath.Join("repository", "old-man-yells-at-"+emoji.Name+".png")
			require.NoError(t, os.WriteFile(out, buf.Bytes(), 0666))
		})
	}
}
