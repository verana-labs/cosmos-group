# group

Apache-2.0 fork of the Cosmos SDK `x/group` module, maintained by Verana Labs.

## Why

Cosmos SDK v0.54 removed `x/group` from the open-source tree and moved it to an
enterprise package under a source-available evaluation license. Everything up to
SDK v0.53.x is Apache-2.0. This repository is `x/group` and `proto/cosmos/group`
as of `cosmos-sdk v0.53.8`, with git history preserved, built against SDK v0.54.x.

Store key, key prefixes, proto packages (`cosmos.group.v1`,
`cosmos.group.module.v1`), message type URLs, amino names and genesis format are
unchanged. Only the Go import path changed, so a chain can switch to this module
without a state migration.

Nothing in this repository is derived from the enterprise package. Do not read
or copy from `cosmos-sdk/enterprise/group`; its license claims derivatives.

## Versioning

`vMAJOR.MINOR` tracks the Cosmos SDK release family the module is built for.
`v0.54.x` builds against `github.com/cosmos/cosmos-sdk v0.54.x`. Patch releases
carry fixes to this module.

## Integration

```go
import (
	group "github.com/verana-labs/cosmos-group"
	groupkeeper "github.com/verana-labs/cosmos-group/keeper"
	groupmodulev1 "github.com/verana-labs/cosmos-group/api/cosmos/group/module/v1"
	_ "github.com/verana-labs/cosmos-group/module" // depinject registration
)
```

Migrating from the in-tree module is an import rewrite:

```bash
grep -rl 'github.com/cosmos/cosmos-sdk/x/group' --include='*.go' . \
  | xargs sed -i 's|github.com/cosmos/cosmos-sdk/x/group|github.com/verana-labs/cosmos-group|g'
sed -i 's|cosmossdk.io/api/cosmos/group|github.com/verana-labs/cosmos-group/api/cosmos/group|g' app/app_config.go
```

The `api/` package replaces `cosmossdk.io/api/cosmos/group`, which SDK v0.54
no longer ships.

## Layout

- root: `package group` (types, msgs, codec, generated `*.pb.go`)
- `keeper/`, `module/`, `client/cli/`, `simulation/`, `migrations/v2/`, `testutil/`
- `api/cosmos/group/...`: pulsar-generated types used by autocli and depinject
- `proto/`: source protos and buf configuration
- `SPEC.md`: upstream module specification

## Development

```
make build test lint vulncheck
make proto-gen   # needs buf, protoc-gen-gocosmos, protoc-gen-go-pulsar, protoc-gen-grpc-gateway, protoc-gen-go-grpc
```

`proto/buf.lock` pins a `buf.build/cosmos/cosmos-sdk` commit that does not ship
`cosmos/group/**`. After any `buf dep update`, check the new pin still has no
`cosmos/group` directory, otherwise the local protos clash with the dependency.

## Maintenance

- Track SDK release families: bump the minor version when moving to a new family.
- While `release/v0.53.x` is supported upstream, backport its `x/group` fixes.
- Report vulnerabilities privately through GitHub security advisories (see SECURITY.md).
