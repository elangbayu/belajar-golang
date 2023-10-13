package db

var connection string

func init() {
  connection = "SQL"
}

func GetDatabase() string {
  return connection
}

