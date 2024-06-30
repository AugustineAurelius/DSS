package ed25519

import (
	"encoding/hex"
	"fmt"
	"testing"
)

const (
	pubHex  = "75d6ff7afcb5033d1ca1a2db271d19342b69c221fd5ccfecb5da398b573f72b1"
	privHex = "0285e50f7c0b5f752369dc24992ffc580fb2300f4bf89ef4bc9a4a563a6fd76c75d6ff7afcb5033d1ca1a2db271d19342b69c221fd5ccfecb5da398b573f72b1"
)

func TestCreate(t *testing.T) {

	public, private := New()

	fmt.Println(hex.EncodeToString(public))
	fmt.Println(hex.EncodeToString(private))
}

func TestSign(t *testing.T) {

	public, private := New()

	msg := public
	signature := MustSign(private, msg)

	if !Verify(public, msg, signature) {
		t.Fatal("fail to virify msg")

	}

}
