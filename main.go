package main

import (
	"github.com/piprate/json-gold/ld"
)

func main() {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	expanded, err := proc.Expand("https://www.w3.org/ns/credentials/v2", options)
	if err != nil {
		panic(err)
	}

	doc := map[string]interface{}{
		"@context": []interface{}{
			"https://www.w3.org/ns/credentials/v2",
			"https://www.w3.org/ns/credentials/examples/v2",
		},
		"id": "http://university.example/credentials/1872",
		"type": []interface{}{
			"VerifiableCredential",
			"ExampleAlumniCredential",
		},
		"issuer":    "https://university.example/issuers/565049",
		"validFrom": "2010-01-01T19:23:24Z",
		"credentialSubject": map[string]interface{}{
			"id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
			"alumniOf": map[string]interface{}{
				"id":   "did:example:c276e12ec21ebfeb1f712ebc6f1",
				"name": "Example University",
			},
		},
	}
	expanded, err = proc.Expand(doc, options)
	if err != nil {
		panic(err)
	}
	ld.PrintDocument("JSON-LD expansion succeeded", expanded)
}
