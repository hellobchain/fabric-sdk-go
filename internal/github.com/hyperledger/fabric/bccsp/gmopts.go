package bccsp

const (
	SM4 = "SM4"

	SM3 = "SM3" //add GM

	SM2 = "SM2"

	SM = "SM"
)

type SM3Opts struct { // add GM
}

// Algorithm add GM
// Algorithm returns the hash algorithm identifier (to be used).
func (opts *SM3Opts) Algorithm() string {
	return SM3
}

// SM2KeyGenOpts contains options for SM2 key generation.
type SM2KeyGenOpts struct {
	Temporary bool
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *SM2KeyGenOpts) Algorithm() string {
	return SM2
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (opts *SM2KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

// SM2HSMKeyGenOpts contains options for SM2 key generation.
type SM2HSMKeyGenOpts struct {
	//KeyIndex  int
	Temporary bool
}

// func (opts *SM2HSMKeyGenOpts) GetKeyIndex() int {
// 	return opts.KeyIndex
// }

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *SM2HSMKeyGenOpts) Algorithm() string {
	return SM2
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (opts *SM2HSMKeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

// SM4KeyGenOpts contains options for SM4 key generation.
type SM4KeyGenOpts struct {
	Temporary bool
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *SM4KeyGenOpts) Algorithm() string {
	return SM4
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (opts *SM4KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
} // Add GM

// SM4ImportKeyOpts contains options for importing SM4 128 keys.
type SM4ImportKeyOpts struct {
	Temporary bool
}

// Algorithm returns the key importation algorithm identifier (to be used).
func (opts *SM4ImportKeyOpts) Algorithm() string {
	return SM4
}

// Ephemeral returns true if the key generated has to be ephemeral,
// false otherwise.
func (opts *SM4ImportKeyOpts) Ephemeral() bool {
	return opts.Temporary
}
