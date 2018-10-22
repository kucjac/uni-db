package mysql

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kucjac/uni-db/gormconv"
	"github.com/kucjac/uni-db/mysqlconv"
)

func init() {
	gormconv.Register("mysql", mysqlconv.New())
}
