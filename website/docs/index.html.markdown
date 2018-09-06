---
layout: "netlify"
page_title: "Provider: Netlify"
sidebar_current: "docs-netlify-index"
description: |-
  The Netlify provider is used to deploy netlify resources
---

# Netlify Provider
[DESCRIPTION]


## Example Usage

```hcl
# Configure the Netlify Provider
provider "netlify" {
  token        = "${var.netlify_token}"
  base_url = "${var.netlify_base_url}"
}

# Define your site
resource "netlify_site" "your_website" {
  # ...
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `token` - (Required) Environment Variable: `NETLIFY_TOKEN`

* `base_url` - (Optional) Environment Variable: `NETLIFY_BASE_URL`
