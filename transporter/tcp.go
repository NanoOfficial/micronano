package transporter

import (
	"github.com/NanoOfficial/micronano/transporter"
)

type Tcp struct {
	transporter.Transporter
}

func New() *Tcp {
	tcp := &Tcp{}

	return tcp
}
