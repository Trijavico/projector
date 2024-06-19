package projector_test

import (
	"encoding/json"
	"testing"

	"github.com/Trijavico/projector/pkg/projector"
)

func getData() projector.Data{
    return projector.Data{
    	Projector: map[string]map[string]string{
            "/": {
                "author": "name", 
                "secret": "greater",
            },
                
            "/foo":{
                "author": "name2",
            },
            "/foo/bar":{
                "author": "name3",
            },
        },
    }
}

func getProjector(pwd string) projector.Projector{
    return projector.Projector{
    	Config: &projector.Config{
    		Operation: projector.PRINT,
    		Opts:      projector.Opts{
    			Args:       []string{},
    			PWD:        pwd,
    			ConfigPATH: "path/to/config/file",
    		},
    	},
    	Data:   getData(),
    }
}

func TestGetAllValues(t *testing.T){
    proj := getProjector("/foo/bar")
    data := map[string]string{
        "secret": "greater",
        "author": "name3",
    }

    expected, _:= json.Marshal(data)
    value, err := proj.GetAllValues()
    if err != nil{
        t.Errorf("Should not error: %v", err)
    }

    if value != string(expected){
        t.Errorf("Expected: %s, but got: %s", string(expected), value)
    }
}

func TestGetValue(t *testing.T){
    proj := getProjector("/foo/bar")

    if value := proj.GetValue("author"); value != "name3"{
        t.Errorf("Expected: %s, but got: %s", "name3", value)
    }
    
    proj = getProjector("/foo")

    if value := proj.GetValue("author"); value != "name2"{
        t.Errorf("Expected: %s, but got: %s", "name2", value)
    }

    if value := proj.GetValue("secret"); value != "greater"{
        t.Errorf("Expected: %s, but got: %s", "greater", value)
    }
}


func TestSetValue(t *testing.T){
    proj := getProjector("/foo/bar") 

    proj.SetValue("author", "lastname")
    if value := proj.GetValue("author"); value != "lastname"{
        t.Errorf("Expected: %s, but got: %s", "lastname", value)
    }

    proj.SetValue("secret", "is_better")
    if value := proj.GetValue("secret"); value != "is_better"{
        t.Errorf("Expected: %s, but got: %s", "is_better", value)
    }

    proj.Config.PWD = "/"
    if value := proj.GetValue("secret"); value != "greater"{
        t.Errorf("Expected: %s, but got: %s", "greater", value)
    }
}


func TestRemoveValue(t *testing.T){
    proj := getProjector("/foo/bar")
    proj.Remove("secret")

    if value := proj.GetValue("secret"); value != "greater"{
        t.Errorf("Expected: %s, but got: %s", "greater", value)
    }

    proj.Remove("author")
    if value := proj.GetValue("author"); value != "name2"{
        t.Errorf("Expected: %s, but got: %s", "name2", value)
    }
}
