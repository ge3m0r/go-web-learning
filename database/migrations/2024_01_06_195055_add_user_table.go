package migrations

import (
    "database/sql"
    "golearning/app/models"
    "golearning/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type AddUserTable struct {
        models.BaseModel

        Name     string `gorm:"type:varchar(255);not null;index"`
        Email    string `gorm:"type:varchar(255);index;default:null"`
        Phone    string `gorm:"type:varchar(20);index;default:null"`
        Password string `gorm:"type:varchar(255)"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&AddUserTable{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&AddUserTable{})
    }

    migrate.Add("2024_01_06_195055_add_user_table", up, down)
}