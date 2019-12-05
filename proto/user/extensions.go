package user

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// BeforeCreate generates and saves a uuid
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return errors.Wrapf(err, "BeforeCreate")
	}

	return scope.SetColumn("Id", uuid.String())
}
