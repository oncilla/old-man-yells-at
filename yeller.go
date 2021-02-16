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

package yeller

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/draw"
	"image/png"

	"github.com/nfnt/resize"
)

var oldman image.Image = func() image.Image {
	reader := base64.NewDecoder(base64.StdEncoding, bytes.NewReader(rawOldMan))
	m, err := png.Decode(reader)
	if err != nil {
		panic(err)
	}
	return m
}()

// YellAt creates an image with Abe Simpson yelling the target.
func YellAt(target image.Image) image.Image {
	bounds := oldman.Bounds()

	yelled := image.NewRGBA(bounds)
	draw.Draw(yelled, bounds, oldman, image.Point{}, draw.Src)

	at := scaleDown(target)
	draw.Draw(yelled, at.Bounds(), at, image.Point{}, draw.Over)

	return yelled
}

func scaleDown(target image.Image) image.Image {
	s := target.Bounds().Size()
	height := float64(s.Y) * (50 / float64(s.X))
	return resize.Resize(50, uint(height), target, resize.Lanczos3)
}
