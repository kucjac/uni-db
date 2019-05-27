package mysql

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/neuronlabs/uni-db/gormconv"
	"github.com/neuronlabs/uni-db/mysqlconv"
)

func init() {
	gormconv.Register("mysql", mysqlconv.New())
}
