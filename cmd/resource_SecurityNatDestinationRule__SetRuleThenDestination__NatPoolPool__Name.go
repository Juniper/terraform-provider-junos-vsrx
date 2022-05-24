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
type xmlSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__Name struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_rule__set  struct {
			XMLName xml.Name `xml:"rule-set"`
			V_name  *string  `xml:"name,omitempty"`
			V_rule  struct {
				XMLName xml.Name `xml:"rule"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_pool  struct {
					XMLName xml.Name `xml:"pool"`
					V_pool__name  *string  `xml:"pool-name,omitempty"`
				} `xml:"then>destination-nat>pool"`
			} `xml:"rule"`
		} `xml:"security>nat>destination>rule-set"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__NameCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_pool__name := d.Get("pool__name").(string)


	config := xmlSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__Name{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_rule__set.V_name = &V_name
	config.Groups.V_rule__set.V_rule.V_name__1 = &V_name__1
	config.Groups.V_rule__set.V_rule.V_pool.V_pool__name = &V_pool__name

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__NameRead(d,m)
}

func junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__NameRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__Name{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_rule__set.V_name)
	d.Set("name__1", config.Groups.V_rule__set.V_rule.V_name__1)
	d.Set("pool__name", config.Groups.V_rule__set.V_rule.V_pool.V_pool__name)

	return nil
}

func junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__NameUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_pool__name := d.Get("pool__name").(string)


	config := xmlSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__Name{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_rule__set.V_name = &V_name
	config.Groups.V_rule__set.V_rule.V_name__1 = &V_name__1
	config.Groups.V_rule__set.V_rule.V_pool.V_pool__name = &V_pool__name

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__NameRead(d,m)
}

func junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__NameDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__Name() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__NameCreate,
		Read: junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__NameRead,
		Update: junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__NameUpdate,
		Delete: junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__NameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_rule__set",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_rule__set.V_rule",
			},
			"pool__name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_rule__set.V_rule.V_pool. Name of Destination NAT pool",
			},
		},
	}
}