package database

import "github.com/gocql/gocql"

type DBConnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}


var connection DBConnection

func SetupDBConnection(){
	connection.cluster = gocql.NewCluster("127.0.0.1")
	connection.cluster.Consistency = gocql.Quorum
	connection.cluster.Keyspace = "first_app"
	connection.session, _ = connection.cluster.CreateSession()
}
