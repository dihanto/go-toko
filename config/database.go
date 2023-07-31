package config

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitDatabaseConnection() *sql.DB {
	InitLoadConfiguration()
	databaseName := viper.GetString("database.name")
	databaseHost := viper.GetString("database.host")
	databasePortInt := viper.GetInt("database.port")
	databasePort := strconv.Itoa(databasePortInt)
	databaseUser := viper.GetString("database.user")
	databasePassword := viper.GetString("database.password")
	databaseDBName := viper.GetString("database.database_name")
	connMaxIdleTime := viper.GetDuration("database_connection.conn_max_idle_time")
	connMaxLifeTime := viper.GetDuration("database_connection.conn_max_life_time")
	maxIdleConn := viper.GetInt("database_connection.max_idle_conn")
	maxOpenConn := viper.GetInt("database_connection.max_open_conn")

	db, err := sql.Open(databaseName, "host="+databaseHost+" port="+databasePort+" user="+databaseUser+" password="+databasePassword+" dbname="+databaseDBName+" sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	db.SetConnMaxLifetime(connMaxLifeTime * time.Second)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	return db

}
