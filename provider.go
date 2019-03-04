package main

import (
	"github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
  "github.com/Rizbe/go-npm"
)

//Provider Name
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NPM_USRNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NPM_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"npm_membership": resourceNPMUser(),
		},
    	ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
client := npm.NewBasicAuthClient(d.Get("username").(string), d.Get("password").(string))
	return client, nil
}
