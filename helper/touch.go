package helper

import (
	"encoding/binary"
	"log"
	"math"
	"os"
	"sync"
	"syscall"
)

var deviceMap sync.Map

func getDevice(deviceName string) (*os.File, error) {
	tmp, ok := deviceMap.Load(deviceName)
	var file *os.File
	var err error
	if !ok {
		file, err = os.OpenFile(deviceName, os.O_WRONLY, 0777)
		if err != nil {
			return file, err
		}
		deviceMap.Store(deviceName, file)
	} else {
		file = tmp.(*os.File)
	}
	return file, nil
}

type InputEvent struct {
	Time  syscall.Timeval // time in seconds since epoch at which event occurred
	Type  uint16          // event type - one of ecodes.EV_*
	Code  uint16          // event code related to the event type
	Value int32           // event value related to the event type
}

func sendEvent(type_ uint16, code uint16, value int32) {
	tf, _ := getDevice("/dev/input/event1")
	_ = binary.Write(tf, binary.LittleEndian, &InputEvent{
		Type:  type_,
		Code:  code,
		Value: value,
	})
}

var maxWidth = 1080
var maxHeight = 1920

func Init() {
	//easycmd.EasyCmdNotPty("")
}
func TouchDown(x float64, y float64, id int32) {
	x = math.Min(1, math.Max(0, x))
	y = math.Min(1, math.Max(0, y))
	sendEvent(EV_ABS, ABS_MT_TRACKING_ID, id)
	sendEvent(EV_KEY, BTN_TOUCH, 0x00000001)
	sendEvent(EV_KEY, BTN_TOOL_FINGER, 0x00000001)
	sendEvent(EV_ABS, ABS_MT_POSITION_X, int32(x*float64(maxWidth)))
	sendEvent(EV_ABS, ABS_MT_POSITION_Y, int32(y*float64(maxHeight)))
	sendEvent(EV_ABS, ABS_MT_TOUCH_MAJOR, 5)
	sendEvent(EV_SYN, SYN_REPORT, 0x00000000)
}

func TouchUp(id int32) {
	sendEvent(EV_ABS, ABS_MT_TRACKING_ID, id)
	sendEvent(EV_KEY, BTN_TOUCH, 0x00000000)
	sendEvent(EV_KEY, BTN_TOOL_FINGER, 0x00000000)
	sendEvent(EV_SYN, SYN_REPORT, 0x00000000)
}

func TouchMove(x float64, y float64, id int32) {
	x = math.Min(1, math.Max(0, x))
	y = math.Min(1, math.Max(0, y))
	log.Println("TouchMove:", x, ",", y)
	sendEvent(EV_ABS, ABS_MT_TRACKING_ID, id)
	sendEvent(EV_ABS, ABS_MT_POSITION_X, int32(x*float64(maxWidth)))
	sendEvent(EV_ABS, ABS_MT_POSITION_Y, int32(y*float64(maxHeight)))
	sendEvent(EV_SYN, SYN_REPORT, 0x00000000)
}

/*
 * Device properties and quirks
 */

const INPUT_PROP_POINTER = 0x00        /* needs a pointer */
const INPUT_PROP_DIRECT = 0x01         /* direct input devices */
const INPUT_PROP_BUTTONPAD = 0x02      /* has button(s) under pad */
const INPUT_PROP_SEMI_MT = 0x03        /* touch rectangle only */
const INPUT_PROP_TOPBUTTONPAD = 0x04   /* softbuttons at top of pad */
const INPUT_PROP_POINTING_STICK = 0x05 /* is a pointing stick */
const INPUT_PROP_ACCELEROMETER = 0x06  /* has accelerometer */

const INPUT_PROP_MAX = 0x1f
const INPUT_PROP_CNT = (INPUT_PROP_MAX + 1)

/*
 * Event types
 */

const EV_SYN = 0x00
const EV_KEY = 0x01
const EV_REL = 0x02
const EV_ABS = 0x03
const EV_MSC = 0x04
const EV_SW = 0x05
const EV_LED = 0x11
const EV_SND = 0x12
const EV_REP = 0x14
const EV_FF = 0x15
const EV_PWR = 0x16
const EV_FF_STATUS = 0x17
const EV_MAX = 0x1f
const EV_CNT = (EV_MAX + 1)

/*
 * Synchronization events.
 */

const SYN_REPORT = 0
const SYN_CONFIG = 1
const SYN_MT_REPORT = 2
const SYN_DROPPED = 3
const SYN_MAX = 0xf
const SYN_CNT = (SYN_MAX + 1)

/*
 * Keys and buttons
 *
 * Most of the keys/buttons are modeled after USB HUT 1.12
 * (see http://www.usb.org/developers/hidpage).
 * Abbreviations in the comments:
 * AC - Application Control
 * AL - Application Launch Button
 * SC - System Control
 */

