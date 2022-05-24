// Copyright (c) 2017-2022, Juniper Networks Inc. All rights reserved.
//
// License: Apache 2.0
//
// THIS SOFTWARE IS PROVIDED BY Juniper Networks, Inc. ''AS IS'' AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL Juniper Networks, Inc. BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlForwarding__OptionsSamplingFamilyInetOutputFileFilename struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_file  struct {
			XMLName xml.Name `xml:"file"`
			V_filename  *string  `xml:"filename,omitempty"`
		} `xml:"forwarding-options>sampling>family>inet>output>file"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosForwarding__OptionsSamplingFamilyInetOutputFileFilenameCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_filename := d.Get("filename").(string)


	config := xmlForwarding__OptionsSamplingFamilyInetOutputFileFilename{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_file.V_filename = &V_filename

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosForwarding__OptionsSamplingFamilyInetOutputFileFilenameRead(d,m)
}

func junosForwarding__OptionsSamplingFamilyInetOutputFileFilenameRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlForwarding__OptionsSamplingFamilyInetOutputFileFilename{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("filename", config.Groups.V_file.V_filename)

	return nil
}

func junosForwarding__OptionsSamplingFamilyInetOutputFileFilenameUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_filename := d.Get("filename").(string)


	config := xmlForwarding__OptionsSamplingFamilyInetOutputFileFilename{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_file.V_filename = &V_filename

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosForwarding__OptionsSamplingFamilyInetOutputFileFilenameRead(d,m)
}

func junosForwarding__OptionsSamplingFamilyInetOutputFileFilenameDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosForwarding__OptionsSamplingFamilyInetOutputFileFilename() *schema.Resource {
	return &schema.Resource{
		Create: junosForwarding__OptionsSamplingFamilyInetOutputFileFilenameCreate,
		Read: junosForwarding__OptionsSamplingFamilyInetOutputFileFilenameRead,
		Update: junosForwarding__OptionsSamplingFamilyInetOutputFileFilenameUpdate,
		Delete: junosForwarding__OptionsSamplingFamilyInetOutputFileFilenameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"filename": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_file. Name of file to contain sampled packet dumps",
			},
		},
	}
}