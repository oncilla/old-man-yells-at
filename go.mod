module github.com/oncilla/old-man-yells-at

go 1.15

require (
	github.com/dgraph-io/ristretto v0.0.3
	github.com/google/uuid v1.3.0
	github.com/lib/pq v1.9.0
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/oncilla/boa v0.1.3
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.4.0
	go.uber.org/zap v1.10.0
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	lukechampine.com/uint128 v1.3.0 // indirect
	modernc.org/ccgo/v3 v3.16.14 // indirect
	modernc.org/sqlite v1.23.1
)

// +heroku goVersion 1.15
// +heroku install ./cmd/...