const KEY_RESERVED = 0
const KEY_ESC = 1
const KEY_1 = 2
const KEY_2 = 3
const KEY_3 = 4
const KEY_4 = 5
const KEY_5 = 6
const KEY_6 = 7
const KEY_7 = 8
const KEY_8 = 9
const KEY_9 = 10
const KEY_0 = 11
const KEY_MINUS = 12
const KEY_EQUAL = 13
const KEY_BACKSPACE = 14
const KEY_TAB = 15
const KEY_Q = 16
const KEY_W = 17
const KEY_E = 18
const KEY_R = 19
const KEY_T = 20
const KEY_Y = 21
const KEY_U = 22
const KEY_I = 23
const KEY_O = 24
const KEY_P = 25
const KEY_LEFTBRACE = 26
const KEY_RIGHTBRACE = 27
const KEY_ENTER = 28
const KEY_LEFTCTRL = 29
const KEY_A = 30
const KEY_S = 31
const KEY_D = 32
const KEY_F = 33
const KEY_G = 34
const KEY_H = 35
const KEY_J = 36
const KEY_K = 37
const KEY_L = 38
const KEY_SEMICOLON = 39
const KEY_APOSTROPHE = 40
const KEY_GRAVE = 41
const KEY_LEFTSHIFT = 42
const KEY_BACKSLASH = 43
const KEY_Z = 44
const KEY_X = 45
const KEY_C = 46
const KEY_V = 47
const KEY_B = 48
const KEY_N = 49
const KEY_M = 50
const KEY_COMMA = 51
const KEY_DOT = 52
const KEY_SLASH = 53
const KEY_RIGHTSHIFT = 54
const KEY_KPASTERISK = 55
const KEY_LEFTALT = 56
const KEY_SPACE = 57
const KEY_CAPSLOCK = 58
const KEY_F1 = 59
const KEY_F2 = 60
const KEY_F3 = 61
const KEY_F4 = 62
const KEY_F5 = 63
const KEY_F6 = 64
const KEY_F7 = 65
const KEY_F8 = 66
const KEY_F9 = 67
const KEY_F10 = 68
const KEY_NUMLOCK = 69
const KEY_SCROLLLOCK = 70
const KEY_KP7 = 71
const KEY_KP8 = 72
const KEY_KP9 = 73
const KEY_KPMINUS = 74
const KEY_KP4 = 75
const KEY_KP5 = 76
const KEY_KP6 = 77
const KEY_KPPLUS = 78
const KEY_KP1 = 79
const KEY_KP2 = 80
const KEY_KP3 = 81
const KEY_KP0 = 82
const KEY_KPDOT = 83

const KEY_ZENKAKUHANKAKU = 85
const KEY_102ND = 86
const KEY_F11 = 87
const KEY_F12 = 88
const KEY_RO = 89
const KEY_KATAKANA = 90
const KEY_HIRAGANA = 91
const KEY_HENKAN = 92
const KEY_KATAKANAHIRAGANA = 93
const KEY_MUHENKAN = 94
const KEY_KPJPCOMMA = 95
const KEY_KPENTER = 96
const KEY_RIGHTCTRL = 97
const KEY_KPSLASH = 98
const KEY_SYSRQ = 99
const KEY_RIGHTALT = 100
const KEY_LINEFEED = 101
const KEY_HOME = 102
const KEY_UP = 103
const KEY_PAGEUP = 104
const KEY_LEFT = 105
const KEY_RIGHT = 106
const KEY_END = 107
const KEY_DOWN = 108
const KEY_PAGEDOWN = 109
const KEY_INSERT = 110
const KEY_DELETE = 111
const KEY_MACRO = 112
const KEY_MUTE = 113
const KEY_VOLUMEDOWN = 114
const KEY_VOLUMEUP = 115
const KEY_POWER = 116 /* SC System Power Down */
const KEY_KPEQUAL = 117
const KEY_KPPLUSMINUS = 118
const KEY_PAUSE = 119
const KEY_SCALE = 120 /* AL Compiz Scale (Expose) */

const KEY_KPCOMMA = 121
const KEY_HANGEUL = 122
const KEY_HANGUEL = KEY_HANGEUL
const KEY_HANJA = 123
const KEY_YEN = 124
const KEY_LEFTMETA = 125
const KEY_RIGHTMETA = 126
const KEY_COMPOSE = 127

