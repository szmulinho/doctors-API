package config

import (
	"testing"
)

func TestStorageConfig_ConnectionString(t *testing.T) {

	a := StorageConfig{
		Host:     "postgres",
		User:     "postgres",
		Password: "L96a1prosniper",
		Dbname:   "postgres",
		Port:     "5432",
		Sslmode:  "disable",
		TimeZone: "Europe/Warsaw",
	}
	result := a.ConnectionString()
	if result != "host=postgres user=postgres password=L96a1prosniper dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Warsaw" {
		t.Fail()
	}

}

//host=postgres user=postgres password=L96a1prosniper dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Warsaw
