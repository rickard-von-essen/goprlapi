// Defines all keycodes key envets see Parallels SDK - PrlKeys.h and PrlEnums.h
package key

/*
#cgo LDFLAGS: -framework ParallelsVirtualizationSDK
#include <ParallelsVirtualizationSDK/Parallels.h>
*/
import "C"

type KeyEvent int

func (e KeyEvent) String() string {
	switch {
	case e == PKE_PRESS:
		return "PKE_PRESS"
	case e == PKE_RELEASE:
		return "PKE_RELEASE"
	}
	return "Unknown PRL_KEY_EVENT"
}

// C.enum_PRL_KEY_EVENT

const (
	PKE_PRESS   KeyEvent = 0x0
	PKE_RELEASE KeyEvent = 0x80
)

type Key C.PRL_UINT32

const (
	PRL_KEY_INVALID             Key = 0
	PRL_KEY_WILDCARD_ANY        Key = 1
	PRL_KEY_WILDCARD_KEYBOARD   Key = 2
	PRL_KEY_WILDCARD_MOUSE      Key = 3
	PRL_KEY_WILDCARD_SHIFT      Key = 4
	PRL_KEY_WILDCARD_CTRL       Key = 5
	PRL_KEY_WILDCARD_ALT        Key = 6
	PRL_KEY_WILDCARD_WIN        Key = 7
	PRL_KEY_ESCAPE              Key = 9
	PRL_KEY_1                   Key = 10
	PRL_KEY_2                   Key = 11
	PRL_KEY_3                   Key = 12
	PRL_KEY_4                   Key = 13
	PRL_KEY_5                   Key = 14
	PRL_KEY_6                   Key = 15
	PRL_KEY_7                   Key = 16
	PRL_KEY_8                   Key = 17
	PRL_KEY_9                   Key = 18
	PRL_KEY_0                   Key = 19
	PRL_KEY_MINUS               Key = 20
	PRL_KEY_EQUAL               Key = 21
	PRL_KEY_BACKSPACE           Key = 22
	PRL_KEY_TAB                 Key = 23
	PRL_KEY_Q                   Key = 24
	PRL_KEY_W                   Key = 25
	PRL_KEY_E                   Key = 26
	PRL_KEY_R                   Key = 27
	PRL_KEY_T                   Key = 28
	PRL_KEY_Y                   Key = 29
	PRL_KEY_U                   Key = 30
	PRL_KEY_I                   Key = 31
	PRL_KEY_O                   Key = 32
	PRL_KEY_P                   Key = 33
	PRL_KEY_LEFT_BRACKET        Key = 34
	PRL_KEY_RIGHT_BRACKET       Key = 35
	PRL_KEY_ENTER               Key = 36
	PRL_KEY_LEFT_CONTROL        Key = 37
	PRL_KEY_A                   Key = 38
	PRL_KEY_S                   Key = 39
	PRL_KEY_D                   Key = 40
	PRL_KEY_F                   Key = 41
	PRL_KEY_G                   Key = 42
	PRL_KEY_H                   Key = 43
	PRL_KEY_J                   Key = 44
	PRL_KEY_K                   Key = 45
	PRL_KEY_L                   Key = 46
	PRL_KEY_SEMICOLON           Key = 47
	PRL_KEY_QUOTE               Key = 48
	PRL_KEY_TILDA               Key = 49
	PRL_KEY_LEFT_SHIFT          Key = 50
	PRL_KEY_BACKSLASH           Key = 51
	PRL_KEY_Z                   Key = 52
	PRL_KEY_X                   Key = 53
	PRL_KEY_C                   Key = 54
	PRL_KEY_V                   Key = 55
	PRL_KEY_B                   Key = 56
	PRL_KEY_N                   Key = 57
	PRL_KEY_M                   Key = 58
	PRL_KEY_COMMA               Key = 59
	PRL_KEY_DOT                 Key = 60
	PRL_KEY_SLASH               Key = 61
	PRL_KEY_RIGHT_SHIFT         Key = 62
	PRL_KEY_NP_STAR             Key = 63
	PRL_KEY_LEFT_ALT            Key = 64
	PRL_KEY_SPACE               Key = 65
	PRL_KEY_CAPS_LOCK           Key = 66
	PRL_KEY_F1                  Key = 67
	PRL_KEY_F2                  Key = 68
	PRL_KEY_F3                  Key = 69
	PRL_KEY_F4                  Key = 70
	PRL_KEY_F5                  Key = 71
	PRL_KEY_F6                  Key = 72
	PRL_KEY_F7                  Key = 73
	PRL_KEY_F8                  Key = 74
	PRL_KEY_F9                  Key = 75
	PRL_KEY_F10                 Key = 76
	PRL_KEY_NUM_LOCK            Key = 77
	PRL_KEY_SCROLL_LOCK         Key = 78
	PRL_KEY_NP_7                Key = 79
	PRL_KEY_NP_8                Key = 80
	PRL_KEY_NP_9                Key = 81
	PRL_KEY_NP_MINUS            Key = 82
	PRL_KEY_NP_4                Key = 83
	PRL_KEY_NP_5                Key = 84
	PRL_KEY_NP_6                Key = 85
	PRL_KEY_NP_PLUS             Key = 86
	PRL_KEY_NP_1                Key = 87
	PRL_KEY_NP_2                Key = 88
	PRL_KEY_NP_3                Key = 89
	PRL_KEY_NP_0                Key = 90
	PRL_KEY_NP_DELETE           Key = 91
	PRL_KEY_PRINT               Key = 92
	PRL_KEY_EUROPE_1            Key = 93
	PRL_KEY_EUROPE_2            Key = 94
	PRL_KEY_F11                 Key = 95
	PRL_KEY_F12                 Key = 96
	PRL_KEY_HOME                Key = 97
	PRL_KEY_UP                  Key = 98
	PRL_KEY_PAGE_UP             Key = 99
	PRL_KEY_LEFT                Key = 100
	PRL_KEY_RIGHT               Key = 102
	PRL_KEY_END                 Key = 103
	PRL_KEY_DOWN                Key = 104
	PRL_KEY_PAGE_DOWN           Key = 105
	PRL_KEY_INSERT              Key = 106
	PRL_KEY_DELETE              Key = 107
	PRL_KEY_NP_ENTER            Key = 108
	PRL_KEY_RIGHT_CONTROL       Key = 109
	PRL_KEY_PAUSE               Key = 110
	PRL_KEY_NP_SLASH            Key = 112
	PRL_KEY_RIGHT_ALT           Key = 113
	PRL_KEY_LEFT_WIN            Key = 115
	PRL_KEY_RIGHT_WIN           Key = 116
	PRL_KEY_MENU                Key = 117
	PRL_KEY_MEDIA_NEXT_TRACK    Key = 118
	PRL_KEY_MEDIA_PREV_TRACK    Key = 119
	PRL_KEY_MEDIA_STOP          Key = 120
	PRL_KEY_MEDIA_PLAY_PAUSE    Key = 121
	PRL_KEY_MUTE                Key = 122
	PRL_KEY_VOLUME_UP           Key = 123
	PRL_KEY_VOLUME_DOWN         Key = 124
	PRL_KEY_MEDIA_SELECT        Key = 125
	PRL_KEY_APP_MAIL            Key = 126
	PRL_KEY_APP_CALCULATOR      Key = 127
	PRL_KEY_APP_MY_COMPUTER     Key = 128
	PRL_KEY_WWW_SEARCH          Key = 129
	PRL_KEY_WWW_HOME            Key = 130
	PRL_KEY_WWW_BACK            Key = 131
	PRL_KEY_WWW_FORWARD         Key = 132
	PRL_KEY_WWW_STOP            Key = 133
	PRL_KEY_WWW_REFRESH         Key = 134
	PRL_KEY_WWW_FAVORITES       Key = 135
	PRL_KEY_EJECT               Key = 136
	PRL_KEY_SYSTEM_POWER        Key = 137
	PRL_KEY_SYSTEM_SLEEP        Key = 138
	PRL_KEY_SYSTEM_WAKE         Key = 139
	PRL_KEY_BRAZILIAN_KEYPAD    Key = 140
	PRL_KEY_RO                  Key = 141
	PRL_KEY_HIRAGANA_KATAKANA   Key = 142
	PRL_KEY_YEN                 Key = 143
	PRL_KEY_HENKAN              Key = 144
	PRL_KEY_MUHENKAN            Key = 145
	PRL_KEY_PC9800_KEYPAD       Key = 146
	PRL_KEY_HANGUEL             Key = 147
	PRL_KEY_HANJA               Key = 148
	PRL_KEY_KATAKANA            Key = 149
	PRL_KEY_HIRAGANA            Key = 150
	PRL_KEY_ZENKAKU_HANKAKU     Key = 151
	PRL_KEY_F13                 Key = 152
	PRL_KEY_F14                 Key = 153
	PRL_KEY_F15                 Key = 154
	PRL_KEY_F16                 Key = 155
	PRL_KEY_F17                 Key = 156
	PRL_KEY_F18                 Key = 157
	PRL_KEY_F19                 Key = 158
	PRL_KEY_F20                 Key = 159
	PRL_KEY_F21                 Key = 160
	PRL_KEY_F22                 Key = 161
	PRL_KEY_F23                 Key = 162
	PRL_KEY_F24                 Key = 163
	PRL_KEY_NP_EQUAL            Key = 164
	PRL_KEY_BREAK               Key = 165
	PRL_KEY_PRINT_WITH_MODIFIER Key = 166
	PRL_KEY_SYSRQ               Key = 167
	PRL_KEY_FN                  Key = 168
	PRL_KEY_EURO                Key = 169
	PRL_KEY_DOLLAR              Key = 170
	PRL_KEY_MON_BRIGHTNESS_DOWN Key = 171
	PRL_KEY_MON_BRIGHTNESS_UP   Key = 172
	PRL_KEY_APP_EXPOSE          Key = 173
	PRL_KEY_APP_DASHBOARD       Key = 174
	PRL_KEY_KBD_BRIGHTNESS_DOWN Key = 175
	PRL_KEY_KBD_BRIGHTNESS_UP   Key = 176
	PRL_KEY_MAC129              Key = 177
	PRL_KEY_LEFT_BUTTON         Key = 178
	PRL_KEY_MIDDLE_BUTTON       Key = 179
	PRL_KEY_RIGHT_BUTTON        Key = 180
	PRL_KEY_MOVE_UP_LEFT        Key = 181
	PRL_KEY_MOVE_UP             Key = 182
	PRL_KEY_MOVE_UP_RIGHT       Key = 183
	PRL_KEY_MOVE_LEFT           Key = 184
	PRL_KEY_MOVE_RIGHT          Key = 185
	PRL_KEY_MOVE_DOWN_LEFT      Key = 186
	PRL_KEY_MOVE_DOWN           Key = 187
	PRL_KEY_MOVE_DOWN_RIGHT     Key = 188
	PRL_KEY_WHEEL_UP            Key = 189
	PRL_KEY_WHEEL_DOWN          Key = 190
	PRL_KEY_WHEEL_LEFT          Key = 191
	PRL_KEY_WHEEL_RIGHT         Key = 192
	PRL_KEY_MAX                 Key = 193
)