const KEY_STOP = 128 /* AC Stop */
const KEY_AGAIN = 129
const KEY_PROPS = 130 /* AC Properties */
const KEY_UNDO = 131  /* AC Undo */
const KEY_FRONT = 132
const KEY_COPY = 133  /* AC Copy */
const KEY_OPEN = 134  /* AC Open */
const KEY_PASTE = 135 /* AC Paste */
const KEY_FIND = 136  /* AC Search */
const KEY_CUT = 137   /* AC Cut */
const KEY_HELP = 138  /* AL Integrated Help Center */
const KEY_MENU = 139  /* Menu (show menu) */
const KEY_CALC = 140  /* AL Calculator */
const KEY_SETUP = 141
const KEY_SLEEP = 142  /* SC System Sleep */
const KEY_WAKEUP = 143 /* System Wake Up */
const KEY_FILE = 144   /* AL Local Machine Browser */
const KEY_SENDFILE = 145
const KEY_DELETEFILE = 146
const KEY_XFER = 147
const KEY_PROG1 = 148
const KEY_PROG2 = 149
const KEY_WWW = 150 /* AL Internet Browser */
const KEY_MSDOS = 151
const KEY_COFFEE = 152 /* AL Terminal Lock/Screensaver */
const KEY_SCREENLOCK = KEY_COFFEE
const KEY_ROTATE_DISPLAY = 153 /* Display orientation for e.g. tablets */
const KEY_DIRECTION = KEY_ROTATE_DISPLAY
const KEY_CYCLEWINDOWS = 154
const KEY_MAIL = 155
const KEY_BOOKMARKS = 156 /* AC Bookmarks */
const KEY_COMPUTER = 157
const KEY_BACK = 158    /* AC Back */
const KEY_FORWARD = 159 /* AC Forward */
const KEY_CLOSECD = 160
const KEY_EJECTCD = 161
const KEY_EJECTCLOSECD = 162
const KEY_NEXTSONG = 163
const KEY_PLAYPAUSE = 164
const KEY_PREVIOUSSONG = 165
const KEY_STOPCD = 166
const KEY_RECORD = 167
const KEY_REWIND = 168
const KEY_PHONE = 169 /* Media Select Telephone */
const KEY_ISO = 170
const KEY_CONFIG = 171   /* AL Consumer Control Configuration */
const KEY_HOMEPAGE = 172 /* AC Home */
const KEY_REFRESH = 173  /* AC Refresh */
const KEY_EXIT = 174     /* AC Exit */
const KEY_MOVE = 175
const KEY_EDIT = 176
const KEY_SCROLLUP = 177
const KEY_SCROLLDOWN = 178
const KEY_KPLEFTPAREN = 179
const KEY_KPRIGHTPAREN = 180
const KEY_NEW = 181  /* AC New */
const KEY_REDO = 182 /* AC Redo/Repeat */

const KEY_F13 = 183
const KEY_F14 = 184
const KEY_F15 = 185
const KEY_F16 = 186
const KEY_F17 = 187
const KEY_F18 = 188
const KEY_F19 = 189
const KEY_F20 = 190
const KEY_F21 = 191
const KEY_F22 = 192
const KEY_F23 = 193
const KEY_F24 = 194

const KEY_PLAYCD = 200
const KEY_PAUSECD = 201
const KEY_PROG3 = 202
const KEY_PROG4 = 203
const KEY_DASHBOARD = 204 /* AL Dashboard */
const KEY_SUSPEND = 205
const KEY_CLOSE = 206 /* AC Close */
const KEY_PLAY = 207
const KEY_FASTFORWARD = 208
const KEY_BASSBOOST = 209
const KEY_PRINT = 210 /* AC Print */
const KEY_HP = 211
const KEY_CAMERA = 212
const KEY_SOUND = 213
const KEY_QUESTION = 214
const KEY_EMAIL = 215
const KEY_CHAT = 216
const KEY_SEARCH = 217
const KEY_CONNECT = 218
const KEY_FINANCE = 219 /* AL Checkbook/Finance */
const KEY_SPORT = 220
const KEY_SHOP = 221
const KEY_ALTERASE = 222
const KEY_CANCEL = 223 /* AC Cancel */
const KEY_BRIGHTNESSDOWN = 224
const KEY_BRIGHTNESSUP = 225
const KEY_MEDIA = 226

const KEY_SWITCHVIDEOMODE = 227 /* Cycle between available video
   outputs (Monitor/LCD/TV-out/etc) */
const KEY_KBDILLUMTOGGLE = 228
const KEY_KBDILLUMDOWN = 229
const KEY_KBDILLUMUP = 230

const KEY_SEND = 231        /* AC Send */
const KEY_REPLY = 232       /* AC Reply */
const KEY_FORWARDMAIL = 233 /* AC Forward Msg */
const KEY_SAVE = 234        /* AC Save */
const KEY_DOCUMENTS = 235

const KEY_BATTERY = 236

const KEY_BLUETOOTH = 237
const KEY_WLAN = 238
const KEY_UWB = 239

const KEY_UNKNOWN = 240

const KEY_VIDEO_NEXT = 241       /* drive next video source */
const KEY_VIDEO_PREV = 242       /* drive previous video source */
const KEY_BRIGHTNESS_CYCLE = 243 /* brightness up, after max is min */
const KEY_BRIGHTNESS_AUTO = 244  /* Set Auto Brightness: manual
                      brightness control is off,
					  rely on ambient */
