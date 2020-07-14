package keypad

import (
	"github.com/bendahl/uinput"
	"github.com/minizbot2012/orbmap/interface/keyevents"
)

var vkm uinput.Keyboard

func init() {
	vkm, _ = uinput.CreateKeyboard("/dev/uinput", []byte("Orbmap"))
}
func procKey(kb chan *keyevents.KeyEvent) {
	for {
		KeyEv := <-kb
		if KeyEv.Type == 1 {
			if KeyEv.Value == 1 {
				vkm.KeyDown(KeyEv.Code)
			} else if KeyEv.Value == 2 {
			} else {
				vkm.KeyUp(KeyEv.Code)
			}
		}
	}
}
