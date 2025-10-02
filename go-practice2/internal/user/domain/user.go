package domain

type User struct {
    ID   int    `json:"user_id" gorm:"primaryKey"`
    Name string `json:"created"`
}
