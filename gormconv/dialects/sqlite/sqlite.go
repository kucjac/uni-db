package sqlite

import (
	"github.com/kucjac/uni-db/gormconv"
	"github.com/kucjac/uni-db/sqliteconv"
	_ "github.com/neuronlabs/gorm/dialects/sqlite"
)

func init() {
	gormconv.Register("sqlite3", sqliteconv.New())
}
