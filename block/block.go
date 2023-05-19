//
//
// @filename: block/block.go
// @author: Krisna Pranav
// @license COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package block

import "github.com/NanoOfficial/micronano/common/types"

type IBlock interface {
	GetName() types.BlockName
	ActionsName() []types.ActionName
}
