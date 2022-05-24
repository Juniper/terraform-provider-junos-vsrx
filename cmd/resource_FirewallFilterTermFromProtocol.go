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
type xmlFirewallFilterTermFromProtocol struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_filter  struct {
			XMLName xml.Name `xml:"filter"`
			V_name  *string  `xml:"name,omitempty"`
			V_term  struct {
				XMLName xml.Name `xml:"term"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_from  struct {
					XMLName xml.Name `xml:"from"`
					V_protocol  *string  `xml:"protocol,omitempty"`
				} `xml:"from"`
			} `xml:"term"`
		} `xml:"firewall>filter"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosFirewallFilterTermFromProtocolCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_protocol := d.Get("protocol").(string)


	config := xmlFirewallFilterTermFromProtocol{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_filter.V_name = &V_name
	config.Groups.V_filter.V_term.V_name__1 = &V_name__1
	config.Groups.V_filter.V_term.V_from.V_protocol = &V_protocol

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosFirewallFilterTermFromProtocolRead(d,m)
}

func junosFirewallFilterTermFromProtocolRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlFirewallFilterTermFromProtocol{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_filter.V_name)
	d.Set("name__1", config.Groups.V_filter.V_term.V_name__1)
	d.Set("protocol", config.Groups.V_filter.V_term.V_from.V_protocol)

	return nil
}

func junosFirewallFilterTermFromProtocolUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_protocol := d.Get("protocol").(string)


	config := xmlFirewallFilterTermFromProtocol{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_filter.V_name = &V_name
	config.Groups.V_filter.V_term.V_name__1 = &V_name__1
	config.Groups.V_filter.V_term.V_from.V_protocol = &V_protocol

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosFirewallFilterTermFromProtocolRead(d,m)
}

func junosFirewallFilterTermFromProtocolDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosFirewallFilterTermFromProtocol() *schema.Resource {
	return &schema.Resource{
		Create: junosFirewallFilterTermFromProtocolCreate,
		Read: junosFirewallFilterTermFromProtocolRead,
		Update: junosFirewallFilterTermFromProtocolUpdate,
		Delete: junosFirewallFilterTermFromProtocolDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_filter",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_filter.V_term",
			},
			"protocol": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_filter.V_term.V_from. ",
			},
		},
	}
}