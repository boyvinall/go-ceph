.PHONY: generate
generate: clean-generated
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i https://raw.githubusercontent.com/ceph/ceph/refs/heads/main/src/pybind/mgr/dashboard/openapi.yaml \
		-g go \
		-o /local \
		--skip-validate-spec \
		--git-host github.com \
		--git-user-id boyvinall \
		--git-repo-id go-ceph \
		--package-name ceph

.PHONY: clean-generated
clean-generated:
	cat .openapi-generator/FILES | xargs rm -rf
	find . -depth -type d -empty -delete
