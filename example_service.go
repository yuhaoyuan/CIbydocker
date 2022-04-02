package utbydocker

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	redisCli *redis.Client
	mysqlCli *gorm.DB
)

func NewRedis() error {
	if redisCli == nil {
		redisCli = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // use default Addr
			Password: "",               // no password set
			DB:       0,                // use default DB
		})
	}

	_, err := redisCli.Ping(context.Background()).Result()
	return err
}

func NewGORM() error {
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	user := "root"
	pwd := "123456"
	dbname := ""

	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, dbname)
	mysqlCli, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	//mysqlCli.Exec("source init.sql")
	return nil
}

// Crop 根据种子创建的作物
type User struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"userId"`
	UserName string `json:"userName"`
}

// TableName 获取分表名
func (u *User) TableName() string {
	return "user_tab"
}

func SetRedisAndMysql(ctx context.Context) error {
	if redisCli == nil {
		_ = NewRedis()
	}
	if mysqlCli == nil {
		_ = NewGORM()
	}

	user := &User{}

	sqlStr := fmt.Sprintf("SELECT * FROM %v WHERE user_id = %v AND meta_id != 0 ORDER BY ID DESC LIMIT 1", "user_tab", "1001")
	err := mysqlCli.Raw(sqlStr).Scan(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}
