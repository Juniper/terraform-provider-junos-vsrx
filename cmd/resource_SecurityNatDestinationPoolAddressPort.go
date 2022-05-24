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
type xmlSecurityNatDestinationPoolAddressPort struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_pool  struct {
			XMLName xml.Name `xml:"pool"`
			V_name  *string  `xml:"name,omitempty"`
			V_address  struct {
				XMLName xml.Name `xml:"address"`
				V_port  *string  `xml:"port,omitempty"`
			} `xml:"address"`
		} `xml:"security>nat>destination>pool"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityNatDestinationPoolAddressPortCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_port := d.Get("port").(string)


	config := xmlSecurityNatDestinationPoolAddressPort{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_pool.V_name = &V_name
	config.Groups.V_pool.V_address.V_port = &V_port

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityNatDestinationPoolAddressPortRead(d,m)
}

func junosSecurityNatDestinationPoolAddressPortRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityNatDestinationPoolAddressPort{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_pool.V_name)
	d.Set("port", config.Groups.V_pool.V_address.V_port)

	return nil
}

func junosSecurityNatDestinationPoolAddressPortUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_port := d.Get("port").(string)


	config := xmlSecurityNatDestinationPoolAddressPort{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_pool.V_name = &V_name
	config.Groups.V_pool.V_address.V_port = &V_port

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityNatDestinationPoolAddressPortRead(d,m)
}

func junosSecurityNatDestinationPoolAddressPortDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityNatDestinationPoolAddressPort() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityNatDestinationPoolAddressPortCreate,
		Read: junosSecurityNatDestinationPoolAddressPortRead,
		Update: junosSecurityNatDestinationPoolAddressPortUpdate,
		Delete: junosSecurityNatDestinationPoolAddressPortDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_pool",
			},
			"port": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_pool.V_address. Specify the port value",
			},
		},
	}
}