# go-jsonld-example

## Usage

```
$ go run main.go
JSON-LD expansion succeeded
[
  {
    "@id": "http://university.example/credentials/1872",
    "@type": [
      "https://www.w3.org/2018/credentials#VerifiableCredential",
      "https://www.w3.org/ns/credentials/examples#ExampleAlumniCredential"
    ],
    "https://www.w3.org/2018/credentials#credentialSubject": [
      {
        "@id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
        "https://www.w3.org/ns/credentials/examples#alumniOf": [
          {
            "@id": "did:example:c276e12ec21ebfeb1f712ebc6f1",
            "https://schema.org/name": [
              {
                "@value": "Example University"
              }
            ]
          }
        ]
      }
    ],
    "https://www.w3.org/2018/credentials#issuer": [
      {
        "@id": "https://university.example/issuers/565049"
      }
    ],
    "https://www.w3.org/2018/credentials#validFrom": [
      {
        "@type": "http://www.w3.org/2001/XMLSchema#dateTime",
        "@value": "2010-01-01T19:23:24Z"
      }
    ]
  }
]
```

## References
- https://github.com/piprate/json-gold
  - https://github.com/piprate/json-gold/issues/56
- https://www.w3.org/TR/vc-data-model-2.0/#example-a-simple-example-of-the-contents-of-a-verifiable-credential
- https://www.w3.org/TR/json-ld11/#the-context
