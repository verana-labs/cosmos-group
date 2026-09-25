package testutil

import (
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	"cosmossdk.io/core/appconfig"

	"github.com/cosmos/cosmos-sdk/testutil/configurator"
	_ "github.com/cosmos/cosmos-sdk/x/auth"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/authz"          // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/bank"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/consensus"      // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/genutil"        // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/mint"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/params"         // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/staking"        // import as blank for app wiring

	groupmodulev1 "github.com/verana-labs/cosmos-group/api/cosmos/group/module/v1"
	_ "github.com/verana-labs/cosmos-group/module" // import as blank for app wiring
)

// GroupModule registers the group module in a configurator-built test app.
func GroupModule() configurator.ModuleOption {
	return func(config *configurator.Config) {
		config.ModuleConfigs["group"] = &appv1alpha1.ModuleConfig{
			Name:   "group",
			Config: appconfig.WrapAny(&groupmodulev1.Module{}),
		}
	}
}

var AppConfig = configurator.NewAppConfig(
	configurator.AuthModule(),
	configurator.BankModule(),
	configurator.StakingModule(),
	configurator.TxModule(),
	configurator.ConsensusModule(),
	configurator.ParamsModule(),
	configurator.GenutilModule(),
	GroupModule(),
)
