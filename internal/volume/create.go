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

package volume

import (
	"strconv"

	"github.com/packethost/packet-cli/internal/output"
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func createCmd(svc packngo.VolumeService, out output.Outputer, facility, plan, projectID, billingCycle, description *string, size *int, locked *bool) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		req := &packngo.VolumeCreateRequest{
			BillingCycle: *billingCycle,
			PlanID:       *plan,
			FacilityID:   *facility,
			Size:         *size,
		}

		if *description != "" {
			req.Description = *description
		}
		if *locked {
			req.Locked = *locked
		}

		v, _, err := svc.Create(req, *projectID)
		if err != nil {
			return errors.Wrap(err, "Could not create Volume")
		}

		header := []string{"ID", "Name", "Size", "State", "Created"}
		data := make([][]string, 1)
		data[0] = []string{v.ID, v.Name, strconv.Itoa(v.Size), v.State, v.Created}

		return out.Output(v, header, &data)
	}
}

func Create(client *VolumeClient, out output.Outputer) *cobra.Command {
	var (
		size                                                 int
		facility, plan, projectID, billingCycle, description string
		locked                                               bool
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a volume",
		Example: `
packet volume create --size [size_in_GB] --plan [plan_UUID]--project-id [project_UUID] --facility [facility_code]`,
		RunE: createCmd(client.VolumeService, out, &facility, &plan, &projectID, &billingCycle, &description, &size, &locked),
	}
	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	cmd.Flags().StringVarP(&plan, "plan", "P", "", "Name of the plan")
	cmd.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility where the volume will be created")
	cmd.Flags().IntVarP(&size, "size", "s", 0, "Size in GB]")

	_ = cmd.MarkFlagRequired("size")
	_ = cmd.MarkFlagRequired("facility")
	_ = cmd.MarkFlagRequired("plan")
	_ = cmd.MarkFlagRequired("project-id")

	cmd.Flags().StringVarP(&billingCycle, "billing-cycle", "b", "hourly", "Billing cycle")
	cmd.Flags().StringVarP(&description, "description", "d", "", "Description of the volume")
	cmd.Flags().BoolVarP(&locked, "locked", "l", false, "Set the volume to be locked")
	return cmd
}