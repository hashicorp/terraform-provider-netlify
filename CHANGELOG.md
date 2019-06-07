## 0.2.0 (Unreleased)

### Features

* `resource/netlify_deploy_key`: support importing [GH-3]
* automatically retry requests [GH-2]

### Improvements

* upgrade to Go 1.11 [GH-12]
* use Go modules with `go mod vendor` [GH-16]
* fix indentation issues in Makefile [GH-23]
* upgrade `netlify/open-api` dependency to v0.11.4 [GH-20]
* upgrade `sirupsen/logrus` dependency to v1.4.2 [GH-27]

### Bug fixes

* docs fixes [GH-9] [GH-10] [GH-11] [GH-19]
* only set repo attribute on read if there is a repo path [GH-22]

## 0.1.0 (September 13, 2018)

Initial Release:

* `resource/netlify_build_hook`
* `resource/netlify_deploy_key`
* `resource/netlify_hook`
* `resource/netlify_site`
