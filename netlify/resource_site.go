package netlify

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/plumbing/operations"
)

func resourceSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceSiteCreate,
		Read:   resourceSiteRead,
		Update: resourceSiteUpdate,
		Delete: resourceSiteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"custom_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"deploy_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"repo": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"command": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"deploy_key_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"dir": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"provider": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"repo": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceSiteCreate(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)

	params := operations.NewCreateSiteParams()
	params.Site = &models.SiteSetup{
		Site: models.Site{
			Name:         d.Get("name").(string),
			CustomDomain: d.Get("custom_domain").(string),
		},
	}

	// If we have a repo config, then configure that
	if v, ok := d.GetOk("repo"); ok {
		vL := v.([]interface{})
		repo := vL[0].(map[string]interface{})

		params.Site.Repo = &models.RepoSetup{
			Branch:      repo["branch"].(string),
			Cmd:         repo["command"].(string),
			DeployKeyID: repo["deploy_key_id"].(string),
			Dir:         repo["dir"].(string),
			Provider:    repo["provider"].(string),
			Repo:        repo["repo"].(string),
		}
	}

	resp, err := meta.Netlify.Operations.CreateSite(params, meta.AuthInfo)
	if err != nil {
		return err
	}

	d.SetId(resp.Payload.ID)
	return resourceSiteRead(d, metaRaw)
}

func resourceSiteRead(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)
	params := operations.NewGetSiteParams()
	params.SiteID = d.Id()
	resp, err := meta.Netlify.Operations.GetSite(params, meta.AuthInfo)
	if err != nil {
		// If it is a 404 it was removed remotely
		if v, ok := err.(*operations.GetSiteDefault); ok && v.Code() == 404 {
			d.SetId("")
			return nil
		}

		return err
	}

	site := resp.Payload
	fmt.Println("foo")
	fmt.Printf("PAYLOAD: %#v", site)
	d.Set("name", site.Name)
	d.Set("custom_domain", site.CustomDomain)
	d.Set("deploy_url", site.DeployURL)

	return nil
}

func resourceSiteUpdate(d *schema.ResourceData, metaRaw interface{}) error {
	return nil
}

func resourceSiteDelete(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)
	params := operations.NewDeleteSiteParams()
	params.SiteID = d.Id()
	_, err := meta.Netlify.Operations.DeleteSite(params, meta.AuthInfo)
	return err
}
