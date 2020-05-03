package sysinit

import (
	"github.com/gorilla/securecookie"
)

var (
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey  = []byte("ykCG%Sg8E6#ow*Ff%u7^vojquz*33Sg^")
	blockKey = []byte("o3@XP5Z58vmGFW5qolI%hF^j@0@kILh8")
	SC       = securecookie.New(hashKey, blockKey)
)
