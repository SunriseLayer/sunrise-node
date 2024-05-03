package share

import (
	"fmt"

	"github.com/sunriselayer/sunrise-da/nodebuilder/node"
	"github.com/sunriselayer/sunrise-da/share/availability/light"
	"github.com/sunriselayer/sunrise-da/share/eds"
	"github.com/sunriselayer/sunrise-da/share/p2p/discovery"
	"github.com/sunriselayer/sunrise-da/share/p2p/peers"
	"github.com/sunriselayer/sunrise-da/share/p2p/shrexeds"
	"github.com/sunriselayer/sunrise-da/share/p2p/shrexnd"
)

// TODO: some params are pointers and other are not, Let's fix this.
type Config struct {
	// EDSStoreParams sets eds store configuration parameters
	EDSStoreParams *eds.Parameters

	UseShareExchange bool
	// ShrExEDSParams sets shrexeds client and server configuration parameters
	ShrExEDSParams *shrexeds.Parameters
	// ShrExNDParams sets shrexnd client and server configuration parameters
	ShrExNDParams *shrexnd.Parameters
	// PeerManagerParams sets peer-manager configuration parameters
	PeerManagerParams peers.Parameters

	LightAvailability light.Parameters `toml:",omitempty"`
	Discovery         *discovery.Parameters
}

func DefaultConfig(tp node.Type) Config {
	cfg := Config{
		EDSStoreParams:    eds.DefaultParameters(),
		Discovery:         discovery.DefaultParameters(),
		ShrExEDSParams:    shrexeds.DefaultParameters(),
		ShrExNDParams:     shrexnd.DefaultParameters(),
		UseShareExchange:  true,
		PeerManagerParams: peers.DefaultParameters(),
	}

	if tp == node.Light {
		cfg.LightAvailability = light.DefaultParameters()
	}

	return cfg
}

// Validate performs basic validation of the config.
func (cfg *Config) Validate(tp node.Type) error {
	if tp == node.Light {
		if err := cfg.LightAvailability.Validate(); err != nil {
			return fmt.Errorf("nodebuilder/share: %w", err)
		}
	}

	if err := cfg.Discovery.Validate(); err != nil {
		return fmt.Errorf("nodebuilder/share: %w", err)
	}

	if err := cfg.ShrExNDParams.Validate(); err != nil {
		return fmt.Errorf("nodebuilder/share: %w", err)
	}

	if err := cfg.ShrExEDSParams.Validate(); err != nil {
		return fmt.Errorf("nodebuilder/share: %w", err)
	}

	if err := cfg.PeerManagerParams.Validate(); err != nil {
		return fmt.Errorf("nodebuilder/share: %w", err)
	}

	return nil
}
