# Changelog

## v0.54.0

Imported `x/group` and `proto/cosmos/group` from cosmos-sdk v0.53.8 (Apache-2.0),
including upstream fix GHSA-47ww-ff84-4jrg.

- Import path is `github.com/verana-labs/cosmos-group`; store key, prefixes, proto
  packages, type URLs and amino names are unchanged.
- Built against cosmos-sdk v0.54.4, `github.com/cosmos/cosmos-sdk/store/v2` and
  `cosmossdk.io/log/v2`.
- Added `api/cosmos/group/{v1,module/v1}` (pulsar) since `cosmossdk.io/api` v1
  no longer ships them; `module.proto` `go_package` points there.
- `testutil.GroupModule()` replaces the removed `configurator.GroupModule()`.
- `internal/orm` test support uses the store/v2 multistore API.
- `keeper.assertMetadataLength` uses `errorsmod.Wrap` (vet printf check).
