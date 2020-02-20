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
	os.Setenv("APP_NUMBER_32", "2147483647")
	os.Setenv("APP_NUMBER_64", "9223372036854775807")
	os.Setenv("APP_DISABLE_DEFAULT", "dummy")

	type Specification struct {
		AppName           string `name:"APP_NAME"`
		IP                string
		AppNumber         int    `name:"APP_NUMBER" required:"true"`
		AppNumber32       int32  `name:"APP_NUMBER_32" required:"true"`
		AppNumber64       int64  `name:"APP_NUMBER_64" required:"true"`
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
		AppNumber32:       2147483647,
		AppNumber64:       9223372036854775807,
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

func TestProcess_AbNormal_5(t *testing.T) {
	os.Clearenv()
	os.Setenv("APP_NUMBER_32", "2147483648")

	type Specification struct {
		AppNumber32 int32 `name:"APP_NUMBER_32" required:"true"`
	}

	var res Specification
	err := Process(&res)
	if err == nil {
		t.Error("Process expect to occur error. Because of overflow as int32")
	}
}

func TestProcess_AbNormal_6(t *testing.T) {
	os.Clearenv()
	os.Setenv("APP_NUMBER", "TEST")

	type Specification struct {
		AppName float64 `name:"APP_NUMBER"`
	}

	var res Specification
	err := Process(&res)
	if err == nil {
		t.Error("Process expect to occur error. Because of APP_NUMBER is float64 but actual string")
	}
}
