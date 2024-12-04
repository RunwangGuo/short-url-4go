package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"short-url-4go/src/infrastrctures"
	"short-url-4go/src/models"
	"sync"
)

var (
	mysqlObj  *MySQLHandler
	mysqlOnce sync.Once
)

type IMySQLHandler interface {
	InitMySQLConnection()
	InitTables()
}

type MySQLHandler struct {
	DBClient *infrastrctures.MySQLClient
}

func (m *MySQLHandler) InitMySQLConnection() {
	// MySQL connection details from environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		EnvVariables.DBUsername,
		EnvVariables.DBPassword,
		EnvVariables.DBHost,
		EnvVariables.DBPort,
		EnvVariables.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info), // 打开sql日志
	})
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to MySQL database: %v", err))
	}

	// 将 gorm.DB 包装为 infrastructures.MySQLClient
	m.DBClient = &infrastrctures.MySQLClient{
		DB: db,
	}

	ZapLogger.Info("MySQL Client Initiated")
}

func (m *MySQLHandler) InitTables() {
	// 使用 AutoMigrate 自动迁移表
	if err := m.DBClient.DB.AutoMigrate(&models.Link{}, &models.AccessLog{}); err != nil {
		panic(fmt.Sprintf("Failed to auto-migrate tables: %v", err))
	}
	ZapLogger.Info("MySQL Tables AutoMigrated")
}

/*func (m *MySQLHandler) InitTables() {
	// Run the SQL scripts from db.sql for table creation
	if err := m.runSQLFile("src/db.sql"); err != nil {
		panic(fmt.Sprintf("Failed to initialize tables: %v", err))
	} else {
		ZapLogger.Info("Tables initialized successfully")
	}
}

func (m *MySQLHandler) runSQLFile(filePath string) error {
	// Read the SQL file
	sqlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading SQL file: %v", err)
	}

	// Run the SQL file
	if err := m.DBClient.DB.Exec(string(sqlFile)).Error; err != nil {
		return fmt.Errorf("error executing SQL file: %v", err)
	}
	return nil
}*/

func MySQL() IMySQLHandler {
	if mysqlObj == nil {
		mysqlOnce.Do(func() {
			mysqlObj = &MySQLHandler{}
		})
	}
	return mysqlObj
}
