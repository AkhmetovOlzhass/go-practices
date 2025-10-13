package db

import (
    "fmt"
    "time"

    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

func Connect(dsn string) (*sqlx.DB, error) {
    db, err := sqlx.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)

    if err := db.Ping(); err != nil {
        return nil, err
    }

    fmt.Println("Успешно подключено к Postgre")
    return db, nil
}
