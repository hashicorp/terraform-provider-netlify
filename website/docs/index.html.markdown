---
layout: "netlify"
page_title: "Provider: Netlify"
sidebar_current: "docs-netlify-index"
description: |-
  The Netlify provider is used to deploy netlify resources
---

# Netlify Provider

Allows you to provision and deploy netlify sites and manage webhooks.

## Example Usage

```hcl
# Configure the Netlify Provider
provider "netlify" {
  token        = "${var.netlify_token}"
  base_url = "${var.netlify_base_url}"
}

# Create a new deploy key for this specific website
resource "netlify_deploy_key" "key" {}

# Define your site
resource "netlify_site" "main" {
  name = "my-site"

  repo {
    repo_branch = "master"
    command = "middleman build"
    deploy_key_id = "${netlify_deploy_key.key.id}"
    dir = "build"
    provider = "github"
    repo_path = "username/reponame"
  }
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `token` - (Required) Environment Variable: `NETLIFY_TOKEN`
* `base_url` - (Optional) Environment Variable: `NETLIFY_BASE_URL`
