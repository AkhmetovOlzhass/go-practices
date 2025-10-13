package internal

import (
    "fmt"
    "go-practice4/internal/config"
    "go-practice4/internal/db"
    "go-practice4/internal/user"
)

type App struct{}

func NewApp() *App {
    return &App{}
}

func (a *App) Run() error {
    cfg := config.New()

    database, err := db.Connect(cfg.DSN)
    if err != nil {
        return fmt.Errorf("ошибка подключения к БД: %w", err)
    }
    defer database.Close()

    repo := user.NewRepository(database)
    fmt.Println("Подключено к Postgre")

    usersToInsert := []user.User{
        {Name: "Olzhas", Email: "mrbll@gmail.com", Balance: 500.0},
        {Name: "Anelya", Email: "anelya@list.ru", Balance: 300.0},
        {Name: "Asan", Email: "asan@mail.ru", Balance: 150.0},
    }

    for _, u := range usersToInsert {
        if err := repo.InsertUser(u); err != nil {
            return fmt.Errorf("insert error for %s: %w", u.Name, err)
        }
    }
    fmt.Println("Пользователи добавлены")

    users, err := repo.GetAllUsers()
    if err != nil {
        return fmt.Errorf("select error: %w", err)
    }

    fmt.Println("Текущие пользователи:")
    for _, u := range users {
        fmt.Printf("%d | %s | %s | %.2f\n", u.ID, u.Name, u.Email, u.Balance)
    }

    if len(users) > 0 {
        firstID := users[0].ID
        oneUser, err := repo.GetUserByID(firstID)
        if err != nil {
            return fmt.Errorf("get by id error: %w", err)
        }
        fmt.Printf("🔍 Пользователь по ID %d: %s | %.2f\n", oneUser.ID, oneUser.Name, oneUser.Balance)
    }

    fmt.Println("\nТестируем перевод 100.00 от Alice к Olzhas")
    err = repo.TransferBalance(users[0].ID, users[1].ID, 100.00)
    if err != nil {
        return fmt.Errorf("transfer error: %w", err)
    }

    updatedUsers, err := repo.GetAllUsers()
    if err != nil {
        return fmt.Errorf("select after transfer error: %w", err)
    }

    fmt.Println("Балансы после перевода:")
    for _, u := range updatedUsers {
        fmt.Printf("%d | %s | %.2f\n", u.ID, u.Name, u.Balance)
    }

    return nil
}
