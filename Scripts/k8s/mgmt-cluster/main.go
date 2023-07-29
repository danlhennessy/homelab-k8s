package main

func main() {
	azCreds, _ := Get_Secret("azure-credentials")

	Apply_String_Secret("crossplane-system", "az-secret", "creds", azCreds)
	Apply_JSON_Secret("default", "json-test", azCreds)
}
