package projector

import (
	"encoding/json"
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


func (p *Projector) SetValue() error{
    key := p.Config.Args[0]
    val := p.Config.Args[1]

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

    return nil
}

func (p *Projector) Remove() error{
    dir := p.Data.Projector[p.Config.PWD]
    if dir != nil {
        delete(dir, p.Config.Args[0])
    }

    err := p.save()
    if err != nil {
        return err
    }
    
    return nil
}
