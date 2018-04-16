// Package crypto with Crypto interface
package crypto

import (
	v "github.com/nsheremet/esrp/value"
)

// Crypto is an interface for crypto engine
//
// Provides ciphersuites for calculating SRP values (SHA256, scrypt, HMAC f.e.).
// May provide one ciphersuite or construct different depends on options
// Also, different crypto providers may be implemented: OpenSSL, Libsodium etc.
type Crypto interface {

	// Interface function: SRP's one way hash function
	//
	// This is very important place for compatibility:
	//
	// 1. Different hashing algoritms may be involved. SHA, SHA256,
	// SHA512, blake2, their name is Legion.
	// 2. RFC5054 assumes that values for hashing is a byte arrays of hexadecimal
	// representation. Some implementations just concatenates hexadecimal strings.
	// So, it's #H's choice which representation should be picked from ESRP::Value
	//
	// Params:
	// - values {[]esrp.Value} values to be hashed
	//
	// Response:
	// - {esrp.Value} one-way hash function result
	H(values ...v.Value) v.Value

	// Interface function: password-based key derivation function
	//
	// PBKDF2, scrypt, bcrypt, argon2 or similar are recommended,
	// but usage of just SHA(salt | password) is possible.
	//
	// Params:
	// - salt     {esrp.Value} random generated salt (s)
	// - password {String}      plain-text password in UTF8 string or concatenated UTF8 string
	//
	// Response:
	// - {esrp.Value}
	PasswordHash(salt v.Value, password string) v.Value

	// Interface function: keyed hash transform function, like HMAC
	//
	// Params:
	// - key {esrp.Value}
	// - msg {esrp.Value}
	//
	// Response:
	// - {esrp.Value}
	KeyedHash(key v.Value, msg v.Value) v.Value

	// Interface function: random string generator
	//
	// Params:
	// - bytesLength {int} length of desired generated bytes
	//
	// Response:
	// - {esrp.Value}
	Random(bytesLength int) v.Value

	// Interface function: constant-time string comparison
	//
	// Compare two strings avoiding timing attacks
	//
	// Params:
	// - a {esrp.Value}
	// - b {esrp.Value}
	//
	// Response:
	// - {bool} true if strings are equal
	SecureCompare(a v.Value, b v.Value) bool
}
