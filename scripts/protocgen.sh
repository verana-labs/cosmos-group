#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/../proto"
buf generate --template buf.gen.pulsar.yaml
for f in $(find cosmos -name '*.proto' | grep -v '/module/'); do
  buf generate --template buf.gen.gogo.yaml "$f"
done
cd ..
cp -R github.com/verana-labs/cosmos-group/. . && rm -rf github.com
if grep -RIn 'github.com/cosmos/cosmos-sdk/x/group' --include='*.go' --include='*.proto' . ; then
  echo "stale upstream import paths" >&2; exit 1
fi
