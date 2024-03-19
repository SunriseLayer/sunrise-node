package light

import (
	"time"

	"github.com/sunrise-zone/sunrise-node/pruner"
)

// Window is the availability window for light nodes in the Celestia
// network (30 days).
const Window = pruner.AvailabilityWindow(time.Second * 86400 * 30)
