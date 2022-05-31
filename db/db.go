package db

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	Instance *sqlx.DB
)

func Init(ctx context.Context) error {
	err := connect()
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	go loop(ctx)
	return nil
}

func connect() error {
	var err error
	var db *sqlx.DB
	//db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "ro", "ro", "content-cdn.projecteq.net", "16033", "peq_content"))
	db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "eqemu", "eqemu", "127.0.0.1", "3306", "peq"))
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	Instance = db
	return nil
}

func loop(ctx context.Context) {
	var err error
	for {
		select {
		case <-time.After(60 * time.Second):
			if err = Instance.Ping(); err != nil {
				err = connect()
				if err != nil {
					fmt.Println("db.loop: connect:", err)
				}
			}
		case <-ctx.Done():
			if Instance != nil {
				Instance.Close()
			}
		}
	}
}
