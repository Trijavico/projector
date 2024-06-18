package main

import (
	"log"

	"github.com/Trijavico/projector/pkg/projector"
)



func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
        return 
	}

    config, err := projector.NewConfig(opts)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    proj, err := projector.NewProjector(config)
    if err != nil{
        log.Fatalf("Error: %v", err)
    }

    if proj.Config.Operation == projector.PRINT{
    }

    if proj.Config.Operation == projector.ADD{
        proj.SetValue()
    }

    if proj.Config.Operation == projector.REMOVE{
        proj.Remove()
    }
}

