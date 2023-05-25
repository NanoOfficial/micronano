package transporter

import "github.com/NanoOfficial/micronano/common/types"

type PayloadDiscovery struct {
	Blocks map[types.BlockName][]types.ActionName `json:"blocks"`
	Event  Event                                  `json:"event"`
}
