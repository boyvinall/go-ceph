# go-ceph [![Go Reference](https://pkg.go.dev/badge/github.com/boyvinall/go-ceph/api.svg)](https://pkg.go.dev/github.com/boyvinall/go-ceph/api)

Go client for [Ceph](https://ceph.io/en/) [REST API](https://docs.ceph.com/en/latest/mgr/ceph_api/), generated from a slightly modified copy of the [OpenAPI spec](https://github.com/ceph/ceph/blob/main/src/pybind/mgr/dashboard/openapi.yaml).

> [!WARNING]
> The upstream openapi spec seems to be erroneous in some places. Some of those have been fixed up in this version. But there may be other errors which have not yet been found.

The intent here is to get these changes upstreamed at some point, when time allows.

## Docs

- [./api/README.md](./api/README.md)
- <https://pkg.go.dev/github.com/boyvinall/go-ceph/api>
- <https://docs.ceph.com/en/latest/mgr/ceph_api/>