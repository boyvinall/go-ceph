VERSION=0.0.1

.PHONY: all
all: generate

# be careful updating this because the file is not really correct, so we have some local changes aiming to push upstream
.PHONY: update-openapi
update-openapi:
	curl -LO https://raw.githubusercontent.com/ceph/ceph/refs/heads/main/src/pybind/mgr/dashboard/openapi.yaml

.PHONY: generate
generate: clean-generated
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:latest-release generate \
		-i /local/openapi.yaml \
		-g go \
		-o /local/api \
		--skip-validate-spec \
		--git-host github.com \
		--git-user-id boyvinall \
		--git-repo-id go-ceph \
		--package-name ceph \
		-t /local/templates \
		--additional-properties=withGoMod=false \
		--additional-properties=packageVersion=${VERSION} \
		--additional-properties=httpUserAgent=boyvinall/go-ceph/${VERSION} \
		--additional-properties=apiURL=https://ceph.example.com

.PHONY: clean-generated
clean-generated:
	# cd api && cat .openapi-generator/FILES | xargs rm -rf
	# find api -depth -type d -empty -delete
	rm -rf api

.PHONY: show-config-help
show-config-help:
	docker run --rm openapitools/openapi-generator-cli:latest-release config-help -g go

.PHONY: templates
templates:
	docker run --rm -v "${PWD}:/local"  openapitools/openapi-generator-cli:latest-release author template -g go --output /local/templates