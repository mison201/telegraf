package mbr_gateway_config

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/Masterminds/sprig"
)

type Data struct {
	GatewayIPs []string
}

type Property struct {
	Name     string
	TypeName string
}

func processTemplate(fileName string, outputFile string, data Data) {
	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(fileName))

	var processed bytes.Buffer
	if err := tmpl.ExecuteTemplate(&processed, fileName, data); err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}

	currentData, err := os.ReadFile(outputFile)
	if err != nil {
		log.Fatal(err)
	}

	if !bytes.Equal(currentData, processed.Bytes()) {
		err = os.WriteFile(outputFile, processed.Bytes(), 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing file: ", outputFile)
	}
}
