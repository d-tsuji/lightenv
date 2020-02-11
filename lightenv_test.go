package lightenv

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProcess_Normal_1(t *testing.T) {
	os.Clearenv()
	os.Setenv("APP_NAME", "TEST")
	os.Setenv("IP", "192.168.1.1")
	os.Setenv("APP_NUMBER", "-1")
	os.Setenv("APP_DISABLE_DEFAULT", "dummy")

	type Specification struct {
		AppName           string `name:"APP_NAME"`
		IP                string
		AppNumber         int    `name:"APP_NUMBER" required:"true"`
		AppEnableDefault  string `name:"APP_ENABLE_DEFAULT" default:"myDefault"`
		AppDisableDefault string `name:"APP_DISABLE_DEFAULT"`
		AppNoParam        string `name:"APP_NO_PARAM"`
	}

	var res Specification
	if err := Process(&res); err != nil {
		t.Error(err)
	}
	expected := Specification{
		AppName:           "TEST",
		IP:                "192.168.1.1",
		AppNumber:         -1,
		AppEnableDefault:  "myDefault",
		AppDisableDefault: "dummy",
		AppNoParam:        "",
	}
	if diff := cmp.Diff(res, expected); diff != "" {
		t.Errorf("TestProcess differs: (-got +want)\n%s", diff)
	}
}

func TestProcess_AbNormal_1(t *testing.T) {
	os.Clearenv()
	os.Setenv("APP_NAME", "TEST")

	type Specification struct {
		AppName string `name:"APP_NAME"`
		IP      string `required:"true"`
	}

	var res Specification
	err := Process(&res)
	if err == nil {
		t.Error("Process expect to occur error. Because of no setting param(IP)")
	}
}

func TestProcess_AbNormal_2(t *testing.T) {
	os.Clearenv()
	os.Setenv("APP_NUMBER", "TEST")

	type Specification struct {
		AppName int `name:"APP_NUMBER"`
	}

	var res Specification
	err := Process(&res)
	if err == nil {
		t.Error("Process expect to occur error. Because of APP_NUMBER is int but actual string")
	}
}

func TestProcess_AbNormal_3(t *testing.T) {
	os.Clearenv()

	type Specification struct {
	}

	var res Specification
	err := Process(res)
	if err == nil {
		t.Error("Process expect to occur error. Because of argument need to pass by pointer")
	}
}

func TestProcess_AbNormal_4(t *testing.T) {
	os.Clearenv()

	type Specification struct {
		IllegalType map[interface{}]interface{}
	}

	var res Specification
	err := Process(&res)
	if err == nil {
		t.Error("Process expect to occur error. Because of illegal input type")
	}
}
