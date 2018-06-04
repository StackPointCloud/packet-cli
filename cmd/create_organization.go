// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

var (
	website string
	twitter string
	logo    string
)

// createOrganizationCmd represents the createOrganization command
var createOrganizationCmd = &cobra.Command{
	Use:   "organization",
	Short: "Command to create an organization.ß",
	Long: `Example:
	
	packet create organization -n [name]
	`,
	Run: func(cmd *cobra.Command, args []string) {
		req := &packngo.OrganizationCreateRequest{
			Name: name,
		}

		if description != "" {
			req.Description = description
		}

		if twitter != "" {
			req.Twitter = twitter
		}

		if logo != "" {
			req.Logo = logo
		}

		org, _, err := PacknGo.Organizations.Create(req)
		if err != nil {
			fmt.Println("Client error:", err)
		}

		data := make([][]string, 1)

		data[0] = []string{org.ID, org.Name, org.Created}
		header := []string{"ID", "Name", "Created"}

		output(org, header, &data)
	},
}

func init() {
	createCmd.AddCommand(createOrganizationCmd)

	createOrganizationCmd.Flags().StringVarP(&name, "name", "n", "", "--name or -n [name]")
	createOrganizationCmd.MarkFlagRequired("name")

	createOrganizationCmd.Flags().StringVarP(&description, "description", "d", "", "--description or -d [description]")
	createOrganizationCmd.Flags().StringVarP(&website, "website", "w", "", "--website or -w [website_url]")
	createOrganizationCmd.Flags().StringVarP(&twitter, "twitter", "t", "", "--twitter or -t [twitter_url]")
	createOrganizationCmd.Flags().StringVarP(&logo, "logo", "l", "", "--twitter or -t [logo_url]")
}