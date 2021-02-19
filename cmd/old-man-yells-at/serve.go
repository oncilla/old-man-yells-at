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
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/oncilla/boa/pkg/boa/flag"
	"github.com/spf13/cobra"

	"github.com/oncilla/old-man-yells-at/server"
	"github.com/oncilla/old-man-yells-at/server/memory"
)

func newServe(pather CommandPather) *cobra.Command {
	flags := struct {
		addr flag.TCPAddr
	}{
		addr: flag.TCPAddr{
			IP:   net.ParseIP("::"),
			Port: 0,
		},
	}

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start an http server that makes Abe yell at stuff",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			store, err := memory.NewStore()
			if err != nil {
				return fmt.Errorf("opening store: %v", err)
			}

			s := &server.Server{
				Store: store,
			}

			mux := http.NewServeMux()
			mux.Handle("/", http.RedirectHandler("upload", http.StatusMovedPermanently))
			mux.HandleFunc("/upload", s.Upload)
			mux.HandleFunc("/image/", s.Image)

			if port := os.Getenv("PORT"); flags.addr.Port == 0 && port != "" {
				p, err := strconv.Atoi(port)
				if err != nil {
					return fmt.Errorf("parsing $PORT: %v", err)
				}
				flags.addr.Port = p
			}

			l, err := net.Listen("tcp", flags.addr.String())
			if err != nil {
				return fmt.Errorf("listening: %v", err)
			}
			fmt.Println("Listening on", l.Addr().String())
			return http.Serve(l, mux)
		},
	}

	cmd.Flags().Var(&flags.addr, "addr", "Address to listen on")

	return cmd
}