const KEY_BRIGHTNESS_ZERO = KEY_BRIGHTNESS_AUTO
const KEY_DISPLAY_OFF = 245 /* display device to off state */

const KEY_WWAN = 246 /* Wireless WAN (LTE, UMTS, GSM, etc.) */
const KEY_WIMAX = KEY_WWAN
const KEY_RFKILL = 247 /* Key that controls all radios */

const KEY_MICMUTE = 248 /* Mute / unmute the microphone */

/* Code 255 is reserved for special needs of AT keyboard driver */

const BTN_MISC = 0x100
const BTN_0 = 0x100
const BTN_1 = 0x101
const BTN_2 = 0x102
const BTN_3 = 0x103
const BTN_4 = 0x104
const BTN_5 = 0x105
const BTN_6 = 0x106
const BTN_7 = 0x107
const BTN_8 = 0x108
const BTN_9 = 0x109

const BTN_MOUSE = 0x110
const BTN_LEFT = 0x110
const BTN_RIGHT = 0x111
const BTN_MIDDLE = 0x112
const BTN_SIDE = 0x113
const BTN_EXTRA = 0x114
const BTN_FORWARD = 0x115
const BTN_BACK = 0x116
const BTN_TASK = 0x117

const BTN_JOYSTICK = 0x120
const BTN_TRIGGER = 0x120
const BTN_THUMB = 0x121
const BTN_THUMB2 = 0x122
const BTN_TOP = 0x123
const BTN_TOP2 = 0x124
const BTN_PINKIE = 0x125
const BTN_BASE = 0x126
const BTN_BASE2 = 0x127
const BTN_BASE3 = 0x128
const BTN_BASE4 = 0x129
const BTN_BASE5 = 0x12a
const BTN_BASE6 = 0x12b
const BTN_DEAD = 0x12f

const BTN_GAMEPAD = 0x130
const BTN_SOUTH = 0x130
const BTN_A = BTN_SOUTH
const BTN_EAST = 0x131
const BTN_B = BTN_EAST
const BTN_C = 0x132
const BTN_NORTH = 0x133
const BTN_X = BTN_NORTH
const BTN_WEST = 0x134
const BTN_Y = BTN_WEST
const BTN_Z = 0x135
const BTN_TL = 0x136
const BTN_TR = 0x137
const BTN_TL2 = 0x138
const BTN_TR2 = 0x139
const BTN_SELECT = 0x13a
const BTN_START = 0x13b
const BTN_MODE = 0x13c
const BTN_THUMBL = 0x13d
const BTN_THUMBR = 0x13e

const BTN_DIGI = 0x140
const BTN_TOOL_PEN = 0x140
const BTN_TOOL_RUBBER = 0x141
const BTN_TOOL_BRUSH = 0x142
const BTN_TOOL_PENCIL = 0x143
const BTN_TOOL_AIRBRUSH = 0x144
const BTN_TOOL_FINGER = 0x145
const BTN_TOOL_MOUSE = 0x146
const BTN_TOOL_LENS = 0x147
const BTN_TOOL_QUINTTAP = 0x148 /* Five fingers on trackpad */
const BTN_TOUCH = 0x14a
const BTN_STYLUS = 0x14b
const BTN_STYLUS2 = 0x14c
const BTN_TOOL_DOUBLETAP = 0x14d
const BTN_TOOL_TRIPLETAP = 0x14e
const BTN_TOOL_QUADTAP = 0x14f /* Four fingers on trackpad */

const BTN_WHEEL = 0x150
const BTN_GEAR_DOWN = 0x150
const BTN_GEAR_UP = 0x151

