package transporter

type Tcp struct {
	// transporter.Transporter
	Transporter
}

func New() *Tcp {
	tcp := &Tcp{}

	return tcp
}
