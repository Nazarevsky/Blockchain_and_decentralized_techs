package main

import (
	"sign/ecdsa"
	"sign/hashing"
)

func main() {
	ecdsa.InitValues()
	message := "Here is a transaction in Bavovnacoin."
	hash := hashing.SHA1(message)

	// The loop here is just to show that signing and verification work fine for different keys.
	for i := 0; i < 10; i++ {
		priv_key := ecdsa.GenPrivKey()
		pub_key := ecdsa.GenPubKey(priv_key)

		sign := ecdsa.Sign(hash, priv_key)
		ver := ecdsa.Verify(pub_key, sign, hash)

		if ver {
			println("Message verified successfuly!")
		} else {
			println("Message is not verified.")
		}
	}
}
