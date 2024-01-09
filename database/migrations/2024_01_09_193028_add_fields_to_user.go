package migrations

import (
    "database/sql"
    "golearning/app/models"
    "golearning/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type AddFieldsToUser struct {
        models.BaseModel

        Name     string `gorm:"type:varchar(255);not null;index"`
        Email    string `gorm:"type:varchar(255);index;default:null"`
        Phone    string `gorm:"type:varchar(20);index;default:null"`
        Password string `gorm:"type:varchar(255)"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&AddFieldsToUser{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&AddFieldsToUser{})
    }

    migrate.Add("2024_01_09_193028_add_fields_to_user", up, down)
}