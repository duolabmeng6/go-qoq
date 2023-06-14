package mymodel

import (
	"strconv"
	"strings"
)

var (
	raw2key = map[uint16]string{ // https://github.com/wesbos/keycodes
		0:     "error",
		3:     "break",
		8:     "backspace",
		9:     "tab",
		12:    "clear",
		13:    "enter",
		16:    "shift",
		17:    "ctrl",
		18:    "alt",
		19:    "pause/break",
		20:    "caps lock",
		21:    "hangul",
		25:    "hanja",
		27:    "escape",
		28:    "conversion",
		29:    "non-conversion",
		32:    "spacebar",
		33:    "page up",
		34:    "page down",
		35:    "end",
		36:    "home",
		37:    "left arrow",
		38:    "up arrow",
		39:    "right arrow",
		40:    "down arrow",
		41:    "select",
		42:    "print",
		43:    "execute",
		44:    "Print Screen",
		45:    "insert",
		46:    "delete",
		47:    "help",
		48:    "0",
		49:    "1",
		50:    "2",
		51:    "3",
		52:    "4",
		53:    "5",
		54:    "6",
		55:    "7",
		56:    "8",
		57:    "9",
		58:    ":",
		59:    ";",
		60:    "<",
		61:    "=",
		63:    "ß",
		64:    "@",
		65:    "a",
		66:    "b",
		67:    "c",
		68:    "d",
		69:    "e",
		70:    "f",
		71:    "g",
		72:    "h",
		73:    "i",
		74:    "j",
		75:    "k",
		76:    "l",
		77:    "m",
		78:    "n",
		79:    "o",
		80:    "p",
		81:    "q",
		82:    "r",
		83:    "s",
		84:    "t",
		85:    "u",
		86:    "v",
		87:    "w",
		88:    "x",
		89:    "y",
		90:    "z",
		91:    "l-super",
		92:    "r-super",
		93:    "apps",
		95:    "sleep",
		96:    "numpad 0",
		97:    "numpad 1",
		98:    "numpad 2",
		99:    "numpad 3",
		100:   "numpad 4",
		101:   "numpad 5",
		102:   "numpad 6",
		103:   "numpad 7",
		104:   "numpad 8",
		105:   "numpad 9",
		106:   "multiply",
		107:   "add",
		108:   "numpad period",
		109:   "subtract",
		110:   "decimal point",
		111:   "divide",
		112:   "f1",
		113:   "f2",
		114:   "f3",
		115:   "f4",
		116:   "f5",
		117:   "f6",
		118:   "f7",
		119:   "f8",
		120:   "f9",
		121:   "f10",
		122:   "f11",
		123:   "f12",
		124:   "f13",
		125:   "f14",
		126:   "f15",
		127:   "f16",
		128:   "f17",
		129:   "f18",
		130:   "f19",
		131:   "f20",
		132:   "f21",
		133:   "f22",
		134:   "f23",
		135:   "f24",
		144:   "num lock",
		145:   "scroll lock",
		160:   "^",
		161:   "!",
		162:   "؛",
		163:   "#",
		164:   "$",
		165:   "ù",
		166:   "page backward",
		167:   "page forward",
		168:   "refresh",
		169:   "closing paren (AZERTY)",
		170:   "*",
		171:   "~ + * key",
		172:   "home key",
		173:   "minus (firefox), mute/unmute",
		174:   "decrease volume level",
		175:   "increase volume level",
		176:   "next",
		177:   "previous",
		178:   "stop",
		179:   "play/pause",
		180:   "e-mail",
		181:   "mute/unmute (firefox)",
		182:   "decrease volume level (firefox)",
		183:   "increase volume level (firefox)",
		186:   "semi-colon / ñ",
		187:   "equal sign",
		188:   "comma",
		189:   "dash",
		190:   "period",
		191:   "forward slash / ç",
		192:   "grave accent / ñ / æ / ö",
		193:   "?, / or °",
		194:   "numpad period (chrome)",
		219:   "open bracket",
		220:   "back slash",
		221:   "close bracket / å",
		222:   "single quote / ø / ä",
		223:   "`",
		224:   "left or right ⌘ key (firefox)",
		225:   "altgr",
		226:   "< /git >, left back slash",
		230:   "GNOME Compose Key",
		231:   "ç",
		233:   "XF86Forward",
		234:   "XF86Back",
		235:   "non-conversion",
		240:   "alphanumeric",
		242:   "hiragana/katakana",
		243:   "half-width/full-width",
		244:   "kanji",
		251:   "unlock trackpad (Chrome/Edge)",
		255:   "toggle touchpad",
		65517: "hyper",
	}

	keytoraw = map[string]uint16{
		"error":                           0,
		"break":                           3,
		"backspace":                       8,
		"tab":                             9,
		"clear":                           12,
		"enter":                           13,
		"shift":                           16,
		"ctrl":                            17,
		"alt":                             18,
		"pause/break":                     19,
		"caps lock":                       20,
		"hangul":                          21,
		"hanja":                           25,
		"escape":                          27,
		"conversion":                      28,
		"non-conversion":                  29,
		"spacebar":                        32,
		"page up":                         33,
		"page down":                       34,
		"end":                             35,
		"home":                            36,
		"left arrow":                      37,
		"up arrow":                        38,
		"right arrow":                     39,
		"down arrow":                      40,
		"select":                          41,
		"print":                           42,
		"execute":                         43,
		"Print Screen":                    44,
		"insert":                          45,
		"delete":                          46,
		"help":                            47,
		"0":                               48,
		"1":                               49,
		"2":                               50,
		"3":                               51,
		"4":                               52,
		"5":                               53,
		"6":                               54,
		"7":                               55,
		"8":                               56,
		"9":                               57,
		":":                               58,
		";":                               59,
		"<":                               60,
		"=":                               61,
		"ß":                               63,
		"@":                               64,
		"a":                               65,
		"b":                               66,
		"c":                               67,
		"d":                               68,
		"e":                               69,
		"f":                               70,
		"g":                               71,
		"h":                               72,
		"i":                               73,
		"j":                               74,
		"k":                               75,
		"l":                               76,
		"m":                               77,
		"n":                               78,
		"o":                               79,
		"p":                               80,
		"q":                               81,
		"r":                               82,
		"s":                               83,
		"t":                               84,
		"u":                               85,
		"v":                               86,
		"w":                               87,
		"x":                               88,
		"y":                               89,
		"z":                               90,
		"l-super":                         91,
		"r-super":                         92,
		"apps":                            93,
		"sleep":                           95,
		"numpad 0":                        96,
		"numpad 1":                        97,
		"numpad 2":                        98,
		"numpad 3":                        99,
		"numpad 4":                        100,
		"numpad 5":                        101,
		"numpad 6":                        102,
		"numpad 7":                        103,
		"numpad 8":                        104,
		"numpad 9":                        105,
		"multiply":                        106,
		"add":                             107,
		"numpad period":                   108,
		"subtract":                        109,
		"decimal point":                   110,
		"divide":                          111,
		"f1":                              112,
		"f2":                              113,
		"f3":                              114,
		"f4":                              115,
		"f5":                              116,
		"f6":                              117,
		"f7":                              118,
		"f8":                              119,
		"f9":                              120,
		"f10":                             121,
		"f11":                             122,
		"f12":                             123,
		"f13":                             124,
		"f14":                             125,
		"f15":                             126,
		"f16":                             127,
		"f17":                             128,
		"f18":                             129,
		"f19":                             130,
		"f20":                             131,
		"f21":                             132,
		"f22":                             133,
		"f23":                             134,
		"f24":                             135,
		"num lock":                        144,
		"scroll lock":                     145,
		"^":                               160,
		"!":                               161,
		"؛":                               162,
		"#":                               163,
		"$":                               164,
		"ù":                               165,
		"page backward":                   166,
		"page forward":                    167,
		"refresh":                         168,
		"closing paren (AZERTY)":          169,
		"*":                               170,
		"~ + * key":                       171,
		"home key":                        172,
		"minus (firefox), mute/unmute":    173,
		"decrease volume level":           174,
		"increase volume level":           175,
		"next":                            176,
		"previous":                        177,
		"stop":                            178,
		"play/pause":                      179,
		"e-mail":                          180,
		"mute/unmute (firefox)":           181,
		"decrease volume level (firefox)": 182,
		"increase volume level (firefox)": 183,
		"semi-colon / ñ":                  186,
		"equal sign":                      187,
		"comma":                           188,
		"dash":                            189,
		"period":                          190,
		"forward slash / ç":               191,
		"grave accent / ñ / æ / ö":        192,
		"?, / or °":                       193,
		"numpad period (chrome)":          194,
		"open bracket":                    219,
		"back slash":                      220,
		"close bracket / å":               221,
		"single quote / ø / ä":            222,
		"`":                               223,
		"left or right ⌘ key (firefox)":   224,
		"altgr":                           225,
		"< /git >, left back slash":       226,
		"GNOME Compose Key":               230,
		"ç":                               231,
		"XF86Forward":                     233,
		"XF86Back":                        234,
		"alphanumeric":                    240,
		"hiragana/katakana":               242,
		"half-width/full-width":           243,
		"kanji":                           244,
		"unlock trackpad (Chrome/Edge)":   251,
		"toggle touchpad":                 255,
		"hyper":                           65517,
	}

	mac_raw2key  = map[uint16]string{}
	mac_keytoraw = map[string]uint16{}
)

