package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{

		Schema: map[string]*schema.Schema{
			"url": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The Octopus service URL",
			},
			"api_key": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The Octopus API key",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{

			"octopusdeploy_libraryvariableset": {

				Read: dataLibraryVariableSetRead,

				Schema: map[string]*schema.Schema{
					"id": {
						Required:    true,
						ForceNew:    true,
						Type:        schema.TypeString,
						Description: "Variable Set ID",
					},
					"content": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
	}
}

func dataLibraryVariableSetRead(d *schema.ResourceData, _ interface{}) error {
	id := d.Get("id").(string)
	content := fmt.Sprintf("<some json from %s>", id)
	checksum := sha1.Sum([]byte(content))

	d.Set("content",  content)
	d.SetId(hex.EncodeToString(checksum[:]))

	return nil
}
