//
//
// @filename: error/errors.go
// @author: Krisna Pranav
// @license COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package errors

import (
	"errors"
)

var (
	ErrInvalidTargetAction      = errors.New("Error: Invalid target action")
	ErrBlockAdded               = errors.New("Error: Block has been already added")
	ErrBlockNodeInstantiated    = errors.New("Error: BlockNode instantiated")
	ErrTransporterNotConfigured = errors.New("Error: Transporter is not configured in config.yml file")
	ErrInvalidFileType          = errors.New("Error: invalid file type")
	ErrConfigDirMissing         = errors.New("Error: CONFIG_DIR is missing....")
)
