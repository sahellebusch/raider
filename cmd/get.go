/*
Copyright © 2019 Sean Hellebusch <sahellebusch@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Package cmd implements the commands for raider.
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sahellebusch/raider/newrelic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var validArgs []string = []string{"alert"}
var allAlertsFlag bool

func run(cmd *cobra.Command, args []string) {

	newRelic := newrelic.New(viper.Get("API_KEY").(string), strconv.Itoa(viper.Get("VERSION").(int)))

	err := cobra.OnlyValidArgs(cmd, args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if allAlertsFlag {
		alerts, error := newRelic.GetAllAlerts()
		if error != nil {
			fmt.Printf("ERROR: %s", error.Error())
			os.Exit(1)
		}

		fmt.Println(alerts)
	}
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run:       run,
	ValidArgs: validArgs,
	Args:      cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&allAlertsFlag, "all", "A", false, "Gets all alerts")
}
