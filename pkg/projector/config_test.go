package projector_test

import (
	"reflect"
	"testing"

	"github.com/Trijavico/projector/pkg/projector"
)

func getOpts(args []string) *projector.Opts{
    return &projector.Opts{
    	Args:       args,
    	PWD:        "",
    	ConfigPATH: "",
    }
}


func testConfig(t *testing.T, args, expectedArgs []string, operation projector.Operation){
    opts := getOpts(args)
    config, err := projector.NewConfig(opts)
    if err != nil{
        t.Errorf("Expected to get no error: %v", err)
    }

    if config.Operation != operation{
        t.Errorf("operation expected this: %v but got this: %v", operation, config.Operation)
    }

    if !reflect.DeepEqual(expectedArgs, config.Args){ 
        t.Errorf("expected args to be %+v array but got: %+v", expectedArgs, config.Args)
    }
}


func TestPrint(t *testing.T){
    testConfig(t, []string{}, []string{}, projector.PRINT)
}

func TestAdd(t *testing.T){
    testConfig(t, []string{"add", "key", "value"}, []string{"key", "value"}, projector.ADD)
}

func TestRemove(t *testing.T){
    testConfig(t, []string{"rm", "key"}, []string{"key"}, projector.REMOVE)
}

