// Copyright © 2018 Alex Burck <me@alexburck.com>

package cmd

import (
	"log"

	"github.com/burck1/go-newrelic/models"

	"github.com/burck1/go-newrelic/client/users"

	"github.com/imdario/mergo"
	"github.com/spf13/cobra"
)

func makeUsersCmd(dst cobra.Command) *cobra.Command {
	src := cobra.Command{
		Use:     "users",
		Aliases: []string{"user"},
	}

	if err := mergo.Merge(&dst, src); err != nil {
		panic(err)
	}

	return &dst
}

var getUsersCmd = makeUsersCmd(cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		client := newClient(cmd)

		// AHH! doesn't return a list!
		// resp, err := client.Users.V2UsersJSONGet(users.NewV2UsersJSONGetParams())

		request := users.NewV2UsersJSONByIDGetParams()
		request.UserID = 2142609 // adam
		resp, err := client.Users.V2UsersJSONByIDGet(request)
		if err != nil {
			log.Fatal(err)
		}

		request2 := users.NewV2UsersJSONByIDGetParams()
		request2.UserID = 2141758 // alex
		resp2, err2 := client.Users.V2UsersJSONByIDGet(request2)
		if err2 != nil {
			log.Fatal(err2)
		}

		users := []models.UserResponseType{*resp.Payload.User, *resp2.Payload.User}
		return outputList(cmd, users)
	},
})

func init() {
	getCmd.AddCommand(getUsersCmd)
}
