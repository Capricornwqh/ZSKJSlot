package service_pgsql

import (
	entity_pgsql "SlotGameServer/pkgs/dao/postgresql/entity"
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PGSQLService struct {
	ctx  context.Context
	DB   *gorm.DB
	err  error
	Done bool
}

// 初始化数据表
func (s *PGSQLService) ititDB(ctx context.Context) error {
	s.ctx = ctx
	s.do("check table exist", s.checkTableExist)
	s.do("sync table", s.syncTable)
	s.do("init version", s.initVersionTable)
	return s.err
}

func (s *PGSQLService) do(taskName string, fn func()) {
	if s.err != nil || s.Done {
		return
	}
	fn()
	if s.err != nil {
		s.err = fmt.Errorf("%s failed: %s", taskName, s.err)
	}
}

func (s *PGSQLService) checkTableExist() {
	// gorm检查表是否存在的方法
	s.Done = s.DB.WithContext(s.ctx).Migrator().HasTable(&entity_pgsql.Version{})
	if s.Done {
		logrus.WithContext(s.ctx).Info("[database] already exists")
	}
}

func (s *PGSQLService) syncTable() {
	// gorm的自动迁移方法
	sliceTables := []any{
		entity_pgsql.Account{},
		entity_pgsql.Version{},
	}
	for _, table := range sliceTables {
		var tableOptions string
		var tableName string
		var sqlString string
		// 根据不同表类型设置对应的comment
		switch v := table.(type) {
		case entity_pgsql.Account:
			tableOptions = v.Comment()
			tableName = v.TableName()
			sqlString = v.SetAccountIdStartValue()
		case entity_pgsql.Version:
			tableOptions = v.Comment()
			tableName = v.TableName()
		default:
		}

		s.err = s.DB.WithContext(s.ctx).AutoMigrate(table)

		if tableOptions != "" {
			s.DB.Exec(fmt.Sprintf("COMMENT ON TABLE %s IS '%s';", tableName, tableOptions))
		}
		if sqlString != "" {
			s.DB.Exec(sqlString)
		}

		if s.err != nil {
			return
		}
	}
}

// 定时器
func (s *PGSQLService) StartScheduler() *gocron.Scheduler {
	ctx := context.Background()
	err := s.ititDB(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	scheduler := gocron.NewScheduler(time.UTC)
	// // 添加Revision表的定时任务
	// tmpSchemaService.MigrateRevision(ctx)
	// nextYear := time.Now().AddDate(1, 0, 0) // 下一年的同一天
	// startTime := time.Date(nextYear.Year(), nextYear.Month(), nextYear.Day(), 0, 0, 0, 0, time.Local)
	// scheduler.Every(1).StartAt(startTime).SingletonMode().Do(func() {
	// 	tmpSchemaService.MigrateRevision(ctx)
	// })

	// 启动调度器
	scheduler.StartAsync()

	return scheduler
}

func (s *PGSQLService) expectedVersion() int64 {
	return int64(minDBVersion + len(migrations))
}

func (s *PGSQLService) initVersionTable() {
	s.err = s.DB.WithContext(s.ctx).Create(&entity_pgsql.Version{Id: 1, VersionNumber: s.expectedVersion()}).Error
}

// 创建Revision 每年生成下一年的表
func (s *PGSQLService) MigrateRevision(ctx context.Context) error {
	// nowTime := time.Now()
	// tmpRecord := entity_pgsql.Record{}

	// for i := range 2 {
	// 	tmpTableName := fmt.Sprintf("record_%d", nowTime.Year()+i)
	// 	if !s.DB.Migrator().HasTable(tmpTableName) {
	// 		err := s.DB.Table(tmpTableName).Migrator().CreateTable(tmpRecord)
	// 		if err != nil {
	// 			logrus.WithContext(ctx).Error(err)
	// 			return err
	// 		}
	// 		if tmpRecord.Comment() != "" {
	// 			s.DB.Exec(fmt.Sprintf("COMMENT ON TABLE %s IS '%s';", tmpTableName, tmpRecord.Comment()))
	// 		}
	// 	}
	// }

	return nil
}
