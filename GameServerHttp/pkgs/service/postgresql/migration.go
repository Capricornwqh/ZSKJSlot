package service_pgsql

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const minDBVersion = 0

// Migration describes on migration from lower version to high version
type Migration interface {
	Version() string
	Description() string
	Migrate(ctx *gin.Context, db *gorm.DB) error
	ShouldCleanCache() bool
}

type migration struct {
	version          string
	description      string
	migrate          func(ctx *gin.Context, db *gorm.DB) error
	shouldCleanCache bool
}

// Version returns the migration's version
func (m *migration) Version() string {
	return m.version
}

// Description returns the migration's description
func (m *migration) Description() string {
	return m.description
}

// Migrate executes the migration
func (m *migration) Migrate(ctx *gin.Context, db *gorm.DB) error {
	return m.migrate(ctx, db)
}

// ShouldCleanCache should clean the cache
func (m *migration) ShouldCleanCache() bool {
	return m.shouldCleanCache
}

// NewMigration creates a new migration
func NewMigration(version, desc string, fn func(ctx *gin.Context, db *gorm.DB) error, shouldCleanCache bool) Migration {
	return &migration{version: version, description: desc, migrate: fn, shouldCleanCache: shouldCleanCache}
}

// Use noopMigration when there is a migration that has been no-oped
var noopMigration = func(_ *gin.Context, _ *gorm.DB) error { return nil }

var migrations = []Migration{
	NewMigration("v0.0.1", "this is first version, no operation", noopMigration, false),
}

func GetMigrations() []Migration {
	return migrations
}

// GetCurrentDBVersion returns the current db version
// func GetCurrentDBVersion(db *gorm.DB) (int64, error) {
// 	// 确保版本表存在
// 	if err := db.AutoMigrate(&entity_pgsql.Version{}); err != nil {
// 		return -1, fmt.Errorf("sync version failed: %v", err)
// 	}

// 	var currentVersion entity_pgsql.Version
// 	result := db.First(&currentVersion, 1) // 查找 ID = 1 的记录
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			// 如果记录不存在，创建新记录
// 			newVersion := entity_pgsql.Version{Id: 1, VersionNumber: 0}
// 			if err := db.Create(&newVersion).Error; err != nil {
// 				return -1, fmt.Errorf("insert first version failed: %v", err)
// 			}
// 			return 0, nil
// 		}
// 		return -1, fmt.Errorf("get first version failed: %v", result.Error)
// 	}
// 	return currentVersion.VersionNumber, nil
// }

// ExpectedVersion returns the expected db version
func ExpectedVersion() int64 {
	return int64(minDBVersion + len(migrations))
}

// Migrate database to current version
// func Migrate(debug bool, dbConf *data.Database, cacheConf *data.CacheConf, upgradeToSpecificVersion string) error {
// 	cache, cacheCleanup, err := data.NewCache(cacheConf)
// 	if err != nil {
// 		fmt.Println("new cache failed:", err)
// 	}
// 	engine, err := data.NewDB(debug, dbConf)
// 	if err != nil {
// 		fmt.Println("new database failed: ", err)
// 		return err
// 	}
// 	defer engine.Close()

// 	currentDBVersion, err := GetCurrentDBVersion(engine)
// 	if err != nil {
// 		return err
// 	}
// 	expectedVersion := ExpectedVersion()
// 	if len(upgradeToSpecificVersion) > 0 {
// 		fmt.Printf("[migrate] user set upgrade to version: %s\n", upgradeToSpecificVersion)
// 		for i, m := range migrations {
// 			if m.Version() == upgradeToSpecificVersion {
// 				currentDBVersion = int64(i)
// 				break
// 			}
// 		}
// 	}

// 	for currentDBVersion < expectedVersion {
// 		fmt.Printf("[migrate] current db version is %d, try to migrate version %d, latest version is %d\n",
// 			currentDBVersion, currentDBVersion+1, expectedVersion)
// 		migrationFunc := migrations[currentDBVersion]
// 		fmt.Printf("[migrate] try to migrate Answer version %s, description: %s\n", migrationFunc.Version(), migrationFunc.Description())
// 		if err := migrationFunc.Migrate(context.Background(), engine); err != nil {
// 			fmt.Printf("[migrate] migrate to db version %d failed: %s\n", currentDBVersion+1, err)
// 			return err
// 		}
// 		if migrationFunc.ShouldCleanCache() {
// 			if err := cache.Flush(context.Background()); err != nil {
// 				fmt.Printf("[migrate] flush cache failed: %s\n", err)
// 			}
// 		}
// 		fmt.Printf("[migrate] migrate to db version %d success\n", currentDBVersion+1)
// 		if _, err := engine.Update(&entity_pgsql.Version{Id: 1, VersionNumber: currentDBVersion + 1}); err != nil {
// 			fmt.Printf("[migrate] migrate to db version %d, update failed: %s", currentDBVersion+1, err)
// 			return err
// 		}
// 		currentDBVersion++
// 	}
// 	if cache != nil {
// 		cacheCleanup()
// 	}
// 	return nil
// }
