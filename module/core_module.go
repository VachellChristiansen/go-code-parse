package module

import (
	"go-code-parse/database"
	"go-code-parse/helper"
)

type CoreModule struct {
	Database database.Database
	Helper   helper.Helper
}

func NewCoreModule() CoreModule {
	return CoreModule{
		Database: database.NewDatabase(),
		Helper:   helper.NewHelper(),
	}
}
