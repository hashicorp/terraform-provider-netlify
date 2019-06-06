## 0.2.0 (Unreleased)

### Features

* `resource/netlify_deploy_key`: support importing [GH-3]
* automatically retry requests [GH-2]

### Improvements

* upgrade to Go 1.11 [GH-12]
* use Go modules with `go mod vendor` [GH-16]
* fix indentation issues in Makefile [GH-23]
* upgrade `netlify/open-api` dependency [GH-20]

### Bug fixes

* docs fixes [GH-9] [GH-10] [GH-11] [GH-19]
* only set repo attribute on read if there is a repo path [GH-22]

## 0.1.0 (September 13, 2018)

Initial Release:

* `resource/netlify_build_hook`
* `resource/netlify_deploy_key`
* `resource/netlify_hook`
* `resource/netlify_site`
