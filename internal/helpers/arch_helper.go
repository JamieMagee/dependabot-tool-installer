package helpers

type Arch int

const (
	x64 Arch = iota
	arm64
)

func FindArch() Arch {
	// This is a placeholder implementation
	return x64
}
