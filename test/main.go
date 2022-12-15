package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/ory/client-go"
)

func main() {
	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4434", // Kratos Admin API
		},
	}
	apiClient := client.NewAPIClient(configuration)
	adminCreateIdentityBody := *client.NewCreateIdentityBody(
		"default",
		map[string]interface{}{
			"email": "foo@example.com",
			"name": map[string]string{
				"first": "foo",
				"last":  "bar",
			},
		},
	) // AdminCreateIdentityBody |  (optional)

	createdIdentity, r, err := apiClient.IdentityApi.CreateIdentity(context.Background()).CreateIdentityBody(adminCreateIdentityBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.AdminCreateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminCreateIdentity`: Identity
	fmt.Println(createdIdentity)

}
