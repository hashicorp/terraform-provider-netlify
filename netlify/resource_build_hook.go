package netlify

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/plumbing/operations"
)

func resourceBuildHook() *schema.Resource {
	return &schema.Resource{
		Create: resourceBuildHookCreate,
		Read:   resourceBuildHookRead,
		Update: resourceBuildHookUpdate,
		Delete: resourceBuildHookDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"branch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceBuildHookCreate(d *schema.ResourceData, metaRaw interface{}) error {
	params := operations.NewCreateSiteBuildHookParams()
	params.SiteID = d.Get("site_id").(string)
	params.BuildHook = resourceBuildHook_struct(d)

	meta := metaRaw.(*Meta)
	resp, err := meta.Netlify.Operations.CreateSiteBuildHook(params, meta.AuthInfo)
	if err != nil {
		return err
	}

	d.SetId(resp.Payload.ID)
	return resourceBuildHookRead(d, metaRaw)
}

func resourceBuildHookRead(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)
	params := operations.NewGetSiteBuildHookParams()
	params.ID = d.Id()
	params.SiteID = d.Get("site_id").(string)
	resp, err := meta.Netlify.Operations.GetSiteBuildHook(params, meta.AuthInfo)
	if err != nil {
		// If it is a 404 it was removed remotely
		if v, ok := err.(*operations.GetSiteBuildHookDefault); ok && v.Code() == 404 {
			d.SetId("")
			return nil
		}

		return err
	}

	hook := resp.Payload
	d.Set("site_id", hook.SiteID)
	d.Set("branch", hook.Branch)
	d.Set("title", hook.Title)
	d.Set("url", hook.URL)

	return nil
}

func resourceBuildHookUpdate(d *schema.ResourceData, metaRaw interface{}) error {
	params := operations.NewUpdateSiteBuildHookParams()
	params.ID = d.Id()
	params.SiteID = d.Get("site_id").(string)
	params.BuildHook = resourceBuildHook_struct(d)

	meta := metaRaw.(*Meta)
	_, err := meta.Netlify.Operations.UpdateSiteBuildHook(params, meta.AuthInfo)
	if err != nil {
		return err
	}

	return resourceBuildHookRead(d, metaRaw)
}

func resourceBuildHookDelete(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)
	params := operations.NewDeleteSiteBuildHookParams()
	params.ID = d.Id()
	params.SiteID = d.Get("site_id").(string)
	_, err := meta.Netlify.Operations.DeleteSiteBuildHook(params, meta.AuthInfo)
	return err
}

// Returns the BuildHook structure that can be used for creation or updating.
func resourceBuildHook_struct(d *schema.ResourceData) *models.BuildHook {
	return &models.BuildHook{
		Branch: d.Get("branch").(string),
		Title:  d.Get("title").(string),
	}
}
