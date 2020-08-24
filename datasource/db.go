package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
	"wumiao/config"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

func InstanceMaster() *xorm.Engine {

	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}
	c := config.MasterDbConfig

	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(config.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster,", err)
		return nil
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(false)
	//engine.SetTZLocation(conf.SysTimeLocation)

	// 性能优化的时候才考虑，加上本机的SQL缓存
	cache := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cache)
	masterEngine = engine
	return engine
}

func InstanceSlave() *xorm.Engine {

	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}
	c := config.SlaveDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(config.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbHelper.DbInstanceMaster error=", err)
		return nil
	} else {
		slaveEngine = engine
		return masterEngine
	}
}