const KEY_OK = 0x160
const KEY_SELECT = 0x161
const KEY_GOTO = 0x162
const KEY_CLEAR = 0x163
const KEY_POWER2 = 0x164
const KEY_OPTION = 0x165
const KEY_INFO = 0x166 /* AL OEM Features/Tips/Tutorial */
const KEY_TIME = 0x167
const KEY_VENDOR = 0x168
const KEY_ARCHIVE = 0x169
const KEY_PROGRAM = 0x16a /* Media Select Program Guide */
const KEY_CHANNEL = 0x16b
const KEY_FAVORITES = 0x16c
const KEY_EPG = 0x16d
const KEY_PVR = 0x16e /* Media Select Home */
const KEY_MHP = 0x16f
const KEY_LANGUAGE = 0x170
const KEY_TITLE = 0x171
const KEY_SUBTITLE = 0x172
const KEY_ANGLE = 0x173
const KEY_ZOOM = 0x174
const KEY_MODE = 0x175
const KEY_KEYBOARD = 0x176
const KEY_SCREEN = 0x177
const KEY_PC = 0x178   /* Media Select Computer */
const KEY_TV = 0x179   /* Media Select TV */
const KEY_TV2 = 0x17a  /* Media Select Cable */
const KEY_VCR = 0x17b  /* Media Select VCR */
const KEY_VCR2 = 0x17c /* VCR Plus */
const KEY_SAT = 0x17d  /* Media Select Satellite */
const KEY_SAT2 = 0x17e
const KEY_CD = 0x17f   /* Media Select CD */
const KEY_TAPE = 0x180 /* Media Select Tape */
const KEY_RADIO = 0x181
const KEY_TUNER = 0x182 /* Media Select Tuner */
const KEY_PLAYER = 0x183
const KEY_TEXT = 0x184
const KEY_DVD = 0x185 /* Media Select DVD */
const KEY_AUX = 0x186
const KEY_MP3 = 0x187
const KEY_AUDIO = 0x188 /* AL Audio Browser */
const KEY_VIDEO = 0x189 /* AL Movie Browser */
const KEY_DIRECTORY = 0x18a
const KEY_LIST = 0x18b
const KEY_MEMO = 0x18c /* Media Select Messages */
const KEY_CALENDAR = 0x18d
const KEY_RED = 0x18e
const KEY_GREEN = 0x18f
const KEY_YELLOW = 0x190
const KEY_BLUE = 0x191
const KEY_CHANNELUP = 0x192   /* Channel Increment */
const KEY_CHANNELDOWN = 0x193 /* Channel Decrement */
const KEY_FIRST = 0x194
const KEY_LAST = 0x195 /* Recall Last */
const KEY_AB = 0x196
const KEY_NEXT = 0x197
const KEY_RESTART = 0x198
const KEY_SLOW = 0x199
const KEY_SHUFFLE = 0x19a
const KEY_BREAK = 0x19b
const KEY_PREVIOUS = 0x19c
const KEY_DIGITS = 0x19d
const KEY_TEEN = 0x19e
const KEY_TWEN = 0x19f
const KEY_VIDEOPHONE = 0x1a0     /* Media Select Video Phone */
const KEY_GAMES = 0x1a1          /* Media Select Games */
const KEY_ZOOMIN = 0x1a2         /* AC Zoom In */
const KEY_ZOOMOUT = 0x1a3        /* AC Zoom Out */
const KEY_ZOOMRESET = 0x1a4      /* AC Zoom */
const KEY_WORDPROCESSOR = 0x1a5  /* AL Word Processor */
const KEY_EDITOR = 0x1a6         /* AL Text Editor */
const KEY_SPREADSHEET = 0x1a7    /* AL Spreadsheet */
const KEY_GRAPHICSEDITOR = 0x1a8 /* AL Graphics Editor */
const KEY_PRESENTATION = 0x1a9   /* AL Presentation App */
const KEY_DATABASE = 0x1aa       /* AL Database App */
const KEY_NEWS = 0x1ab           /* AL Newsreader */
const KEY_VOICEMAIL = 0x1ac      /* AL Voicemail */
const KEY_ADDRESSBOOK = 0x1ad    /* AL Contacts/Address Book */
const KEY_MESSENGER = 0x1ae      /* AL Instant Messaging */
const KEY_DISPLAYTOGGLE = 0x1af  /* Turn display (LCD) on and off */
const KEY_BRIGHTNESS_TOGGLE = KEY_DISPLAYTOGGLE
const KEY_SPELLCHECK = 0x1b0 /* AL Spell Check */
const KEY_LOGOFF = 0x1b1     /* AL Logoff */

const KEY_DOLLAR = 0x1b2
const KEY_EURO = 0x1b3

const KEY_FRAMEBACK = 0x1b4 /* Consumer - transport controls */
const KEY_FRAMEFORWARD = 0x1b5
const KEY_CONTEXT_MENU = 0x1b6   /* GenDesc - system context menu */
const KEY_MEDIA_REPEAT = 0x1b7   /* Consumer - transport control */
const KEY_10CHANNELSUP = 0x1b8   /* 10 channels up (10+) */
const KEY_10CHANNELSDOWN = 0x1b9 /* 10 channels down (10-) */
const KEY_IMAGES = 0x1ba         /* AL Image Browser */

const KEY_DEL_EOL = 0x1c0
const KEY_DEL_EOS = 0x1c1
const KEY_INS_LINE = 0x1c2
const KEY_DEL_LINE = 0x1c3

const KEY_FN = 0x1d0
const KEY_FN_ESC = 0x1d1
const KEY_FN_F1 = 0x1d2
const KEY_FN_F2 = 0x1d3
const KEY_FN_F3 = 0x1d4
const KEY_FN_F4 = 0x1d5
const KEY_FN_F5 = 0x1d6
const KEY_FN_F6 = 0x1d7
const KEY_FN_F7 = 0x1d8
const KEY_FN_F8 = 0x1d9
const KEY_FN_F9 = 0x1da
const KEY_FN_F10 = 0x1db
const KEY_FN_F11 = 0x1dc
const KEY_FN_F12 = 0x1dd
const KEY_FN_1 = 0x1de
const KEY_FN_2 = 0x1df
const KEY_FN_D = 0x1e0
const KEY_FN_E = 0x1e1
const KEY_FN_F = 0x1e2
const KEY_FN_S = 0x1e3
const KEY_FN_B = 0x1e4

