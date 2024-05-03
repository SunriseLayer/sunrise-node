package state

import (
	libfraud "github.com/celestiaorg/go-fraud"
	"github.com/celestiaorg/go-header/sync"
	apptypes "github.com/sunriselayer/sunrise/x/blob/types"

	"github.com/sunrise-zone/sunrise-node/header"
	"github.com/sunrise-zone/sunrise-node/nodebuilder/core"
	modfraud "github.com/sunrise-zone/sunrise-node/nodebuilder/fraud"
	"github.com/sunrise-zone/sunrise-node/share/eds/byzantine"
	"github.com/sunrise-zone/sunrise-node/state"
)

// coreAccessor constructs a new instance of state.Module over
// a celestia-core connection.
func coreAccessor(
	corecfg core.Config,
	signer *apptypes.KeyringSigner,
	sync *sync.Syncer[*header.ExtendedHeader],
	fraudServ libfraud.Service[*header.ExtendedHeader],
) (*state.CoreAccessor, Module, *modfraud.ServiceBreaker[*state.CoreAccessor, *header.ExtendedHeader]) {
	ca := state.NewCoreAccessor(signer, sync, corecfg.IP, corecfg.RPCPort, corecfg.GRPCPort)

	return ca, ca, &modfraud.ServiceBreaker[*state.CoreAccessor, *header.ExtendedHeader]{
		Service:   ca,
		FraudType: byzantine.BadEncoding,
		FraudServ: fraudServ,
	}
}
