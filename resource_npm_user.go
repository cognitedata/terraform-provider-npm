package main

import (
	// "fmt"
	"github.com/Rizbe/go-npm"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceNPMUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceNPMUserCreateOrUpdate,
		Read:   resourceNPMUserRead,
		Update: resourceNPMUserCreateOrUpdate,
		Delete: resourceNPMUserDelete,

		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"org": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"role": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateValueFunc([]string{"developer", "admin", "owner"}),
			},
		},
	}
}

func resourceNPMUserCreateOrUpdate(d *schema.ResourceData, m interface{}) error {
	user := d.Get("user").(string)
	org := d.Get("org").(string)
	role := d.Get("role").(string)
	client := m.(*npm.Client)
	err := client.AddUser(org, user, role)
	if err != nil {
		// return resourceNPMUserRead(d, m)
	}

	d.SetId(user)
	return resourceNPMUserRead(d, m)

}

func resourceNPMUserRead(d *schema.ResourceData, m interface{}) error {
	user := d.Get("user").(string)
	org := d.Get("org").(string)

  client := m.(*npm.Client)
	t, err := client.GetUsers(org, user)
	if err != nil {
		d.SetId("")
		return nil
	}
	d.Set("user", user)
	d.Set("role", t[user])
	return nil
}

func resourceNPMUserDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
