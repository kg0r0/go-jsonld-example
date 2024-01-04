package main

import (
	"fmt"

	"github.com/go-ap/jsonld"
)

// https://www.w3.org/TR/vc-data-model-2.0/#example-a-simple-example-of-the-contents-of-a-verifiable-credential
type Claims struct {
	Context           []string          `jsonld:"@context"`
	Id                string            `json:"id"`
	Type              []string          `json:"type"`
	Issuer            string            `json:"issuer"`
	ValidFrom         string            `json:"validFrom"`
	CredentialSubject CredentialSubject `json:"credentialSubject"`
}

type CredentialSubject struct {
	Id       string   `json:"id"`
	AlumniOf AlumniOf `json:"alumniOf"`
}

type AlumniOf struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	c := Claims{}
	err := jsonld.Unmarshal([]byte(`{"@context": ["https://www.w3.org/ns/credentials/v2", "https://www.w3.org/ns/credentials/examples/v2"], "id": "http://university.example/credentials/1872", "type": ["VerifiableCredential", "ExampleAlumniCredential"], "issuer": "https://university.example/issuers/565049", "validFrom": "2010-01-01T19:23:24Z", "credentialSubject": {"id": "did:example:ebfeb1f712ebc6f1c276e12ec21", "alumniOf": {"id": "did:example:c276e12ec21ebfeb1f712ebc6f1",  "name": "Example University"}}}`), &c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", c)
}