func Win键代码取名称(键代码 uint16) string {
	if 名称, 存在 := raw2key[键代码]; 存在 {
		return 名称
	}
	return ""
}
func E初始化mac键代码() {
	文本 := `-- 功能键
escape --> 53
F1 --> 122
F2 --> 120
F3 --> 99
F4 --> 118
F5 --> 96
F6 --> 97
F7 --> 98
F8 --> 100
F9 --> 101
F10 --> 109
F11 --> 103
F12 --> 111
F13 --> 105
F14 --> 107
F15 --> 113
F16 --> 106
F17 --> 64
F18 --> 79
F19 --> 80
F20 --> 90

keypad0 --> 82
keypad1 --> 83
keypad2 --> 84
keypad3 --> 85
keypad4 --> 86
keypad5 --> 87
keypad6 --> 88
keypad7 --> 89
keypad8 --> 91
keypad9 --> 92
keypadDecimal --> 65
keypadMultiply --> 67
keypadPlus --> 69
keypadClear --> 71
keypadDivide --> 75
keypadEnter --> 76
keypadMinus --> 78

0 --> 29
1 --> 18
2 --> 19
3 --> 20
4 --> 21
5 --> 23
6 --> 22
7 --> 26
8 --> 28
9 --> 25

tab --> 48
space --> 49
return --> 36
delete --> 51
forward delete --> 117
escape --> 53
command --> 55
shift --> 56
caps lock --> 57
option --> 58
control --> 59
right shift --> 60
right option --> 61
right control --> 62
function --> 63

a --> 0
b --> 11
c --> 8
d --> 2
e --> 14
f --> 3
g --> 5
h --> 4
i --> 34
j --> 38
k --> 40
l --> 37
m --> 46
n --> 45
o --> 31
p --> 35
q --> 12
r --> 15
s --> 1
t --> 17
u --> 32
v --> 9
w --> 13
x --> 7
y --> 16
z --> 6

grave --> 50
minus --> 27
equal --> 24
left bracket --> 33
right bracket --> 30
backslash --> 42
semicolon --> 41
quote --> 39
comma --> 43
period --> 47
slash --> 44

help --> 114
home --> 115
page up --> 116
page down --> 121
end --> 119
keypadEquals --> 81`
	// 把上述文本分割 加入 mac_raw2key
	// 清空 mac_raw2key
	// 清空 mac_keytoraw

	mac_raw2key = map[uint16]string{}
	mac_keytoraw = map[string]uint16{}

	行列表 := strings.Split(文本, "\n")
	for _, 行 := range 行列表 {
		行 = strings.TrimSpace(行)
		if len(行) == 0 {
			continue
		}
		//
		键名arr := strings.Split(行, " --> ")
		if len(键名arr) == 1 {
			continue
		}

		键名 := 键名arr[0]
		键代码, _ := strconv.Atoi(strings.Split(行, " --> ")[1])
		mac_raw2key[uint16(键代码)] = 键名
		mac_keytoraw[键名] = uint16(键代码)
	}

}
func Mac键代码取键名称(键代码 uint16) string {
	if 名称, 存在 := mac_raw2key[键代码]; 存在 {
		return 名称
	}
	return "null-" + strconv.Itoa(int(键代码))
}
func Mac键名称取键代码(键代码 uint16) string {
	if 名称, 存在 := mac_raw2key[键代码]; 存在 {
		return 名称
	}
	return ""
}
func ContainsAll(s []string, values ...string) bool {
	//先检查长度 以便提前退出
	if len(s) != len(values) {
		return false
	}
	for _, v := range values {
		//转换小写 以便比较
		v = strings.ToLower(v)
		found := false
		for _, item := range s {
			item = strings.ToLower(item)
			if v == item {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
