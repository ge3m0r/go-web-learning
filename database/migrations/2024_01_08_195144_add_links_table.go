package migrations

import (
    "database/sql"
    "golearning/app/models"
    "golearning/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type Link struct {
        models.BaseModel

        Name string `gorm:"type:varchar(255);not null"`
        URL  string `gorm:"type:varchar(255);default:null"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&Link{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&Link{})
    }

    migrate.Add("2024_01_08_195144_add_links_table", up, down)
}