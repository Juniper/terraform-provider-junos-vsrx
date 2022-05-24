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
type xmlSystemRoot__AuthenticationEncrypted__Password struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_root__authentication  struct {
			XMLName xml.Name `xml:"root-authentication"`
			V_encrypted__password  *string  `xml:"encrypted-password,omitempty"`
		} `xml:"system>root-authentication"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemRoot__AuthenticationEncrypted__PasswordCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_encrypted__password := d.Get("encrypted__password").(string)


	config := xmlSystemRoot__AuthenticationEncrypted__Password{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_root__authentication.V_encrypted__password = &V_encrypted__password

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemRoot__AuthenticationEncrypted__PasswordRead(d,m)
}

func junosSystemRoot__AuthenticationEncrypted__PasswordRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemRoot__AuthenticationEncrypted__Password{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("encrypted__password", config.Groups.V_root__authentication.V_encrypted__password)

	return nil
}

func junosSystemRoot__AuthenticationEncrypted__PasswordUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_encrypted__password := d.Get("encrypted__password").(string)


	config := xmlSystemRoot__AuthenticationEncrypted__Password{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_root__authentication.V_encrypted__password = &V_encrypted__password

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSystemRoot__AuthenticationEncrypted__PasswordRead(d,m)
}

func junosSystemRoot__AuthenticationEncrypted__PasswordDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSystemRoot__AuthenticationEncrypted__Password() *schema.Resource {
	return &schema.Resource{
		Create: junosSystemRoot__AuthenticationEncrypted__PasswordCreate,
		Read: junosSystemRoot__AuthenticationEncrypted__PasswordRead,
		Update: junosSystemRoot__AuthenticationEncrypted__PasswordUpdate,
		Delete: junosSystemRoot__AuthenticationEncrypted__PasswordDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"encrypted__password": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_root__authentication. Encrypted password string",
			},
		},
	}
}