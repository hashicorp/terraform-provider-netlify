package netlify

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/plumbing/operations"
)

func resourceHook() *schema.Resource {
	return &schema.Resource{
		Create: resourceHookCreate,
		Read:   resourceHookRead,
		Update: resourceHookUpdate,
		Delete: resourceHookDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"event": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"data": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}

func resourceHookCreate(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)

	params := operations.NewCreateHookBySiteIDParams()
	params.SiteID = d.Get("site_id").(string)
	params.Hook = &models.Hook{
		Data:  d.Get("data").(map[string]interface{}),
		Event: d.Get("event").(string),
		Type:  d.Get("type").(string),
	}

	resp, err := meta.Netlify.Operations.CreateHookBySiteID(params, meta.AuthInfo)
	if err != nil {
		return err
	}

	d.SetId(resp.Payload.ID)
	return resourceHookRead(d, metaRaw)
}

func resourceHookRead(d *schema.ResourceData, metaRaw interface{}) error {
	// Not implemented yet
	return nil
}

func resourceHookUpdate(d *schema.ResourceData, metaRaw interface{}) error {
	// Not implemented yet
	return nil
}

func resourceHookDelete(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)
	params := operations.NewDeleteHookBySiteIDParams()
	params.HookID = d.Id()
	_, err := meta.Netlify.Operations.DeleteHookBySiteID(params, meta.AuthInfo)
	return err
}
