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
type xmlSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesName struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_security__zone  struct {
			XMLName xml.Name `xml:"security-zone"`
			V_name  *string  `xml:"name,omitempty"`
			V_system__services  struct {
				XMLName xml.Name `xml:"system-services"`
				V_name__1  *string  `xml:"name,omitempty"`
			} `xml:"host-inbound-traffic>system-services"`
		} `xml:"security>zones>security-zone"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesNameCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)


	config := xmlSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesName{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_security__zone.V_name = &V_name
	config.Groups.V_security__zone.V_system__services.V_name__1 = &V_name__1

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesNameRead(d,m)
}

func junosSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesNameRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesName{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_security__zone.V_name)
	d.Set("name__1", config.Groups.V_security__zone.V_system__services.V_name__1)

	return nil
}

func junosSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesNameUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)


	config := xmlSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesName{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_security__zone.V_name = &V_name
	config.Groups.V_security__zone.V_system__services.V_name__1 = &V_name__1

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesNameRead(d,m)
}

func junosSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesNameDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesName() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesNameCreate,
		Read: junosSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesNameRead,
		Update: junosSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesNameUpdate,
		Delete: junosSecurityZonesSecurity__ZoneHost__Inbound__TrafficSystem__ServicesNameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_security__zone",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_security__zone.V_system__services. ",
			},
		},
	}
}