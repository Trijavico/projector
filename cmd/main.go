package main

import (
	"fmt"

	"github.com/Trijavico/projector/pkg/projector"
)



func main() {
	opts, err := projector.GetOpts()
	if err != nil {
        fmt.Printf("Error: %v\n", err)
        return 
	}

    config, err := projector.NewConfig(opts)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return 
    }

    proj, err := projector.NewProjector(config)
    if err != nil{
        fmt.Printf("Error: %v\n", err)
        return 
    }

    switch operation := proj.Config.Operation; operation {

    case projector.PRINT:

        if len(config.Args) == 0 {

            values, err := proj.GetAllValues()
            if err != nil{
                fmt.Printf("Error: %v\n", err)
            }

            fmt.Println(values)

        }else{
            value := proj.GetValue(config.Args[0])
            fmt.Println(value)
        }

    case projector.ADD:
        proj.SetValue(config.Args[0], config.Args[1])

    case projector.REMOVE:
        proj.Remove(config.Args[0])
    }
}

