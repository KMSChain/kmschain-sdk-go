package kmschain

//go:generate ../builder.sh

// Main KmsChain structure for having API functions referenced from it
type KMSChain struct {
  CPointer
  cm *CryptoMagic
  Keys KeyPair
}

// Clearing C/C++ allocations
// For example:
//    sc := kmschain.NewKmsChain()
//    defer sc.Clean()
func (sc *KMSChain) Clean() {
  sc.cm.Clean()
}

// Making new KMSChain object with default crypto configurations
func NewKmsChain() *KMSChain {
  cm := NewCryptoMagic()
  return &KMSChain{cm: cm, Keys: KeyPair{cm: cm}}
}

// Getting capsule object from given raw data converted from Capsule.ToBytes
func (sc *KMSChain) CapsuleFromBytes(capsuleData []byte) *Capsule {
  return CapsuleFromBytes(sc.cm, capsuleData)
}

// Getting private key object from given raw byte data converted with PrivateKey.ToBytes
func (sc *KMSChain) PrivateKeyFromBytes(skData []byte) *PrivateKey {
  return PrivateKeyFromBytes(sc.cm, skData)
}

// Getting public key object from given raw byte data converted with PublicKey.ToBytes
func (sc *KMSChain) PublicKeyFromBytes(pkData []byte) *PublicKey {
  return PublicKeyFromBytes(sc.cm, pkData)
}

// ReEncrypting capsule
func (sc *KMSChain) ReEncrypt(capsule *Capsule, rkk *ReEncryptionKey) *Capsule {
  return rkk.ReEncrypt(capsule)
}

func (sc *KMSChain) ReEncryptionKeyFromBytes(data []byte) *ReEncryptionKey {
  return ReEncryptionKeyFromBytes(sc.cm, data)
}
