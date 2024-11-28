package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
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
	DBClient *gorm.DB
}

func (m *MySQLHandler) InitMySQLConnection() {
	// MySQL connection details from environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		EnvVariables.DBUsername,
		EnvVariables.DBPassword,
		EnvVariables.DBHost,
		EnvVariables.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to MySQL database: %v", err))
	}

	m.DBClient = db
	ZapLogger.Info("MySQL Client Initiated")
}

func (m *MySQLHandler) InitTables() {
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
	if err := m.DBClient.Exec(string(sqlFile)).Error; err != nil {
		return fmt.Errorf("error executing SQL file: %v", err)
	}
	return nil
}

func MySQL() IMySQLHandler {
	if mysqlObj == nil {
		mysqlOnce.Do(func() {
			mysqlObj = &MySQLHandler{}
		})
	}
	return mysqlObj
}
