package migrations

import (
    "database/sql"
    "golearning/app/models"
    "golearning/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type Category struct {
        models.BaseModel

        Name        string `gorm:"type:varchar(255);not null;index"`
        Description string `gorm:"type:varchar(255);default:null"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&Category{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&Category{})
    }

    migrate.Add("2024_01_07_101315_add_categories_table", up, down)
}