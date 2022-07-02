package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
)

type ToDo struct {
	Id   int
	Name string
}

func Create(db *sql.DB, t *ToDo) error {
	_, err := db.Exec("INSERT INTO todo(name) VALUES ( ? )", t.Name)
	return err
}

func createContainer() (*dockertest.Resource, *dockertest.Pool) {
	// Dockerコンテナへのファイルマウント時に絶対パスが必要
	pwd, _ := os.Getwd()

	// Dockerとの接続
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Dockerコンテナ起動時の細かいオプションを指定する
	// テーブル定義などはここで流し込むのが良さそう
	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "5.7.33",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=secret",
		},
		Mounts: []string{
			pwd + "/my.cnf:/etc/mysql/my.cnf",                      // MySQLの設定ファイル
			pwd + "/todo.sql:/docker-entrypoint-initdb.d/todo.sql", // コンテナ起動時に実行したいSQL
		},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
		},
	}

	// コンテナを起動
	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource, pool
}

func closeContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	// コンテナの終了
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func connectDB(resource *dockertest.Resource, pool *dockertest.Pool) *sql.DB {
	// DB(コンテナ)との接続
	var db *sql.DB
	if err := pool.Retry(func() error {
		// DBコンテナが立ち上がってから疎通可能になるまで少しかかるのでちょっと待ったほうが良さそう
		time.Sleep(time.Second * 10)

		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mydb?parseTime=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return db
}

func TestCreateWithDB(t *testing.T) {
	t.Run(
		"レコードが正しく追加されることのテスト",
		func(t *testing.T) {
			// Arrange
			// コンテナ(DB)の立ち上げ, 接続
			resource, pool := createContainer()
			defer closeContainer(resource, pool)
			db := connectDB(resource, pool)
			expected := &ToDo{
				Name: "testToDo",
			}

			// Act
			// テストDBに対して通常通りCreateを実行
			err := Create(db, expected)
			if err != nil {
				t.Error(err.Error())
			}

			// Assert
			// 検証用にINSERTしたもののIDを取得する
			// かっこ悪いのでCreate()がidを返すように書き換えたほうが良さそう
			var id *int
			err = db.QueryRow("SELECT LAST_INSERT_ID() FROM todo").Scan(&id)
			if err != nil {
				t.Error(err.Error())
			}
			// 挿入されたレコードを取得して値を検証する
			actual := &ToDo{}
			err = db.QueryRow("SELECT id, name FROM todo WHERE id = ?", id).Scan(&actual.Id, &actual.Name)
			if err != nil {
				t.Error(err.Error())
			}
			if actual.Name != expected.Name {
				t.Errorf("expected: %s, actual: %s", expected.Name, actual.Name)
			}
		},
	)
}
