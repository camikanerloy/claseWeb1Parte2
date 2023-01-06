package product

import (
	"encoding/json"
	"io"
	"os"

	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
)

var Products []domain.Product

func OpenJsonFile() (jsonFile *os.File, err error) {
	jsonFile, err = os.Open("/Users/CKANER/bootcamp/claseWeb1Parte2/products.json")

	if err != nil {
		return
	}
	return
}

func GetProductsStruct() (err error) {

	jsonFile, err := OpenJsonFile()

	if err != nil {
		return
	}

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return
	}

	if err = json.Unmarshal(byteValue, &Products); err != nil {
		return
	}

	defer jsonFile.Close()
	return
}
