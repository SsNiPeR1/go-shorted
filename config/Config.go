package config

const (
	Host     = ""
	Port     = 5432
	User     = "postgres"
	Password = ""
	Dbname   = "shorted"
	AdminKey = "" // this is the admin key, you can change it however you want
	// this key is used to access the /db route, also used to create vanity URLs or deleting shorted URLs from the database

	Website = ""
	Charset = "qwertyuiopasdfghjklzxcvbnm1234567890" // it is possible to change this however you want
)
