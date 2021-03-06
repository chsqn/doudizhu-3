package algorithm

import "testing"

func TestAnalyseCards_ColorRecover(t *testing.T) {
	var cards AnalyseCards = 0

	//handcards:=[]byte{0xEF, 0xFF, 0x0a, 0x1a, 0x2a, 0x3a, 0x9, 0x19, 0x29, 0x39}
	//cards.Set([]byte{0xEF, 0xFF,0x9, 0x19, 0x39, 0x29})

	handcards := []byte{0xEF, 0xFF, 0x02, 0x12, 0x22, 0x32, 0x9, 0x19, 0x29}
	cards.Set([]byte{0xEF, 0xFF})

	for i := THREE; i <= 17; i++ {
		t.Logf("%v %v ", i, cards.Get(i))
	}

	t.Logf("%#v ", cards.ColorRecover(handcards))
}

func TestAnalyseCards_AnalyseCards(t *testing.T) {
	a := Cards{}
	//a.AnalyseUnSort( []byte{0x05, 0x15, 0x25,0x06, 0x16, 0x26,  0x27,0x37,0x07,  0x18, 0x28,0x08, 0x38, 0x2D, 0x2D})

	//t.Log(len(a.Cards))

	//handcards := []byte{0xEF, 0xFF, 0x0a, 0x1a, 0x2a, 0x3a, 0x9, 0x19, 0x29}
	handcards := []byte{0xEF, 0xFF, 0x9, 0x19, 0x39, 0x29}
	//discards := []byte{0x04}
	//a.AnalyseUnSort([]byte{0xEF,0xFF,0x02, 0x12, 0x03, 0x13, 0x23, 0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x36, 0x27, 0x37})
	a.AnalyseUnSort(handcards)

	//for k,v:=range a.analyseCards{
	//	t.Log(k,v)
	//}
	////
	t.Log(a.analyseCards, "====")
	for i := uint8(0); i <= 17; i++ {
		//t.Log(i, a.Get(i))
	}

	var cards AnalyseCards = 0

	//handcards:=[]byte{0xEF, 0xFF, 0x0a, 0x1a, 0x2a, 0x3a, 0x9, 0x19, 0x29, 0x39}
	cards.Set([]byte{0xEF, 0xFF, 0x9, 0x91, 0x39, 0x29})

	//for i:=THREE ;i<21;i++{
	//	t.Logf("%v %v ",i,cards.Get(i))
	//}

	//t.Log(a.Get(232))
}

func TestAnalyseStartEnd(t *testing.T) {
	handcards := []byte{0xEF, 0xFF, 0x9, 0x19, 0x39, 0x29}
	var an AnalyseCards
	an.Set(handcards)

	t.Log(an.Start(), an.End(),an.Len())
}