const KEY_BRL_DOT1 = 0x1f1
const KEY_BRL_DOT2 = 0x1f2
const KEY_BRL_DOT3 = 0x1f3
const KEY_BRL_DOT4 = 0x1f4
const KEY_BRL_DOT5 = 0x1f5
const KEY_BRL_DOT6 = 0x1f6
const KEY_BRL_DOT7 = 0x1f7
const KEY_BRL_DOT8 = 0x1f8
const KEY_BRL_DOT9 = 0x1f9
const KEY_BRL_DOT10 = 0x1fa

const KEY_NUMERIC_0 = 0x200 /* used by phones, remote controls, */
const KEY_NUMERIC_1 = 0x201 /* and other keypads */
const KEY_NUMERIC_2 = 0x202
const KEY_NUMERIC_3 = 0x203
const KEY_NUMERIC_4 = 0x204
const KEY_NUMERIC_5 = 0x205
const KEY_NUMERIC_6 = 0x206
const KEY_NUMERIC_7 = 0x207
const KEY_NUMERIC_8 = 0x208
const KEY_NUMERIC_9 = 0x209
const KEY_NUMERIC_STAR = 0x20a
const KEY_NUMERIC_POUND = 0x20b
const KEY_NUMERIC_A = 0x20c /* Phone key A - HUT Telephony 0xb9 */
const KEY_NUMERIC_B = 0x20d
const KEY_NUMERIC_C = 0x20e
const KEY_NUMERIC_D = 0x20f

const KEY_CAMERA_FOCUS = 0x210
const KEY_WPS_BUTTON = 0x211 /* WiFi Protected Setup key */

const KEY_TOUCHPAD_TOGGLE = 0x212 /* Request switch touchpad on or off */
const KEY_TOUCHPAD_ON = 0x213
const KEY_TOUCHPAD_OFF = 0x214

const KEY_CAMERA_ZOOMIN = 0x215
const KEY_CAMERA_ZOOMOUT = 0x216
const KEY_CAMERA_UP = 0x217
const KEY_CAMERA_DOWN = 0x218
const KEY_CAMERA_LEFT = 0x219
const KEY_CAMERA_RIGHT = 0x21a

const KEY_ATTENDANT_ON = 0x21b
const KEY_ATTENDANT_OFF = 0x21c
const KEY_ATTENDANT_TOGGLE = 0x21d /* Attendant call on or off */
const KEY_LIGHTS_TOGGLE = 0x21e    /* Reading light on or off */

const BTN_DPAD_UP = 0x220
const BTN_DPAD_DOWN = 0x221
const BTN_DPAD_LEFT = 0x222
const BTN_DPAD_RIGHT = 0x223

const KEY_ALS_TOGGLE = 0x230 /* Ambient light sensor */

const KEY_BUTTONCONFIG = 0x240 /* AL Button Configuration */
const KEY_TASKMANAGER = 0x241  /* AL Task/Project Manager */
const KEY_JOURNAL = 0x242      /* AL Log/Journal/Timecard */
const KEY_CONTROLPANEL = 0x243 /* AL Control Panel */
const KEY_APPSELECT = 0x244    /* AL Select Task/Application */
const KEY_SCREENSAVER = 0x245  /* AL Screen Saver */
const KEY_VOICECOMMAND = 0x246 /* Listening Voice Command */
const KEY_ASSISTANT = 0x247    /* AL Context-aware desktop assistant */

const KEY_BRIGHTNESS_MIN = 0x250 /* Set Brightness to Minimum */
const KEY_BRIGHTNESS_MAX = 0x251 /* Set Brightness to Maximum */

const KEY_KBDINPUTASSIST_PREV = 0x260
const KEY_KBDINPUTASSIST_NEXT = 0x261
const KEY_KBDINPUTASSIST_PREVGROUP = 0x262
const KEY_KBDINPUTASSIST_NEXTGROUP = 0x263
const KEY_KBDINPUTASSIST_ACCEPT = 0x264
const KEY_KBDINPUTASSIST_CANCEL = 0x265

/* Diagonal movement keys */
const KEY_RIGHT_UP = 0x266
const KEY_RIGHT_DOWN = 0x267
const KEY_LEFT_UP = 0x268
const KEY_LEFT_DOWN = 0x269

const KEY_ROOT_MENU = 0x26a /* Show Device's Root Menu */
/* Show Top Menu of the Media (e.g. DVD) */
const KEY_MEDIA_TOP_MENU = 0x26b
const KEY_NUMERIC_11 = 0x26c
const KEY_NUMERIC_12 = 0x26d

