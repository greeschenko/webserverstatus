package main

import (
	"testing"
)

func TestPropReadConfigFile(t *testing.T) {
	getConfig()
	if Conf.Server.Domen != "polonex.com.ua" {
		t.Fatal("not prop server domen")
	}
	return
}

func TestExecuteOne(t *testing.T) {
	r := executeOne("printf OK")
	if r != "OK" {
		t.Fatal("execute commands from strings dont work")
	}
	return
}

func TestToJson(t *testing.T) {
	data := Status{"OK"}
	r := toJson(data)
	if len(r) == 0 {
		t.Fatal("convert to json dont work")
	}
	return
}
