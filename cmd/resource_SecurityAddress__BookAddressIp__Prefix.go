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
type xmlSecurityAddress__BookAddressIp__Prefix struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_address__book  struct {
			XMLName xml.Name `xml:"address-book"`
			V_name  *string  `xml:"name,omitempty"`
			V_address  struct {
				XMLName xml.Name `xml:"address"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_ip__prefix  *string  `xml:"ip-prefix,omitempty"`
			} `xml:"address"`
		} `xml:"security>address-book"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityAddress__BookAddressIp__PrefixCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_ip__prefix := d.Get("ip__prefix").(string)


	config := xmlSecurityAddress__BookAddressIp__Prefix{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_address__book.V_name = &V_name
	config.Groups.V_address__book.V_address.V_name__1 = &V_name__1
	config.Groups.V_address__book.V_address.V_ip__prefix = &V_ip__prefix

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityAddress__BookAddressIp__PrefixRead(d,m)
}

func junosSecurityAddress__BookAddressIp__PrefixRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityAddress__BookAddressIp__Prefix{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_address__book.V_name)
	d.Set("name__1", config.Groups.V_address__book.V_address.V_name__1)
	d.Set("ip__prefix", config.Groups.V_address__book.V_address.V_ip__prefix)

	return nil
}

func junosSecurityAddress__BookAddressIp__PrefixUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_ip__prefix := d.Get("ip__prefix").(string)


	config := xmlSecurityAddress__BookAddressIp__Prefix{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_address__book.V_name = &V_name
	config.Groups.V_address__book.V_address.V_name__1 = &V_name__1
	config.Groups.V_address__book.V_address.V_ip__prefix = &V_ip__prefix

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityAddress__BookAddressIp__PrefixRead(d,m)
}

func junosSecurityAddress__BookAddressIp__PrefixDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityAddress__BookAddressIp__Prefix() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityAddress__BookAddressIp__PrefixCreate,
		Read: junosSecurityAddress__BookAddressIp__PrefixRead,
		Update: junosSecurityAddress__BookAddressIp__PrefixUpdate,
		Delete: junosSecurityAddress__BookAddressIp__PrefixDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_address__book",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_address__book.V_address",
			},
			"ip__prefix": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_address__book.V_address. Numeric IPv4 or IPv6 address with prefix",
			},
		},
	}
}