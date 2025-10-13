package user

import (
    "fmt"
    "github.com/jmoiron/sqlx"
)

type Repository struct {
    db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) InsertUser(u User) error {
    query := `
        INSERT INTO users (name, email, balance)
        VALUES (:name, :email, :balance)
    `
    _, err := r.db.NamedExec(query, u)
    return err
}

func (r *Repository) GetAllUsers() ([]User, error) {
    var users []User
    err := r.db.Select(&users, "SELECT * FROM users")
    return users, err
}

func (r *Repository) GetUserByID(id int) (User, error) {
    var u User
    err := r.db.Get(&u, "SELECT * FROM users WHERE id=$1", id)
    return u, err
}

func (r *Repository) TransferBalance(fromID, toID int, amount float64) error {
    tx, err := r.db.Beginx()
    if err != nil {
        return err
    }

    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        }
    }()

    var sender User
    err = tx.Get(&sender, "SELECT * FROM users WHERE id=$1", fromID)
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("sender not found: %w", err)
    }

    if sender.Balance < amount {
        tx.Rollback()
        return fmt.Errorf("insufficient funds")
    }

    _, err = tx.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", amount, fromID)
    if err != nil {
        tx.Rollback()
        return err
    }

    _, err = tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, toID)
    if err != nil {
        tx.Rollback()
        return err
    }
    return tx.Commit()
}
