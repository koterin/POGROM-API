package db

import (
	"database/sql"
	"fmt"
	"pogrom/internal/entity"
)

var DB *sql.DB

func FindUser(username string) (entity.UserInfo, error) {
	var info entity.UserInfo

	err := DB.QueryRow(`SELECT * FROM users WHERE username=$1;`,
		username).Scan(&info.Id, &info.Username, &info.Role,
		&info.Status, &info.ChatId, &info.CreatedAt)
	if err != nil {
		err = fmt.Errorf("\nError in FindUser: SELECT * FROM users WHERE username=%s: %w", username, err)
	}

	return info, err
}

func GetItemInfo(id string) (entity.Item, error) {
	var info entity.Item

	err := DB.QueryRow(`SELECT * FROM items WHERE id=$1;`,
		id).Scan(&info.Title, &info.Desc, &info.Status, &info.Category,
		&info.ImageLink, &info.RawPrice, &info.StartPrice, &info.OwnerId,
		&info.CurrentPrice, &info.CreatedAt, &info.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("\nError in GetItem: SELECT * FROM items WHERE id=%s: %w", id, err)
	}

	return info, err
}
