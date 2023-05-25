//
//
// @filename: transporter/error.go
// @author: Krisna Pranav
// @license COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package transporter

import (
	"errors"
)

var (
	ErrInvalidDeliveryMethod = errors.New("Error: invalid delivery method..")
)
