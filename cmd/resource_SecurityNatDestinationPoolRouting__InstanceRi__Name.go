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
type xmlSecurityNatDestinationPoolRouting__InstanceRi__Name struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_pool  struct {
			XMLName xml.Name `xml:"pool"`
			V_name  *string  `xml:"name,omitempty"`
			V_routing__instance  struct {
				XMLName xml.Name `xml:"routing-instance"`
				V_ri__name  *string  `xml:"ri-name,omitempty"`
			} `xml:"routing-instance"`
		} `xml:"security>nat>destination>pool"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityNatDestinationPoolRouting__InstanceRi__NameCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_ri__name := d.Get("ri__name").(string)


	config := xmlSecurityNatDestinationPoolRouting__InstanceRi__Name{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_pool.V_name = &V_name
	config.Groups.V_pool.V_routing__instance.V_ri__name = &V_ri__name

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityNatDestinationPoolRouting__InstanceRi__NameRead(d,m)
}

func junosSecurityNatDestinationPoolRouting__InstanceRi__NameRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityNatDestinationPoolRouting__InstanceRi__Name{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_pool.V_name)
	d.Set("ri__name", config.Groups.V_pool.V_routing__instance.V_ri__name)

	return nil
}

func junosSecurityNatDestinationPoolRouting__InstanceRi__NameUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_ri__name := d.Get("ri__name").(string)


	config := xmlSecurityNatDestinationPoolRouting__InstanceRi__Name{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_pool.V_name = &V_name
	config.Groups.V_pool.V_routing__instance.V_ri__name = &V_ri__name

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityNatDestinationPoolRouting__InstanceRi__NameRead(d,m)
}

func junosSecurityNatDestinationPoolRouting__InstanceRi__NameDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityNatDestinationPoolRouting__InstanceRi__Name() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityNatDestinationPoolRouting__InstanceRi__NameCreate,
		Read: junosSecurityNatDestinationPoolRouting__InstanceRi__NameRead,
		Update: junosSecurityNatDestinationPoolRouting__InstanceRi__NameUpdate,
		Delete: junosSecurityNatDestinationPoolRouting__InstanceRi__NameDelete,

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
			"ri__name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_pool.V_routing__instance. Routing-instance name",
			},
		},
	}
}