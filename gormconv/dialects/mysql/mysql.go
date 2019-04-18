package mysql

import (
	"github.com/kucjac/uni-db/gormconv"
	"github.com/kucjac/uni-db/mysqlconv"
	_ "github.com/neuronlabs/gorm/dialects/mysql"
)

func init() {
	gormconv.Register("mysql", mysqlconv.New())
}
