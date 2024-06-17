package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/Trijavico/projector/pkg/projector"
)

type Operation int

const (
	PRINT Operation = iota
	ADD
	REMOVE
)

type Data struct {
	Projector map[string]map[string]string `json:"projector"`
}

func main() {
	cliParser, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	cliParser.PWD, err = os.Getwd()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	if getOperation(cliParser.Args) == PRINT {
		fmt.Println("PRINT")
	}

	if getOperation(cliParser.Args) == ADD {
		err := getConfig(*cliParser)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}

	if getOperation(cliParser.Args) == REMOVE {
		fmt.Println("REMOVE")
	}

	fmt.Printf("%+v\n", cliParser)

}

func getConfig(opts projector.Opts) error {
	config := path.Join(opts.ConfigPATH, "projector", "projector.json")
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

	return nil
}

func getPWD() string {
	pwd, _ := os.Getwd()
	return pwd
}

func getOperation(args []string) Operation {
	if len(args) == 0 {
		return PRINT
	}

	switch op := args[0]; op {
	case "print":
		return PRINT
	case "add":
		return ADD
	case "remove":
		return REMOVE
	}

	return PRINT
}
