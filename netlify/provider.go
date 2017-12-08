package netlify

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETLIFY_TOKEN", nil),
				Description: descriptions["token"],
			},

			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETLIFY_BASE_URL", defaultBaseUrl),
				Description: descriptions["base_url"],
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"netlify_deploy_key": resourceDeployKey(),
			"netlify_hook":       resourceHook(),
			"netlify_site":       resourceSite(),
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ConfigureFunc: providerConfigure,
	}
}

// The default Netlify base URL.
const defaultBaseUrl = "https://api.netlify.com/api/v1"

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"token":    "The OAuth token used to connect to GitHub.",
		"base_url": "The Netlify Base API URL",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   d.Get("token").(string),
		BaseURL: d.Get("base_url").(string),
	}

	return config.Client()
}
