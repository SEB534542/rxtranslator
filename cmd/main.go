package main

import (
	"os"
	"fmt"
	"log"

	rxt "github.com/SEB534542/rxtranslator"
)
func main() {
	if len(os .Args) != 4 {
		log.Fatal("Usage: resume deepl-translator <API_KEY> <REACTIVE_RESUME_JSON_FILENAME> <TARGET_LANG>") 
		// https://developers.deepl.com/docs/api-reference/translate/openapi-spec-for-text-translation
		return
	}

	apiKey := os.Args[1]
	fname := os.Args[2]
	targetLang := os.Args[3]

	rsm, err := rxt.GetResume(fname)
	if err != nil {
		log.Fatalf("Error importing JSON: %s", err)
		return
	}

	rsmTranslated, err := rsm.Translate(apiKey, targetLang)
	if err != nil {
		log.Printf("Error translating Resume: %s", err)
	}

	fnameTranslated := fmt.Sprintf("%s %s.json", fname[:len(fname)-5], targetLang)
	err = rsmTranslated.Export(fnameTranslated)
	if err != nil {
		log.Print("Error exporting Resume:", err)
		fmt.Print(rsmTranslated)
		return
	}
	log.Print("Succesfully translated resume and exported to:", fmt.Sprintf("%s %s", fname, targetLang))
}
