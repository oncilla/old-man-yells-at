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

package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	yeller "github.com/oncilla/old-man-yells-at"
)

// CommandPather returns the path to a command.
type CommandPather interface {
	CommandPath() string
}

func main() {
	var flags struct {
		out string
	}

	executable := filepath.Base(os.Args[0])
	cmd := &cobra.Command{
		Use:   executable + " <target-file>",
		Short: executable + " makes Abe yell at stuff!",
		Long: `Enjoy Abe yelling at stuff!

Provide an target image and Abe Simpson will yell at.

By default, the resulting image is created in the current working directory
as 'old-man-yells-at-<target-basename>.png'. If Abe should redirect his yelling,
you have the following options:

  - <filename>.png: Create image at the specified filename.
  - png: Create image at 'old-man-yells-at-<target-basename>.png'.
  - hex: Write image hex-encoded to stdout.
  - b64: Write image b64-encoded to stdout.
`,
		SilenceErrors: true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateOutput(flags.out); err != nil {
				return err
			}
			cmd.SilenceUsage = true
			return yell(args[0], flags.out)
		},
	}
	cmd.AddCommand(
		newCompletion(cmd),
		newVersion(cmd),
	)
	cmd.Flags().StringVarP(&flags.out, "output", "o", "png", `[png, b64, hex, <filename>.png]`)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func yell(filename, output string) error {
	input, err := loadImage(filename)
	if err != nil {
		return err
	}
	m := yeller.YellAt(input)

	switch {
	case output == "png":
		basename := filepath.Base(filename)
		name := strings.TrimSuffix(basename, filepath.Ext(basename))
		name = "old-man-yells-at-" + name + ".png"

		if err := writeImageFile(m, name); err != nil {
			return fmt.Errorf("writing image: %v", err)
		}
		return nil
	case output == "b64":
		enc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
		return png.Encode(enc, m)
	case output == "hex":
		enc := hex.NewEncoder(os.Stdout)
		return png.Encode(enc, m)
	case filepath.Ext(output) == ".png":
		if err := writeImageFile(m, output); err != nil {
			return fmt.Errorf("writing image: %v", err)
		}
		return nil
	default:
		return fmt.Errorf("unsupported output: %s", output)
	}
}

func loadImage(filename string) (image.Image, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("reading input file: %v", err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("decoding input image: %v", err)
	}
	return m, err
}

func writeImageFile(m image.Image, filename string) error {
	if err := os.MkdirAll(filepath.Dir(filename), 0644); err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(f, m); err != nil {
		return err
	}
	fmt.Printf("Abe is yelling: %s\n", filename)
	return nil
}

func validateOutput(output string) error {
	switch {
	case output == "png", output == "b64", output == "hex", filepath.Ext(output) == ".png":
		return nil
	default:
		return fmt.Errorf("unsupported output: %s", output)
	}
}
