// this file was generated by gomacro command: import _b "crypto"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto"
	"io"
)

// reflection: allow interpreted code to import "crypto"
func init() {
	Packages["crypto"] = Package{
	Binds: map[string]Value{
		"MD4":	ValueOf(crypto.MD4),
		"MD5":	ValueOf(crypto.MD5),
		"MD5SHA1":	ValueOf(crypto.MD5SHA1),
		"RIPEMD160":	ValueOf(crypto.RIPEMD160),
		"RegisterHash":	ValueOf(crypto.RegisterHash),
		"SHA1":	ValueOf(crypto.SHA1),
		"SHA224":	ValueOf(crypto.SHA224),
		"SHA256":	ValueOf(crypto.SHA256),
		"SHA384":	ValueOf(crypto.SHA384),
		"SHA3_224":	ValueOf(crypto.SHA3_224),
		"SHA3_256":	ValueOf(crypto.SHA3_256),
		"SHA3_384":	ValueOf(crypto.SHA3_384),
		"SHA3_512":	ValueOf(crypto.SHA3_512),
		"SHA512":	ValueOf(crypto.SHA512),
		"SHA512_224":	ValueOf(crypto.SHA512_224),
		"SHA512_256":	ValueOf(crypto.SHA512_256),
	},Types: map[string]Type{
		"Decrypter":	TypeOf((*crypto.Decrypter)(nil)).Elem(),
		"DecrypterOpts":	TypeOf((*crypto.DecrypterOpts)(nil)).Elem(),
		"Hash":	TypeOf((*crypto.Hash)(nil)).Elem(),
		"PrivateKey":	TypeOf((*crypto.PrivateKey)(nil)).Elem(),
		"PublicKey":	TypeOf((*crypto.PublicKey)(nil)).Elem(),
		"Signer":	TypeOf((*crypto.Signer)(nil)).Elem(),
		"SignerOpts":	TypeOf((*crypto.SignerOpts)(nil)).Elem(),
	},Proxies: map[string]Type{
		"Decrypter":	TypeOf((*Decrypter_crypto)(nil)).Elem(),
		"DecrypterOpts":	TypeOf((*DecrypterOpts_crypto)(nil)).Elem(),
		"PrivateKey":	TypeOf((*PrivateKey_crypto)(nil)).Elem(),
		"PublicKey":	TypeOf((*PublicKey_crypto)(nil)).Elem(),
		"Signer":	TypeOf((*Signer_crypto)(nil)).Elem(),
		"SignerOpts":	TypeOf((*SignerOpts_crypto)(nil)).Elem(),
	},
	}
}

// --------------- proxy for crypto.Decrypter ---------------
type Decrypter_crypto struct {
	Object	interface{}
	Decrypt_	func(_proxy_obj_ interface{}, rand io.Reader, msg []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)
	Public_	func(interface{}) crypto.PublicKey
}
func (Proxy *Decrypter_crypto) Decrypt(rand io.Reader, msg []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error) {
	return Proxy.Decrypt_(Proxy.Object, rand, msg, opts)
}
func (Proxy *Decrypter_crypto) Public() crypto.PublicKey {
	return Proxy.Public_(Proxy.Object)
}

// --------------- proxy for crypto.DecrypterOpts ---------------
type DecrypterOpts_crypto struct {
	Object	interface{}
}

// --------------- proxy for crypto.PrivateKey ---------------
type PrivateKey_crypto struct {
	Object	interface{}
}

// --------------- proxy for crypto.PublicKey ---------------
type PublicKey_crypto struct {
	Object	interface{}
}

// --------------- proxy for crypto.Signer ---------------
type Signer_crypto struct {
	Object	interface{}
	Public_	func(interface{}) crypto.PublicKey
	Sign_	func(_proxy_obj_ interface{}, rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error)
}
func (Proxy *Signer_crypto) Public() crypto.PublicKey {
	return Proxy.Public_(Proxy.Object)
}
func (Proxy *Signer_crypto) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	return Proxy.Sign_(Proxy.Object, rand, digest, opts)
}

// --------------- proxy for crypto.SignerOpts ---------------
type SignerOpts_crypto struct {
	Object	interface{}
	HashFunc_	func(interface{}) crypto.Hash
}
func (Proxy *SignerOpts_crypto) HashFunc() crypto.Hash {
	return Proxy.HashFunc_(Proxy.Object)
}