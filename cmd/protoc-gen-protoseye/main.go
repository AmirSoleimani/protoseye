package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AmirSoleimani/protoseye/internal/protoparser"
	"google.golang.org/protobuf/compiler/protogen"
)

type Flags struct {
	// MARK: put your flags here
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	var fs flag.FlagSet

	protoOpts := protogen.Options{ParamFunc: fs.Set}
	protoOpts.
		Run(func(plugin *protogen.Plugin) error {
			for _, file := range plugin.Files {
				if !file.Generate {
					continue
				}

				for _, service := range file.Services {
					log.Println("processing:", file.Desc.Path())

					protoParser, err := protoparser.New()
					if err != nil {
						return err
					}

					// Generate OpenAPI spec for the service.
					rpcs, err := protoParser.RegisterService(service)
					if err != nil {
						return fmt.Errorf("An error occurred while processing the service: %v", err)
					}

					out := fmt.Sprintf("%s_%s", string(file.GoPackageName), string(service.Desc.Name()))

					for methodName, output := range rpcs {
						g := plugin.NewGeneratedFile(out+"_"+methodName+"_request.json", file.GoImportPath)
						if _, err := g.Write(output.MarshalRequest()); err != nil {
							return err
						}

						g = plugin.NewGeneratedFile(out+"_"+methodName+"_response.json", file.GoImportPath)
						if _, err := g.Write(output.MarshalResponse()); err != nil {
							return err
						}
					}
				}

			}
			return nil
		})
}
