---
layout: "netlify"
page_title: "Netlify: netlify_site"
sidebar_current: "docs-netlify-resource-site"
description: |-
  Provides an site resource.
---

# netlify_site

Primary settings for a Netlify site - should contain the bulk of your configuration. Allows configuration of most aspects of your Netlify site.

## Example Usage

```hcl
resource "netlify_site" "main" {
  name = "my-site"

  repo {
    command       = "middleman build"
    deploy_key_id = "${netlify_deploy_key.key.id}"
    dir           = "/build"
    provider      = "github"
    repo_path     = "username/repo"
    repo_branch   = "master"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) - Name of your site on Netlify (e.g. **mysite**.netlify.com)
* `repo` - (Required) - See [Repository](#repo)
* `custom_domain` - (Optional) - Custom domain of the site, must be configured using a CNAME in accordance with [Netlify's docs](https://www.netlify.com/docs/custom-domains). (e.g. `www.example.com`)
* `deploy_url` - (Optional)

### Repository

`repo` supports the following arguments:

* `command` - (Optional) - Shell command to run before deployment, typically used to build the site
* `deploy_key_id` - (Optional) - A deploy key id from the `deploy_key` resource
* `dir` - (Optional) - Directory to deploy, typically where the build puts the processed files
* `provider` - (Required) - Name of your VCS provider (e.g. `github`)
* `repo_path` - (Required) - path to your repo, typically `username/reponame`
* `repo_branch` - (Required) - branch to be deployed
