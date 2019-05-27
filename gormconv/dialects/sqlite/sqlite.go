package sqlite

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/neuronlabs/uni-db/gormconv"
	"github.com/neuronlabs/uni-db/sqliteconv"
)

func init() {
	gormconv.Register("sqlite3", sqliteconv.New())
}
