package version

import (
	"runtime"

	"gorm.io/gorm"

	"odmp-admin/cmd/migrate/migration"
	"odmp-admin/cmd/migrate/migration/models"
	common "odmp-admin/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1614073723691Test)
}

func _1614073723691Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		// TODO: 这里开始写入要变更的内容

		err := tx.Model(&models.SysMenu{}).Where("path = ?", "/api/v1/syscontentList").Update("path", "/api/v1/syscontent").Error
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
