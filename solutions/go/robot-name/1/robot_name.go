package robotname

// Define the Robot type here.

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	r.name = "RX811"
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""

}
