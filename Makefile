VERSION=0.0.1

.PHONY: generate
generate: clean-generated
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:latest-release generate \
		-i https://raw.githubusercontent.com/ceph/ceph/refs/heads/main/src/pybind/mgr/dashboard/openapi.yaml \
		-g go \
		-o /local \
		--skip-validate-spec \
		--git-host github.com \
		--git-user-id boyvinall \
		--git-repo-id go-ceph \
		--package-name ceph \
		-t /local/templates \
		--additional-properties=withGoMod=false,packageVersion=${VERSION},httpUserAgent=boyvinall/go-ceph/${VERSION}

.PHONY: clean-generated
clean-generated:
	cat .openapi-generator/FILES | xargs rm -rf
	find . -depth -type d -empty -delete

.PHONY: show-config-help
show-config-help:
	docker run --rm openapitools/openapi-generator-cli:latest-release config-help -g go

.PHONY: templates
templates:
	docker run --rm -v "${PWD}:/local"  openapitools/openapi-generator-cli:latest-release author template -g go --output /local/templates