/*
 * Toggle Audio Description: refers to an audio service that helps blind and
 * visually impaired consumers understand the action in a program. Note: in
 * some countries this is referred to as "Video Description".
 */
const KEY_AUDIO_DESC = 0x26e
const KEY_3D_MODE = 0x26f
const KEY_NEXT_FAVORITE = 0x270
const KEY_STOP_RECORD = 0x271
const KEY_PAUSE_RECORD = 0x272
const KEY_VOD = 0x273 /* Video on Demand */
const KEY_UNMUTE = 0x274
const KEY_FASTREVERSE = 0x275
const KEY_SLOWREVERSE = 0x276

/*
 * Control a data application associated with the currently viewed channel,
 * e.g. teletext or data broadcast application (MHEG, MHP, HbbTV, etc.)
 */
const KEY_DATA = 0x277
const KEY_ONSCREEN_KEYBOARD = 0x278

const BTN_TRIGGER_HAPPY = 0x2c0
const BTN_TRIGGER_HAPPY1 = 0x2c0
const BTN_TRIGGER_HAPPY2 = 0x2c1
const BTN_TRIGGER_HAPPY3 = 0x2c2
const BTN_TRIGGER_HAPPY4 = 0x2c3
const BTN_TRIGGER_HAPPY5 = 0x2c4
const BTN_TRIGGER_HAPPY6 = 0x2c5
const BTN_TRIGGER_HAPPY7 = 0x2c6
const BTN_TRIGGER_HAPPY8 = 0x2c7
const BTN_TRIGGER_HAPPY9 = 0x2c8
const BTN_TRIGGER_HAPPY10 = 0x2c9
const BTN_TRIGGER_HAPPY11 = 0x2ca
const BTN_TRIGGER_HAPPY12 = 0x2cb
const BTN_TRIGGER_HAPPY13 = 0x2cc
const BTN_TRIGGER_HAPPY14 = 0x2cd
const BTN_TRIGGER_HAPPY15 = 0x2ce
const BTN_TRIGGER_HAPPY16 = 0x2cf
const BTN_TRIGGER_HAPPY17 = 0x2d0
const BTN_TRIGGER_HAPPY18 = 0x2d1
const BTN_TRIGGER_HAPPY19 = 0x2d2
const BTN_TRIGGER_HAPPY20 = 0x2d3
const BTN_TRIGGER_HAPPY21 = 0x2d4
const BTN_TRIGGER_HAPPY22 = 0x2d5
const BTN_TRIGGER_HAPPY23 = 0x2d6
const BTN_TRIGGER_HAPPY24 = 0x2d7
const BTN_TRIGGER_HAPPY25 = 0x2d8
const BTN_TRIGGER_HAPPY26 = 0x2d9
const BTN_TRIGGER_HAPPY27 = 0x2da
const BTN_TRIGGER_HAPPY28 = 0x2db
const BTN_TRIGGER_HAPPY29 = 0x2dc
const BTN_TRIGGER_HAPPY30 = 0x2dd
const BTN_TRIGGER_HAPPY31 = 0x2de
const BTN_TRIGGER_HAPPY32 = 0x2df
const BTN_TRIGGER_HAPPY33 = 0x2e0
const BTN_TRIGGER_HAPPY34 = 0x2e1
const BTN_TRIGGER_HAPPY35 = 0x2e2
const BTN_TRIGGER_HAPPY36 = 0x2e3
const BTN_TRIGGER_HAPPY37 = 0x2e4
const BTN_TRIGGER_HAPPY38 = 0x2e5
const BTN_TRIGGER_HAPPY39 = 0x2e6
const BTN_TRIGGER_HAPPY40 = 0x2e7

/* We avoid low common keys in module aliases so they don't get huge. */
const KEY_MIN_INTERESTING = KEY_MUTE
const KEY_MAX = 0x2ff
const KEY_CNT = (KEY_MAX + 1)

/*
 * Relative axes
 */

const REL_X = 0x00
const REL_Y = 0x01
const REL_Z = 0x02
const REL_RX = 0x03
const REL_RY = 0x04
const REL_RZ = 0x05
const REL_HWHEEL = 0x06
const REL_DIAL = 0x07
const REL_WHEEL = 0x08
const REL_MISC = 0x09
const REL_MAX = 0x0f
const REL_CNT = (REL_MAX + 1)

/*
 * Absolute axes
 */

