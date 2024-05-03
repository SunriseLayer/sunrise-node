package state

import (
	libfraud "github.com/celestiaorg/go-fraud"
	"github.com/celestiaorg/go-header/sync"
	apptypes "github.com/sunriselayer/sunrise/x/blob/types"

	"github.com/sunriselayer/sunrise-da/header"
	"github.com/sunriselayer/sunrise-da/nodebuilder/core"
	modfraud "github.com/sunriselayer/sunrise-da/nodebuilder/fraud"
	"github.com/sunriselayer/sunrise-da/share/eds/byzantine"
	"github.com/sunriselayer/sunrise-da/state"
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
