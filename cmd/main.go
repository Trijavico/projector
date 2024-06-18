package main

import (
	"fmt"
	"log"

	"github.com/Trijavico/projector/pkg/projector"
)


type Data struct {
	Projector map[string]map[string]string `json:"projector"`
}

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
        return 
	}

    cliConfig, err := projector.NewConfig(opts)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

	fmt.Printf("%+v\n", cliConfig)

}

