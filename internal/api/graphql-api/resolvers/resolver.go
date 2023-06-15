package resolvers

import "github.com/jinzhu/gorm"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your cmd, add any dependencies you require here.

type Resolver struct {
	Database *gorm.DB
}
