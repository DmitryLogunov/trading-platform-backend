package graph

import (
	"github.com/jinzhu/gorm"
)

type Resolver struct {
	Database *gorm.DB
}
