package postgres

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kucjac/uni-db/gormconv"
	"github.com/kucjac/uni-db/pgconv"
)

func init() {
	gormconv.Register("postgres", pgconv.New())
}
