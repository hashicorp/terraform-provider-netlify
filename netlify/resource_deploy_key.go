package netlify

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/netlify/open-api/go/plumbing/operations"
)

func resourceDeployKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeployKeyCreate,
		Read:   resourceDeployKeyRead,
		Delete: resourceDeployKeyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDeployKeyCreate(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)

	resp, err := meta.Netlify.Operations.CreateDeployKey(
		operations.NewCreateDeployKeyParams(), meta.AuthInfo)
	if err != nil {
		return err
	}

	d.SetId(resp.Payload.ID)
	d.Set("public_key", resp.Payload.PublicKey)
	return nil
}

func resourceDeployKeyRead(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)
	params := operations.NewGetDeployKeyParams()
	params.KeyID = d.Id()
	resp, err := meta.Netlify.Operations.GetDeployKey(params, meta.AuthInfo)
	if err != nil {
		// Deleted remotely
		if v, ok := err.(*operations.GetDeployKeyDefault); ok && v.Code() == 404 {
			d.SetId("")
			return nil
		}

		return err
	}

	d.Set("public_key", resp.Payload.PublicKey)
	return nil
}

func resourceDeployKeyDelete(d *schema.ResourceData, metaRaw interface{}) error {
	meta := metaRaw.(*Meta)
	params := operations.NewDeleteDeployKeyParams()
	params.KeyID = d.Id()
	_, err := meta.Netlify.Operations.DeleteDeployKey(params, meta.AuthInfo)
	return err
}
