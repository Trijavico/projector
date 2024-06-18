package projector

import (
	"fmt"
	"os"
	"path"
)


type Operation int

const (
	PRINT Operation = iota
	ADD
	REMOVE
)

type Config struct{
    Operation 
    Opts
}


func getConfig(opts *Opts) string{
	return path.Join(opts.ConfigPATH, "projector", "projector.json")
    
    /*
	configDir := filepath.Dir(config)

	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		return err
	}
	data := Data{
		Projector: map[string]map[string]string{
			getPWD(): {
				opts.Args[1]: opts.Args[2],
			},
		},
	}

	jsonData, err := json.Marshal(data)

	err = os.WriteFile(config, jsonData, 0644)
	if err != nil {
		return err
	}
    */
}

func getArgs(opts *Opts) ([]string, error){
    operation := getOperation(opts)

    if operation == PRINT{
        if len(opts.Args) > 1{
            return nil, fmt.Errorf("expected 0 arguments, but received %v", len(opts.Args) - 1)
        }

        return opts.Args, nil
    }

    if operation == ADD{
        if len(opts.Args) != 3{
            return nil, fmt.Errorf("expected 2 arguments, but received %v", len(opts.Args) - 1)
        }

        return opts.Args[1:], nil
    }

    if operation == REMOVE{
        if len(opts.Args) != 2{
            return nil, fmt.Errorf("expected 1 argument, but received %v", len(opts.Args) - 1)
        }
    }

    return opts.Args[1:], nil
}

func getPWD() (string, error){
	pwd, err := os.Getwd()
    if err != nil {
        return "", err
    }

	return pwd, nil
}

func getOperation(opts *Opts) Operation {
	if len(opts.Args) == 0 {
		return PRINT
	}

	switch op := opts.Args[0]; op {
	case "print":
		return PRINT
	case "add":
		return ADD
	case "remove":
		return REMOVE
	}

	return PRINT
}


func NewConfig(opts *Opts) (*Config, error){
    operartion := getOperation(opts)

    configPath := getConfig(opts)

    pwd, err := getPWD()
    if err != nil {
        return nil, err
    }

    args, err := getArgs(opts)
    if err != nil {
        return nil, err
    }

    
    return &Config{
    	Operation: operartion,
    	Opts: Opts{
    		Args:       args,
    		PWD:        pwd,
    		ConfigPATH: configPath,
    	},
    }, nil
}
