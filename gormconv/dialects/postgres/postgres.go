package postgres

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/neuronlabs/uni-db/gormconv"
	"github.com/neuronlabs/uni-db/pgconv"
)

func init() {
	gormconv.Register("postgres", pgconv.New())
}
