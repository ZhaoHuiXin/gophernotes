// this file was generated by gomacro command: import _b "hash"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"hash"
)

// reflection: allow interpreted code to import "hash"
func init() {
	Packages["hash"] = Package{
	Types: map[string]Type{
		"Hash":	TypeOf((*hash.Hash)(nil)).Elem(),
		"Hash32":	TypeOf((*hash.Hash32)(nil)).Elem(),
		"Hash64":	TypeOf((*hash.Hash64)(nil)).Elem(),
	}, Proxies: map[string]Type{
		"Hash":	TypeOf((*P_hash_Hash)(nil)).Elem(),
		"Hash32":	TypeOf((*P_hash_Hash32)(nil)).Elem(),
		"Hash64":	TypeOf((*P_hash_Hash64)(nil)).Elem(),
	}, 
	}
}

// --------------- proxy for hash.Hash ---------------
type P_hash_Hash struct {
	Object	interface{}
	BlockSize_	func(interface{}) int
	Reset_	func(interface{}) 
	Size_	func(interface{}) int
	Sum_	func(_proxy_obj_ interface{}, b []byte) []byte
	Write_	func(_proxy_obj_ interface{}, p []byte) (n int, err error)
}
func (P *P_hash_Hash) BlockSize() int {
	return P.BlockSize_(P.Object)
}
func (P *P_hash_Hash) Reset()  {
	P.Reset_(P.Object)
}
func (P *P_hash_Hash) Size() int {
	return P.Size_(P.Object)
}
func (P *P_hash_Hash) Sum(b []byte) []byte {
	return P.Sum_(P.Object, b)
}
func (P *P_hash_Hash) Write(p []byte) (n int, err error) {
	return P.Write_(P.Object, p)
}

// --------------- proxy for hash.Hash32 ---------------
type P_hash_Hash32 struct {
	Object	interface{}
	BlockSize_	func(interface{}) int
	Reset_	func(interface{}) 
	Size_	func(interface{}) int
	Sum_	func(_proxy_obj_ interface{}, b []byte) []byte
	Sum32_	func(interface{}) uint32
	Write_	func(_proxy_obj_ interface{}, p []byte) (n int, err error)
}
func (P *P_hash_Hash32) BlockSize() int {
	return P.BlockSize_(P.Object)
}
func (P *P_hash_Hash32) Reset()  {
	P.Reset_(P.Object)
}
func (P *P_hash_Hash32) Size() int {
	return P.Size_(P.Object)
}
func (P *P_hash_Hash32) Sum(b []byte) []byte {
	return P.Sum_(P.Object, b)
}
func (P *P_hash_Hash32) Sum32() uint32 {
	return P.Sum32_(P.Object)
}
func (P *P_hash_Hash32) Write(p []byte) (n int, err error) {
	return P.Write_(P.Object, p)
}

// --------------- proxy for hash.Hash64 ---------------
type P_hash_Hash64 struct {
	Object	interface{}
	BlockSize_	func(interface{}) int
	Reset_	func(interface{}) 
	Size_	func(interface{}) int
	Sum_	func(_proxy_obj_ interface{}, b []byte) []byte
	Sum64_	func(interface{}) uint64
	Write_	func(_proxy_obj_ interface{}, p []byte) (n int, err error)
}
func (P *P_hash_Hash64) BlockSize() int {
	return P.BlockSize_(P.Object)
}
func (P *P_hash_Hash64) Reset()  {
	P.Reset_(P.Object)
}
func (P *P_hash_Hash64) Size() int {
	return P.Size_(P.Object)
}
func (P *P_hash_Hash64) Sum(b []byte) []byte {
	return P.Sum_(P.Object, b)
}
func (P *P_hash_Hash64) Sum64() uint64 {
	return P.Sum64_(P.Object)
}
func (P *P_hash_Hash64) Write(p []byte) (n int, err error) {
	return P.Write_(P.Object, p)
}