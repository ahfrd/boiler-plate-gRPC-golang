package database

import (
	"database/sql"
	"errors"
	"log"
	"grpc-boiler-plate-go/env"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocql/gocql"
)

type Env struct {
	Params *env.EnvironmentParameters
	Error  error
}

func NewMySQLDB(config env.EnvironmentParameters) (*sql.DB, error) {
	conn, err := sql.Open(config.Database.MySQLDB.DBType, config.SetupMySQLDBConnection().Database.MySQLDB.DBConfig)
	if err != nil {
		return conn, errors.New("failed connecting to database : " + err.Error())
	}

	if err = conn.Ping(); err != nil {
		return conn, errors.New("failed pinging to database : " + err.Error())
	}

	return conn, err
}

func NewScyllaDB(config env.EnvironmentParameters) (*gocql.Session, error) {
	cluster := gocql.NewCluster(config.Database.ScyllaDB.DBCluster...)

	// set scylla cluster config
	cluster.Keyspace = config.Database.ScyllaDB.DBKeyspace
	cluster.Timeout = 1 * time.Second
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: config.Database.ScyllaDB.DBUser, Password: config.Database.ScyllaDB.DBPass}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(config.Database.ScyllaDB.DBPolicy)

	// create scylla session
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		log.Printf("Failed connecting to scylla db on keyspace : %s ", config.Database.ScyllaDB.DBKeyspace)
		return session, err
	}

	log.Printf("Connected to scylla db on keyspace : %s ", config.Database.ScyllaDB.DBKeyspace)
	return session, nil
}
