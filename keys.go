package gocliselectv2

import "github.com/buger/goterm"

//See http://www.climagic.org/mirrors/VT100_Escape_Codes.html

type KeyCode byte

const (
	unknownKey KeyCode = iota
	Key_a
	Key_b
	Key_c
	Key_d
	Key_e
	Key_f
	Key_g
	Key_h
	Key_i
	Key_j
	Key_k
	Key_l
	Key_m
	Key_n
	Key_o
	Key_p
	Key_q
	Key_r
	Key_s
	Key_t
	Key_u
	Key_v
	Key_w
	Key_x
	Key_y
	Key_z

	Key_A
	Key_B
	Key_C
	Key_D
	Key_E
	Key_F
	Key_G
	Key_H
	Key_I
	Key_J
	Key_K
	Key_L
	Key_M
	Key_N
	Key_O
	Key_P
	Key_Q
	Key_R
	Key_S
	Key_T
	Key_U
	Key_V
	Key_W
	Key_X
	Key_Y
	Key_Z

	Key_0
	Key_1
	Key_2
	Key_3
	Key_4
	Key_5
	Key_6
	Key_7
	Key_8
	Key_9

	upKey
	downKey
	//rightKey
	//leftKey
	escKey
	enterKey
	tabKey
	shiftTabKey
	pageUpKey
	pageDownKey
	homeKey
	endKey
	ctrlCKey
	ctrlDKey
)

const (
	a = byte(97)
	b = byte(98)
	c = byte(99)
	d = byte(100)
	e = byte(101)
	f = byte(102)
	g = byte(103)
	h = byte(104)
	i = byte(105)
	j = byte(106)
	k = byte(107)
	l = byte(108)
	m = byte(109)
	n = byte(110)
	o = byte(111)
	p = byte(112)
	q = byte(113)
	r = byte(114)
	s = byte(115)
	t = byte(116)
	u = byte(117)
	v = byte(118)
	w = byte(119)
	x = byte(120)
	y = byte(121)
	z = byte(122)

	aA = byte(65)
	bB = byte(66)
	cC = byte(67)
	dD = byte(68)
	eE = byte(69)
	fF = byte(70)
	gG = byte(71)
	hH = byte(72)
	iI = byte(73)
	jJ = byte(74)
	kK = byte(75)
	lL = byte(76)
	mM = byte(77)
	nN = byte(78)
	oO = byte(79)
	pP = byte(80)
	qQ = byte(81)
	rR = byte(82)
	sS = byte(83)
	tT = byte(84)
	uU = byte(85)
	vV = byte(86)
	wW = byte(87)
	xX = byte(88)
	yY = byte(89)
	zZ = byte(90)

	n0 = byte(48)
	n1 = byte(49)
	n2 = byte(50)
	n3 = byte(51)
	n4 = byte(52)
	n5 = byte(53)
	n6 = byte(54)
	n7 = byte(55)
	n8 = byte(56)
	n9 = byte(57)

	escape = byte(27)
	enter  = byte(13)
	tab    = byte(9)
	ctrlC  = byte(3)
	ctrlD  = byte(4)
)

var (
	up   = []byte{27, 91, 65}
	down = []byte{27, 91, 66}
	//right    = []byte{27, 91, 67}
	//left     = []byte{27, 91, 68}
	shiftTab = []byte{27, 91, 90}
	pageUp   = []byte{27, 91, 53}
	pageDown = []byte{27, 91, 54}
	home     = []byte{27, 91, 72}
	end      = []byte{27, 91, 70}
)

var inputToKeyMap1 = map[byte]KeyCode{
	a:      Key_a,
	b:      Key_b,
	c:      Key_c,
	d:      Key_d,
	e:      Key_e,
	f:      Key_f,
	g:      Key_g,
	h:      Key_h,
	i:      Key_i,
	j:      Key_j,
	k:      Key_k,
	l:      Key_l,
	m:      Key_m,
	n:      Key_n,
	o:      Key_o,
	p:      Key_p,
	q:      Key_q,
	r:      Key_r,
	s:      Key_s,
	t:      Key_t,
	u:      Key_u,
	v:      Key_v,
	w:      Key_w,
	x:      Key_x,
	y:      Key_y,
	z:      Key_z,
	aA:     Key_A,
	bB:     Key_B,
	cC:     Key_C,
	dD:     Key_D,
	eE:     Key_E,
	fF:     Key_F,
	gG:     Key_G,
	hH:     Key_H,
	iI:     Key_I,
	jJ:     Key_J,
	kK:     Key_K,
	lL:     Key_L,
	mM:     Key_M,
	nN:     Key_N,
	oO:     Key_O,
	pP:     Key_P,
	qQ:     Key_Q,
	rR:     Key_R,
	sS:     Key_S,
	tT:     Key_T,
	uU:     Key_U,
	vV:     Key_V,
	wW:     Key_W,
	xX:     Key_X,
	yY:     Key_Y,
	zZ:     Key_Z,
	n0:     Key_0,
	n1:     Key_1,
	n2:     Key_2,
	n3:     Key_3,
	n4:     Key_4,
	n5:     Key_5,
	n6:     Key_6,
	n7:     Key_7,
	n8:     Key_8,
	n9:     Key_9,
	escape: escKey,
	enter:  enterKey,
	tab:    tabKey,
	ctrlC:  ctrlCKey,
	ctrlD:  ctrlDKey,
}
var keyToInputMap1 = map[KeyCode]byte{}
var inputToKeyMap3 = map[byte]KeyCode{
	up[2]:   upKey,
	down[2]: downKey,
	//right[2]:    rightKey,
	//left[2]:     leftKey,
	shiftTab[2]: shiftTabKey,
	pageUp[2]:   pageUpKey,
	pageDown[2]: pageDownKey,
	home[2]:     homeKey,
	end[2]:      endKey,
}

type SelectedColor int

const (
	BLACK   SelectedColor = goterm.BLACK
	RED     SelectedColor = goterm.RED
	GREEN   SelectedColor = goterm.GREEN
	YELLOW  SelectedColor = goterm.YELLOW
	BLUE    SelectedColor = goterm.BLUE
	MAGENTA SelectedColor = goterm.MAGENTA
	CYAN    SelectedColor = goterm.CYAN
	WHITE   SelectedColor = goterm.WHITE
)

func init() {
	for k, v := range inputToKeyMap1 {
		if k == escape || k == enter || k == tab || k == ctrlC || k == ctrlD {
			continue
		}
		keyToInputMap1[v] = k
	}
}
