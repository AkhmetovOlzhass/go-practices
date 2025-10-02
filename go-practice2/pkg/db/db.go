package db

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "go-prcatice2/internal/user/domain"
)

func InitDB() (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "host=localhost user=postgres password=postgres dbname=practice2 port=5432 sslmode=disable TimeZone=UTC",
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    if err := db.AutoMigrate(&domain.User{}); err != nil {
        return nil, err
    }

    return db, nil
}
