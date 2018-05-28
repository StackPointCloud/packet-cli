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
	"strconv"

	"github.com/spf13/cobra"
)

var (
	volumeID string
)

// retriveVolumeCmd represents the retriveVolume command
var retriveVolumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Gets volume list or volume details.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if projectID != "" && volumeID != "" {
			fmt.Println("Either projectID or volumeID can be set.")
			return
		} else if projectID == "" && volumeID == "" {
			fmt.Println("Either projectID or volumeID should be set.")
			return
		} else if projectID != "" {
			volumes, _, err := PacknGo.Volumes.List(projectID, nil)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
			data := make([][]string, len(volumes))

			for i, v := range volumes {
				data[i] = []string{v.ID, v.Name, strconv.Itoa(v.Size), v.State, v.Created}
			}
			header := []string{"ID", "Name", "Size", "State", "Created"}

			output(volumes, header, &data)
		} else if volumeID != "" {

			v, _, err := PacknGo.Volumes.Get(volumeID)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}

			header := []string{"ID", "Name", "Size", "State", "Created"}
			data := make([][]string, 1)
			data[0] = []string{v.ID, v.Name, strconv.Itoa(v.Size), v.State, v.Created}

			output(v, header, &data)
		}
	},
}

func init() {
	getCmd.AddCommand(retriveVolumeCmd)
	retriveVolumeCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "--project-id or -p [UUID]")
	retriveVolumeCmd.Flags().StringVarP(&volumeID, "id", "i", "", "--id or -i [UUID]")
}