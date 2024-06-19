package projector

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Data struct {
	Projector map[string]map[string]string `json:"projector"`
}

type Projector struct{
    Config *Config
    Data Data
}

func NewProjector(confing *Config) (*Projector, error){
    _, err := os.Stat(confing.ConfigPATH)
    if os.IsNotExist(err) {
        return &Projector{
        	Config: confing,
        	Data:   Data{
        		Projector: map[string]map[string]string{},
        	},
        }, nil 
    }

    jsonFile, err := os.ReadFile(confing.ConfigPATH)
    if err != nil{
        return nil, err
    }

    var data Data
    err = json.Unmarshal(jsonFile, &data)
    if err != nil{
        return nil, err
    }

    return &Projector{
    	Config: confing,
    	Data:   data,
    }, nil
}



func (p *Projector) save() error{
	configDir := filepath.Dir(p.Config.ConfigPATH)
    
    if _, err := os.Stat(configDir); os.IsNotExist(err){
        err := os.MkdirAll(configDir, 0755)
        if err != nil {
            return err
        }
    }

	jsonData, err := json.Marshal(p.Data)

	err = os.WriteFile(p.Config.ConfigPATH, jsonData, 0644)
	if err != nil {
		return err
	}

    fmt.Printf("Saving data: %s\n", string(jsonData))

    return nil
}

func (p *Projector) GetAllValues() (string, error){
    prevDir := ""
    currDir := p.Config.PWD 
    paths := make([]string, 0)

    for {
       prevDir = currDir
       paths = append(paths, currDir)
       currDir = filepath.Dir(currDir)

       if currDir != prevDir{
           break
       }
    }

    data := make(map[string]string)
    start := len(paths) - 1

    for i := start; i >= 0; i--{
        if p.Data.Projector[paths[i]] != nil{
            data = p.Data.Projector[paths[i]]
        }
    }

    jsonData, err := json.Marshal(data)
    if err != nil{
        return "", err
    }

    return string(jsonData), nil
}

func (p *Projector) GetValue(key string) string {
    prevDir := ""
    currDir := p.Config.PWD 
    value := p.Data.Projector[currDir][key]

    for currDir != prevDir{
       value = p.Data.Projector[currDir][key] 
       if value != ""{
           break
       }

       prevDir = currDir
       currDir = filepath.Dir(currDir)
    }

    return value 
}

func (p *Projector) SetValue(key, val string) error{
    if p.Data.Projector[p.Config.PWD] == nil {
        p.Data.Projector[p.Config.PWD] = make(map[string]string)
    }

    p.Data.Projector[p.Config.PWD][key] = val

    err := p.save()
    if err != nil {
        return err
    }

    return nil
}


func (p *Projector) Remove(key string) error{
    dir := p.Data.Projector[p.Config.PWD]
    if dir != nil {
        delete(dir, key)
    }

    err := p.save()
    if err != nil {
        return err
    }
    
    return nil
}
