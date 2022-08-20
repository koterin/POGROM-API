package db

import (
	"database/sql"
	"fmt"
	"pogrom/internal/entity"
)

var DB *sql.DB

func FindUser(username string) (entity.UserInfo, error) {
	var info entity.UserInfo

	info.Username = username

	err := DB.QueryRow(`SELECT id, role, status, chat_id, created_at FROM users
		WHERE username=$1;`, username).Scan(&info.Id, &info.Role,
		&info.Status, &info.ChatId, &info.CreatedAt)
	if err != nil {
		err = fmt.Errorf("\n35: SELECT id FROM users WHERE username=%s: %w", username, err)
	}

	return info, err
}
