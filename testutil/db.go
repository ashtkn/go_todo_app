package testutil

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	port := 33306
	if _, defined := os.LookupEnv("CI"); defined {
		port = 3306
	}

	db, err := sql.Open("mysql", fmt.Sprintf("todo:todo@tcp(127.0.0.1:%d)/todo?parseTime=true", port))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { db.Close() })

	return sqlx.NewDb(db, "mysql")
}

func OpenRedisForTest(t *testing.T) *redis.Client {
	t.Helper()

	host := "127.0.0.1"
	port := 36379
	if _, defined := os.LookupEnv("CI"); defined {
		port = 6379
	}

	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       0,
	})

	if err := cli.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("failed to connect redis: %s", err)
	}

	return cli
}
