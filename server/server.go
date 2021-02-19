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

package server

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"image"
	"image/png"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	yeller "github.com/oncilla/old-man-yells-at"
	"go.uber.org/zap"
)

// Server implements the HTTP server.
type Server struct {
	Store  Store
	Logger zap.Logger
}

func (s *Server) Upload(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("server/upload.gtpl")
		t.Execute(w, token)
	case r.Method == "POST":
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		mm, _, err := image.Decode(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		var b bytes.Buffer
		if err := png.Encode(&b, yeller.YellAt(mm)); err != nil {
			fmt.Println(err)
			return
		}

		m := Image{
			Name: handler.Filename[:len(handler.Filename)-len(filepath.Ext(handler.Filename))],
			Raw:  b.Bytes(),
			UUID: uuid.New(),
		}
		fmt.Println(m.UUID.String())

		if err := s.Store.Add(r.Context(), m); err != nil {
			fmt.Println(err)
			return
		}

		// TODO(): horrible
		w.Write([]byte(
			fmt.Sprintf(`
			<html>
			<head>
			</head>
			<body>
			<img src="image/%s/old-man-yells-at-%s.png">
			</body>
			</html>
			`, m.UUID, m.Name),
		))
	default:
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
	}
}

func (s *Server) Image(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	list := strings.Split(r.URL.Path, "/")
	if len(list) == 0 {
		http.Error(w, "UUID not provided", http.StatusBadRequest)
		return
	}
	// TODO(): Horrible
	id, err := uuid.Parse(list[len(list)-2])
	if err != nil {
		http.Error(w, "UUID malformed", http.StatusBadRequest)
		return
	}
	m, err := s.Store.Get(r.Context(), id)
	if err != nil {
		http.Error(w, "UUID not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(m.Raw)
}