const ABS_X = 0x00
const ABS_Y = 0x01
const ABS_Z = 0x02
const ABS_RX = 0x03
const ABS_RY = 0x04
const ABS_RZ = 0x05
const ABS_THROTTLE = 0x06
const ABS_RUDDER = 0x07
const ABS_WHEEL = 0x08
const ABS_GAS = 0x09
const ABS_BRAKE = 0x0a
const ABS_HAT0X = 0x10
const ABS_HAT0Y = 0x11
const ABS_HAT1X = 0x12
const ABS_HAT1Y = 0x13
const ABS_HAT2X = 0x14
const ABS_HAT2Y = 0x15
const ABS_HAT3X = 0x16
const ABS_HAT3Y = 0x17
const ABS_PRESSURE = 0x18
const ABS_DISTANCE = 0x19
const ABS_TILT_X = 0x1a
const ABS_TILT_Y = 0x1b
const ABS_TOOL_WIDTH = 0x1c

const ABS_VOLUME = 0x20

const ABS_MISC = 0x28

const ABS_MT_SLOT = 0x2f        /* MT slot being modified */
const ABS_MT_TOUCH_MAJOR = 0x30 /* Major axis of touching ellipse */
const ABS_MT_TOUCH_MINOR = 0x31 /* Minor axis (omit if circular) */
const ABS_MT_WIDTH_MAJOR = 0x32 /* Major axis of approaching ellipse */
const ABS_MT_WIDTH_MINOR = 0x33 /* Minor axis (omit if circular) */
const ABS_MT_ORIENTATION = 0x34 /* Ellipse orientation */
const ABS_MT_POSITION_X = 0x35  /* Center X touch position */
const ABS_MT_POSITION_Y = 0x36  /* Center Y touch position */
const ABS_MT_TOOL_TYPE = 0x37   /* Type of touching device */
const ABS_MT_BLOB_ID = 0x38     /* Group a set of packets as a blob */
const ABS_MT_TRACKING_ID = 0x39 /* Unique ID of initiated contact */
const ABS_MT_PRESSURE = 0x3a    /* Pressure on contact area */
const ABS_MT_DISTANCE = 0x3b    /* Contact hover distance */
const ABS_MT_TOOL_X = 0x3c      /* Center X tool position */
const ABS_MT_TOOL_Y = 0x3d      /* Center Y tool position */

const ABS_MAX = 0x3f
const ABS_CNT = (ABS_MAX + 1)

/*
 * Switch events
 */

const SW_LID = 0x00              /* set = lid shut */
const SW_TABLET_MODE = 0x01      /* set = tablet mode */
const SW_HEADPHONE_INSERT = 0x02 /* set = inserted */
const SW_RFKILL_ALL = 0x03       /* rfkill master switch, type "any"
   set = radio enabled */
const SW_RADIO = SW_RFKILL_ALL       /* deprecated */
const SW_MICROPHONE_INSERT = 0x04    /* set = inserted */
const SW_DOCK = 0x05                 /* set = plugged into dock */
const SW_LINEOUT_INSERT = 0x06       /* set = inserted */
const SW_JACK_PHYSICAL_INSERT = 0x07 /* set = mechanical switch set */
const SW_VIDEOOUT_INSERT = 0x08      /* set = inserted */
const SW_CAMERA_LENS_COVER = 0x09    /* set = lens covered */
const SW_KEYPAD_SLIDE = 0x0a         /* set = keypad slide out */
const SW_FRONT_PROXIMITY = 0x0b      /* set = front proximity sensor active */
const SW_ROTATE_LOCK = 0x0c          /* set = rotate locked/disabled */
const SW_LINEIN_INSERT = 0x0d        /* set = inserted */
const SW_MUTE_DEVICE = 0x0e          /* set = device disabled */
const SW_PEN_INSERTED = 0x0f         /* set = pen inserted */
const SW_MAX = 0x0f
const SW_CNT = (SW_MAX + 1)

/*
 * Misc events
 */

const MSC_SERIAL = 0x00
const MSC_PULSELED = 0x01
const MSC_GESTURE = 0x02
const MSC_RAW = 0x03
const MSC_SCAN = 0x04
const MSC_TIMESTAMP = 0x05
const MSC_MAX = 0x07
const MSC_CNT = (MSC_MAX + 1)

/*
 * LEDs
 */

const LED_NUML = 0x00
const LED_CAPSL = 0x01
const LED_SCROLLL = 0x02
const LED_COMPOSE = 0x03
const LED_KANA = 0x04
const LED_SLEEP = 0x05
const LED_SUSPEND = 0x06
const LED_MUTE = 0x07
const LED_MISC = 0x08
const LED_MAIL = 0x09
const LED_CHARGING = 0x0a
const LED_MAX = 0x0f
const LED_CNT = (LED_MAX + 1)

/*
 * Autorepeat values
 */

const REP_DELAY = 0x00
const REP_PERIOD = 0x01
const REP_MAX = 0x01
const REP_CNT = (REP_MAX + 1)

/*
 * Sounds
 */

const SND_CLICK = 0x00
const SND_BELL = 0x01
const SND_TONE = 0x02
const SND_MAX = 0x07
const SND_CNT = (SND_MAX + 1)
