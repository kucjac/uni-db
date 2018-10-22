package sqlite

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kucjac/uni-db/gormconv"
	"github.com/kucjac/uni-db/sqliteconv"
)

func init() {
	gormconv.Register("sqlite3", sqliteconv.New())
}
