package kong

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"jwt": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				Description: "User supplied JWT token for authentication with Kong",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"kong_api":                            resourceKongAPI(),
			"kong_consumer":                       resourceKongConsumer(),
			"kong_api_plugin":                     resourceKongPlugin(),
			"kong_plugin":                         resourceKongPlugin(),
			"kong_consumer_basic_auth_credential": resourceKongBasicAuthCredential(),
			"kong_consumer_key_auth_credential":   resourceKongKeyAuthCredential(),
			"kong_consumer_jwt_credential":        resourceKongJWTCredential(),
			"kong_api_plugin_key_auth":            resourceKongKeyAuthPlugin(),
			"kong_consumer_acl_group":             resourceKongConsumerACLGroup(),
			"kong_certificate":                    resourceKongCertificate(),
			"kong_sni":                            resourceKongSNI(),
			"kong_upstream":                       resourceKongUpstream(),
			"kong_target":                         resourceKongTarget(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Address:  d.Get("address").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		JWT: d.Get("jwt").(string),
	}

	return config.Client()
}
