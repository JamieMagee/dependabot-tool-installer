package helpers

type Distro struct {
	Name    string
	Version string
}

func ReadDistro() (Distro, error) {
	// This is a placeholder implementation
	return Distro{Name: "ubuntu", Version: "24.04"}, nil
}
