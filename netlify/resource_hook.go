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
			"site_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			"event": {
				Type:     schema.TypeString,
				Required: true,
			},

			"data": {
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}

func resourceHookCreate(d *schema.ResourceData, metaRaw interface{}) error {
	params := operations.NewCreateHookBySiteIDParams()
	params.SiteID = d.Get("site_id").(string)
	params.Hook = resourceHook_struct(d)

	meta := metaRaw.(*Meta)
	resp, err := meta.Netlify.Operations.CreateHookBySiteID(params, meta.AuthInfo)
	if err != nil {
		return err
	}

	d.SetId(resp.Payload.ID)
	return resourceHookRead(d, metaRaw)
}

func resourceHookRead(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)
	params := operations.NewGetHookParams()
	params.HookID = d.Id()
	resp, err := meta.Netlify.Operations.GetHook(params, meta.AuthInfo)
	if err != nil {
		// If it is a 404 it was removed remotely
		if v, ok := err.(*operations.GetHookDefault); ok && v.Code() == 404 {
			d.SetId("")
			return nil
		}

		return err
	}

	hook := resp.Payload
	d.Set("site_id", hook.SiteID)
	d.Set("type", hook.Type)
	d.Set("event", hook.Event)
	d.Set("data", hook.Data)

	return nil
}

func resourceHookUpdate(d *schema.ResourceData, metaRaw interface{}) error {
	params := operations.NewUpdateHookParams()
	params.HookID = d.Id()
	params.Hook = resourceHook_struct(d)

	meta := metaRaw.(*Meta)
	_, err := meta.Netlify.Operations.UpdateHook(params, meta.AuthInfo)
	if err != nil {
		return err
	}

	return resourceHookRead(d, metaRaw)
}

func resourceHookDelete(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)
	params := operations.NewDeleteHookBySiteIDParams()
	params.HookID = d.Id()
	_, err := meta.Netlify.Operations.DeleteHookBySiteID(params, meta.AuthInfo)
	return err
}

// Returns the Hook structure that can be used for creation or updating.
func resourceHook_struct(d *schema.ResourceData) *models.Hook {
	return &models.Hook{
		Data:  d.Get("data").(map[string]interface{}),
		Event: d.Get("event").(string),
		Type:  d.Get("type").(string),
	}
}
