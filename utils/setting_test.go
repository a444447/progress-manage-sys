package utils

import (
	"testing"
)

func TestSetting(t *testing.T) {

	if AppMode == "debug" {
		t.Log("OK!")
		t.Logf("Server: appmode %s httpPort %s \n", AppMode, HttpPort)
		t.Logf("DataBase: Db %s DbName %s DbPort %s", Db, DbName, DbPort)
	}
}
