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

// Package writer provides convenience wrappers for writing NewRelic results
package writer

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	newrelic "github.com/sahellebusch/raider/newrelic"
)

// PolicyWriter contains a table writer to help write pilicies
type PolicyWriter struct {
	tbWriter *tablewriter.Table
}

// NewPolicyWriter creates a new PolicyWriter object
func NewPolicyWriter(writer io.Writer) *PolicyWriter {
	pr := &PolicyWriter{
		tbWriter: tablewriter.NewWriter(writer),
	}

	return pr
}

// WritePolicies writes the policies using the defined PolicyWriter
func (pw *PolicyWriter) WritePolicies(policies []newrelic.Policy) {
	for _, policy := range policies {
		createdAt := time.Unix((policy.CreatedAt / int64(time.Microsecond)), 0)
		updatedAt := time.Unix((policy.UpdatedAt / int64(time.Microsecond)), 0)

		pw.tbWriter.Append([]string{strconv.Itoa(policy.ID), policy.IncidentPreference, createdAt.String(), updatedAt.String()})
	}
	pw.tbWriter.SetHeader([]string{"Id", "Incident Preference", "Created", "Updated"})
	pw.tbWriter.Render()

	fmt.Println("\nFor more information about alert policy types, visit the link below.\nhttps://docs.newrelic.com/docs/alerts/new-relic-alerts/configuring-alert-policies/specify-when-new-relic-creates-incidents")

}
