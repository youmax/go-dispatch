package tests

import (
	"reflect"
	"testing"
	"upay/configs"
	"upay/envs"
)

func LoadDatabaseConfiguration(conf map[string]string, t *testing.T) {
	db := configs.LoadDatabaseConfiguration(conf)
	s := reflect.ValueOf(db).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if reflect.DeepEqual(f.Interface(), reflect.Zero(reflect.TypeOf(f.Interface())).Interface()) {
			t.Errorf("%s = %v\n", typeOfT.Field(i).Name, f.Interface())
		}
	}
	t.Log(db)
}

func TestConfig(t *testing.T) {
	LoadDatabaseConfiguration(envs.Development, t)
	LoadDatabaseConfiguration(envs.Production, t)
}
