package postgres

import (
	"github.com/kucjac/uni-db/gormconv"
	"github.com/kucjac/uni-db/pgconv"
	_ "github.com/neuronlabs/gorm/dialects/postgres"
)

func init() {
	gormconv.Register("postgres", pgconv.New())
}
