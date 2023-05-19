//
//
// @filename: db/errors.go
// @author: Krisna Pranav
// @license COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package traffic

import "errors"

var (
	ErrNoNodeFound = errors.New("Error: Can't find any available nodes for action")
)
