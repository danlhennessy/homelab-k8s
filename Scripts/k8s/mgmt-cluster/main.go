package main

func main() {
	azCreds, _ := Get_Secret("azure-credentials")
	sv_principal, _ := Get_Secrets([]string{"aks-clientSecret", "aks-clientID"})

	Apply_String_Secret("crossplane-system", "azure-secret", "creds", azCreds)
	Apply_Multi_String_Secret("default", "sv-principal", sv_principal)
	// Apply_JSON_Secret("default", "json-test", azCreds)
}
