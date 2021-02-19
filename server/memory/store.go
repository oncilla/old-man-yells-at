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

package memory

import (
	"context"
	"fmt"

	"github.com/dgraph-io/ristretto"
	"github.com/google/uuid"
	"github.com/oncilla/old-man-yells-at/server"
)

// Store is a sqlite backed implementation of the store interface.
type Store struct {
	cache *ristretto.Cache
}

func NewStore() (*Store, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 28,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	return &Store{cache: cache}, nil

}

func (s *Store) Add(ctx context.Context, m server.Image) error {
	if !s.cache.Set(m.UUID[:], m, int64(len(m.Name)+len(m.Raw)+len(m.UUID)+50)) {
		return fmt.Errorf("image not added to the cache: %s", m.UUID)
	}
	return nil
}

func (s *Store) Get(ctx context.Context, id uuid.UUID) (server.Image, error) {
	mm, ok := s.cache.Get(id[:])
	if !ok {
		return server.Image{}, fmt.Errorf("image not found: %s", id)
	}
	return mm.(server.Image), nil
}

func (s *Store) Search(ctx context.Context, name string) ([]server.Image, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *Store) List(ctx context.Context, pageSize, pageNumber int) ([]server.Image, error) {
	return nil, fmt.Errorf("not implemented")
}
