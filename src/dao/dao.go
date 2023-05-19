package dao

import (
	"GoTemplate/src/utils"
	"encoding/json"
	"fmt"

	"github.com/syklinux/golib/log"
	"github.com/syklinux/golib/mysql"
	"github.com/syklinux/golib/redis"

	"golang.org/x/sync/errgroup"
)

// Dao Dao
type Dao struct {
	closeFc []func() error
}

// NewDao NewDao
func NewDao(isAutoMigrate bool) *Dao {
	d := new(Dao)
	d.InitMysql(isAutoMigrate)
	return d
}

// InitMysql InitMysql
func (dao *Dao) InitMysql(isAutoMigrate bool) {
	defer func() {
		dao.closeFc = append(dao.closeFc, mysql.Close)
	}()
	strJSON, _ := json.Marshal(utils.Conf.Mysql)
	var client *mysql.MySQLConf
	err := json.Unmarshal([]byte(strJSON), &client)
	if err != nil {
		log.Fatalf("InitMysql err", err)
	}
	mysql.InitDb(client, isAutoMigrate)
}

// Close Close1
func (dao *Dao) Close() error {
	fmt.Println("停止mysql、redis连接")
	var g errgroup.Group
	for i := 0; i < len(dao.closeFc); i++ {
		fn := dao.closeFc[i]
		g.Go(func() error {
			return fn()
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

// InitRedis InitRedis1
func InitRedis() {
	redis.InitRedisCon(utils.Conf.RedisConf.Addr, utils.Conf.RedisConf.Password, utils.Conf.RedisConf.Db)
}

// RedisClose RedisClose
func RedisClose() {
	fmt.Println("关闭redis连接")
	_ = redis.Close()
}
