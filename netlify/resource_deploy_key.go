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

		Schema: map[string]*schema.Schema{
			"public_key": &schema.Schema{
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

func resourceDeployKeyRead(d *schema.ResourceData, meta interface{}) error {
	// There is currently no read endpoint
	return nil
}

func resourceDeployKeyDelete(d *schema.ResourceData, meta interface{}) error {
	// Currently no delete endpoint, so just remove from state.
	return nil
}
