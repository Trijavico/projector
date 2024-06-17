package projector

import (
	"os"

	"github.com/hellflame/argparse"
)

type Opts struct{
    Args []string
    PWD string
    ConfigPATH string
}

func GetOpts() (*Opts, error){
    parser := argparse.NewParser("projector", "gets all values", &argparse.ParserConfig{
        DisableDefaultShowHelp: true,
    })

    args := parser.Strings("a", "args", &argparse.Option{
		Positional: true,
		Required:   false,
		Default:    "",
	})

    configPATH, err := os.UserConfigDir()
    if err != nil{
        return nil, err
    }

	config := parser.String("c", "config", &argparse.Option{
		Required: false,
		Default: configPATH, 
	})

	pwd := parser.String("p", "pwd", &argparse.Option{
		Required: false,
		Default:  "",
	})

	err = parser.Parse(nil)
	if err != nil {
		return nil, err
	}

    return &Opts{
    	Args:       *args,
    	PWD:        *pwd,
    	ConfigPATH: *config,
    }, nil 
    
}
