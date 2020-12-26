package main

var themes = map[string]map[string]string{
	"default": {
		"fgcol": "#8C8C8C",
		"bgcol": "#282828",

		"hicol2": "#805b13",
		"hicol3": "#b4801b",
		"hicol":  "#ffffff",
		"errcol": "#a10705",
	},

	//Generated from terminal themes, probably suboptimal
	"3024-day": map[string]string{
		"hicol2": "#db2d20",
		"hicol3": "#db2d20",
		"hicol":  "#a5a2a2",
		"bgcol":  "#f7f7f7",
		"fgcol":  "#4a4543",
		"errcol": "#ff0000",
	},

	"3024-night": map[string]string{
		"hicol2": "#db2d20",
		"hicol3": "#db2d20",
		"hicol":  "#a5a2a2",
		"bgcol":  "#090300",
		"fgcol":  "#a5a2a2",
		"errcol": "#ff0000",
	},

	"aci": map[string]string{
		"hicol2": "#ff0883",
		"hicol3": "#ff0883",
		"hicol":  "#b6b6b6",
		"bgcol":  "#0d1926",
		"fgcol":  "#b4e1fd",
		"errcol": "#ff0000",
	},

	"aco": map[string]string{
		"hicol2": "#ff0883",
		"hicol3": "#ff0883",
		"hicol":  "#bebebe",
		"bgcol":  "#1f1305",
		"fgcol":  "#b4e1fd",
		"errcol": "#ff0000",
	},

	"adventuretime": map[string]string{
		"hicol2": "#bd0013",
		"hicol3": "#bd0013",
		"hicol":  "#f8dcc0",
		"bgcol":  "#1f1d45",
		"fgcol":  "#f8dcc0",
		"errcol": "#ff0000",
	},

	"afterglow": map[string]string{
		"hicol2": "#a53c23",
		"hicol3": "#a53c23",
		"hicol":  "#d0d0d0",
		"bgcol":  "#222222",
		"fgcol":  "#d0d0d0",
		"errcol": "#ff0000",
	},

	"alien-blood": map[string]string{
		"hicol2": "#7f2b27",
		"hicol3": "#7f2b27",
		"hicol":  "#647d75",
		"bgcol":  "#0f1610",
		"fgcol":  "#637d75",
		"errcol": "#ff0000",
	},

	"argonaut": map[string]string{
		"hicol2": "#ff000f",
		"hicol3": "#ff000f",
		"hicol":  "#ffffff",
		"bgcol":  "#0e1019",
		"fgcol":  "#fffaf4",
		"errcol": "#ff0000",
	},

	"arthur": map[string]string{
		"hicol2": "#cd5c5c",
		"hicol3": "#cd5c5c",
		"hicol":  "#bbaa99",
		"bgcol":  "#1c1c1c",
		"fgcol":  "#ddeedd",
		"errcol": "#ff0000",
	},

	"atom": map[string]string{
		"hicol2": "#fd5ff1",
		"hicol3": "#fd5ff1",
		"hicol":  "#e0e0e0",
		"bgcol":  "#161719",
		"fgcol":  "#c5c8c6",
		"errcol": "#ff0000",
	},

	"azu": map[string]string{
		"hicol2": "#ac6d74",
		"hicol3": "#ac6d74",
		"hicol":  "#e6e6e6",
		"bgcol":  "#09111a",
		"fgcol":  "#d9e6f2",
		"errcol": "#ff0000",
	},

	"_base": map[string]string{
		"hicol2": "##C54133",
		"hicol3": "##C54133",
		"hicol":  "##C9CCCD",
		"bgcol":  "#260346",
		"fgcol":  "#DADADA",
		"errcol": "#ff0000",
	},

	"belafonte-day": map[string]string{
		"hicol2": "#be100e",
		"hicol3": "#be100e",
		"hicol":  "#968c83",
		"bgcol":  "#d5ccba",
		"fgcol":  "#45373c",
		"errcol": "#ff0000",
	},

	"belafonte-night": map[string]string{
		"hicol2": "#be100e",
		"hicol3": "#be100e",
		"hicol":  "#968c83",
		"bgcol":  "#20111b",
		"fgcol":  "#968c83",
		"errcol": "#ff0000",
	},

	"bim": map[string]string{
		"hicol2": "#f557a0",
		"hicol3": "#f557a0",
		"hicol":  "#918988",
		"bgcol":  "#012849",
		"fgcol":  "#a9bed8",
		"errcol": "#ff0000",
	},

	"birds-of-paradise": map[string]string{
		"hicol2": "#be2d26",
		"hicol3": "#be2d26",
		"hicol":  "#e0dbb7",
		"bgcol":  "#2a1f1d",
		"fgcol":  "#e0dbb7",
		"errcol": "#ff0000",
	},

	"blazer": map[string]string{
		"hicol2": "#b87a7a",
		"hicol3": "#b87a7a",
		"hicol":  "#d9d9d9",
		"bgcol":  "#0d1926",
		"fgcol":  "#d9e6f2",
		"errcol": "#ff0000",
	},

	"borland": map[string]string{
		"hicol2": "#ff6c60",
		"hicol3": "#ff6c60",
		"hicol":  "#eeeeee",
		"bgcol":  "#0000a4",
		"fgcol":  "#ffff4e",
		"errcol": "#ff0000",
	},

	"broadcast": map[string]string{
		"hicol2": "#da4939",
		"hicol3": "#da4939",
		"hicol":  "#ffffff",
		"bgcol":  "#2b2b2b",
		"fgcol":  "#e6e1dc",
		"errcol": "#ff0000",
	},

	"brogrammer": map[string]string{
		"hicol2": "#f81118",
		"hicol3": "#f81118",
		"hicol":  "#d6dbe5",
		"bgcol":  "#131313",
		"fgcol":  "#d6dbe5",
		"errcol": "#ff0000",
	},

	"c64": map[string]string{
		"hicol2": "#883932",
		"hicol3": "#883932",
		"hicol":  "#ffffff",
		"bgcol":  "#40318d",
		"fgcol":  "#7869c4",
		"errcol": "#ff0000",
	},

	"cai": map[string]string{
		"hicol2": "#ca274d",
		"hicol3": "#ca274d",
		"hicol":  "#808080",
		"bgcol":  "#09111a",
		"fgcol":  "#d9e6f2",
		"errcol": "#ff0000",
	},

	"chalkboard": map[string]string{
		"hicol2": "#c37372",
		"hicol3": "#c37372",
		"hicol":  "#d9d9d9",
		"bgcol":  "#29262f",
		"fgcol":  "#d9e6f2",
		"errcol": "#ff0000",
	},

	"chalk": map[string]string{
		"hicol2": "#F58E8E",
		"hicol3": "#F58E8E",
		"hicol":  "#D4D4D4",
		"bgcol":  "#2D2D2D",
		"fgcol":  "#D4D4D4",
		"errcol": "#ff0000",
	},

	"ciapre": map[string]string{
		"hicol2": "#810009",
		"hicol3": "#810009",
		"hicol":  "#aea47f",
		"bgcol":  "#191c27",
		"fgcol":  "#aea47a",
		"errcol": "#ff0000",
	},

	"clone-of-ubuntu": map[string]string{
		"hicol2": "#CC0000",
		"hicol3": "#CC0000",
		"hicol":  "#D3D7CF",
		"bgcol":  "#300a24",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"clrs": map[string]string{
		"hicol2": "#f8282a",
		"hicol3": "#f8282a",
		"hicol":  "#b3b3b3",
		"bgcol":  "#ffffff",
		"fgcol":  "#262626",
		"errcol": "#ff0000",
	},

	"cobalt2": map[string]string{
		"hicol2": "#ff0000",
		"hicol3": "#ff0000",
		"hicol":  "#bbbbbb",
		"bgcol":  "#132738",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"cobalt-neon": map[string]string{
		"hicol2": "#ff2320",
		"hicol3": "#ff2320",
		"hicol":  "#ba46b2",
		"bgcol":  "#142838",
		"fgcol":  "#8ff586",
		"errcol": "#ff0000",
	},

	"crayon-pony-fish": map[string]string{
		"hicol2": "#91002b",
		"hicol3": "#91002b",
		"hicol":  "#68525a",
		"bgcol":  "#150707",
		"fgcol":  "#68525a",
		"errcol": "#ff0000",
	},

	"dark-pastel": map[string]string{
		"hicol2": "#ff5555",
		"hicol3": "#ff5555",
		"hicol":  "#bbbbbb",
		"bgcol":  "#000000",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"darkside": map[string]string{
		"hicol2": "#e8341c",
		"hicol3": "#e8341c",
		"hicol":  "#bababa",
		"bgcol":  "#222324",
		"fgcol":  "#bababa",
		"errcol": "#ff0000",
	},

	"desert": map[string]string{
		"hicol2": "#ff2b2b",
		"hicol3": "#ff2b2b",
		"hicol":  "#f5deb3",
		"bgcol":  "#333333",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"dimmed-monokai": map[string]string{
		"hicol2": "#be3f48",
		"hicol3": "#be3f48",
		"hicol":  "#b9bcba",
		"bgcol":  "#1f1f1f",
		"fgcol":  "#b9bcba",
		"errcol": "#ff0000",
	},

	"dracula": map[string]string{
		"hicol2": "#ff5555",
		"hicol3": "#ff5555",
		"hicol":  "#94A3A5",
		"bgcol":  "#282a36",
		"fgcol":  "#94A3A5",
		"errcol": "#ff0000",
	},

	"earthsong": map[string]string{
		"hicol2": "#c94234",
		"hicol3": "#c94234",
		"hicol":  "#e5c6aa",
		"bgcol":  "#292520",
		"fgcol":  "#e5c7a9",
		"errcol": "#ff0000",
	},

	"elemental": map[string]string{
		"hicol2": "#98290f",
		"hicol3": "#98290f",
		"hicol":  "#807974",
		"bgcol":  "#22211d",
		"fgcol":  "#807a74",
		"errcol": "#ff0000",
	},

	"elementary": map[string]string{
		"hicol2": "#e1321a",
		"hicol3": "#e1321a",
		"hicol":  "#f2f2f2",
		"bgcol":  "#101010",
		"fgcol":  "#f2f2f2",
		"errcol": "#ff0000",
	},

	"elic": map[string]string{
		"hicol2": "#e1321a",
		"hicol3": "#e1321a",
		"hicol":  "#2aa7e7",
		"bgcol":  "#4A453E",
		"fgcol":  "#f2f2f2",
		"errcol": "#ff0000",
	},

	"elio": map[string]string{
		"hicol2": "#e1321a",
		"hicol3": "#e1321a",
		"hicol":  "#f2f2f2",
		"bgcol":  "#041A3B",
		"fgcol":  "#f2f2f2",
		"errcol": "#ff0000",
	},

	"espresso-libre": map[string]string{
		"hicol2": "#cc0000",
		"hicol3": "#cc0000",
		"hicol":  "#d3d7cf",
		"bgcol":  "#2a211c",
		"fgcol":  "#b8a898",
		"errcol": "#ff0000",
	},

	"espresso": map[string]string{
		"hicol2": "#d25252",
		"hicol3": "#d25252",
		"hicol":  "#eeeeec",
		"bgcol":  "#323232",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"fishtank": map[string]string{
		"hicol2": "#c6004a",
		"hicol3": "#c6004a",
		"hicol":  "#ecf0fc",
		"bgcol":  "#232537",
		"fgcol":  "#ecf0fe",
		"errcol": "#ff0000",
	},

	"flatland": map[string]string{
		"hicol2": "#f18339",
		"hicol3": "#f18339",
		"hicol":  "#ffffff",
		"bgcol":  "#1d1f21",
		"fgcol":  "#b8dbef",
		"errcol": "#ff0000",
	},

	"flat": map[string]string{
		"hicol2": "#c0392b",
		"hicol3": "#c0392b",
		"hicol":  "#bdc3c7",
		"bgcol":  "#1F2D3A",
		"fgcol":  "#1abc9c",
		"errcol": "#ff0000",
	},

	"foxnightly": map[string]string{
		"hicol2": "#B98EFF",
		"hicol3": "#B98EFF",
		"hicol":  "#FFFFFF",
		"bgcol":  "#2A2A2E",
		"fgcol":  "#D7D7DB",
		"errcol": "#ff0000",
	},

	"freya": map[string]string{
		"hicol2": "#dc322f",
		"hicol3": "#dc322f",
		"hicol":  "#94a3a5",
		"bgcol":  "#252e32",
		"fgcol":  "#94a3a5",
		"errcol": "#ff0000",
	},

	"frontend-delight": map[string]string{
		"hicol2": "#f8511b",
		"hicol3": "#f8511b",
		"hicol":  "#adadad",
		"bgcol":  "#1b1c1d",
		"fgcol":  "#adadad",
		"errcol": "#ff0000",
	},

	"frontend-fun-forrest": map[string]string{
		"hicol2": "#d6262b",
		"hicol3": "#d6262b",
		"hicol":  "#ddc265",
		"bgcol":  "#251200",
		"fgcol":  "#dec165",
		"errcol": "#ff0000",
	},

	"frontend-galaxy": map[string]string{
		"hicol2": "#f9555f",
		"hicol3": "#f9555f",
		"hicol":  "#bbbbbb",
		"bgcol":  "#1d2837",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"github": map[string]string{
		"hicol2": "#970b16",
		"hicol3": "#970b16",
		"hicol":  "#ffffff",
		"bgcol":  "#f4f4f4",
		"fgcol":  "#3e3e3e",
		"errcol": "#ff0000",
	},

	"gooey": map[string]string{
		"hicol2": "#BB4F6C",
		"hicol3": "#BB4F6C",
		"hicol":  "#858893",
		"bgcol":  "#0D101B",
		"fgcol":  "#EBEEF9",
		"errcol": "#ff0000",
	},

	"google-dark": map[string]string{
		"hicol2": "#CC342B",
		"hicol3": "#CC342B",
		"hicol":  "#C5C8C6",
		"bgcol":  "#1D1F21",
		"fgcol":  "#B4B7B4",
		"errcol": "#ff0000",
	},

	"google-light": map[string]string{
		"hicol2": "#CC342B",
		"hicol3": "#CC342B",
		"hicol":  "#373B41",
		"bgcol":  "#FFFFFF",
		"fgcol":  "#373B41",
		"errcol": "#ff0000",
	},

	"grape": map[string]string{
		"hicol2": "#ed2261",
		"hicol3": "#ed2261",
		"hicol":  "#9e9ea0",
		"bgcol":  "#171423",
		"fgcol":  "#9f9fa1",
		"errcol": "#ff0000",
	},

	"grass": map[string]string{
		"hicol2": "#bb0000",
		"hicol3": "#bb0000",
		"hicol":  "#bbbbbb",
		"bgcol":  "#13773d",
		"fgcol":  "#fff0a5",
		"errcol": "#ff0000",
	},

	"gruvbox-dark": map[string]string{
		"hicol2": "#cc241d",
		"hicol3": "#cc241d",
		"hicol":  "#a89984",
		"bgcol":  "#282828",
		"fgcol":  "#ebdbb2",
		"errcol": "#ff0000",
	},

	"gruvbox": map[string]string{
		"hicol2": "#cc241d",
		"hicol3": "#cc241d",
		"hicol":  "#7c6f64",
		"bgcol":  "#fbf1c7",
		"fgcol":  "#3c3836",
		"errcol": "#ff0000",
	},

	"hardcore": map[string]string{
		"hicol2": "#f92672",
		"hicol3": "#f92672",
		"hicol":  "#ccccc6",
		"bgcol":  "#121212",
		"fgcol":  "#a0a0a0",
		"errcol": "#ff0000",
	},

	"harper": map[string]string{
		"hicol2": "#f8b63f",
		"hicol3": "#f8b63f",
		"hicol":  "#a8a49d",
		"bgcol":  "#010101",
		"fgcol":  "#a8a49d",
		"errcol": "#ff0000",
	},

	"hemisu-dark": map[string]string{
		"hicol2": "#FF0054",
		"hicol3": "#FF0054",
		"hicol":  "#EDEDED",
		"bgcol":  "#000000",
		"fgcol":  "#FFFFFF",
		"errcol": "#ff0000",
	},

	"hemisu-light": map[string]string{
		"hicol2": "#FF0055",
		"hicol3": "#FF0055",
		"hicol":  "#999999",
		"bgcol":  "#EFEFEF",
		"fgcol":  "#444444",
		"errcol": "#ff0000",
	},

	"highway": map[string]string{
		"hicol2": "#d00e18",
		"hicol3": "#d00e18",
		"hicol":  "#ededed",
		"bgcol":  "#222225",
		"fgcol":  "#ededed",
		"errcol": "#ff0000",
	},

	"hipster-green": map[string]string{
		"hicol2": "#b6214a",
		"hicol3": "#b6214a",
		"hicol":  "#bfbfbf",
		"bgcol":  "#100b05",
		"fgcol":  "#84c138",
		"errcol": "#ff0000",
	},

	"homebrew": map[string]string{
		"hicol2": "#990000",
		"hicol3": "#990000",
		"hicol":  "#bfbfbf",
		"bgcol":  "#000000",
		"fgcol":  "#00ff00",
		"errcol": "#ff0000",
	},

	"hurtado": map[string]string{
		"hicol2": "#ff1b00",
		"hicol3": "#ff1b00",
		"hicol":  "#cbcccb",
		"bgcol":  "#000000",
		"fgcol":  "#dbdbdb",
		"errcol": "#ff0000",
	},

	"hybrid": map[string]string{
		"hicol2": "#A54242",
		"hicol3": "#A54242",
		"hicol":  "#969896",
		"bgcol":  "#141414",
		"fgcol":  "#94a3a5",
		"errcol": "#ff0000",
	},

	"ibm3270": map[string]string{
		"hicol2": "#F01818",
		"hicol3": "#F01818",
		"hicol":  "#A5A5A5",
		"bgcol":  "#000000",
		"fgcol":  "#FDFDFD",
		"errcol": "#ff0000",
	},

	"ic-green-ppl": map[string]string{
		"hicol2": "#fb002a",
		"hicol3": "#fb002a",
		"hicol":  "#e0ffef",
		"bgcol":  "#3a3d3f",
		"fgcol":  "#d9efd3",
		"errcol": "#ff0000",
	},

	"ic-orange-ppl": map[string]string{
		"hicol2": "#c13900",
		"hicol3": "#c13900",
		"hicol":  "#ffc88a",
		"bgcol":  "#262626",
		"fgcol":  "#ffcb83",
		"errcol": "#ff0000",
	},

	"idle-toes": map[string]string{
		"hicol2": "#d25252",
		"hicol3": "#d25252",
		"hicol":  "#eeeeec",
		"bgcol":  "#323232",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"ir-black": map[string]string{
		"hicol2": "#ff6c60",
		"hicol3": "#ff6c60",
		"hicol":  "#eeeeee",
		"bgcol":  "#000000",
		"fgcol":  "#eeeeee",
		"errcol": "#ff0000",
	},

	"jackie-brown": map[string]string{
		"hicol2": "#ef5734",
		"hicol3": "#ef5734",
		"hicol":  "#bfbfbf",
		"bgcol":  "#2c1d16",
		"fgcol":  "#ffcc2f",
		"errcol": "#ff0000",
	},

	"japanesque": map[string]string{
		"hicol2": "#cf3f61",
		"hicol3": "#cf3f61",
		"hicol":  "#fafaf6",
		"bgcol":  "#1e1e1e",
		"fgcol":  "#f7f6ec",
		"errcol": "#ff0000",
	},

	"jellybeans": map[string]string{
		"hicol2": "#e27373",
		"hicol3": "#e27373",
		"hicol":  "#dedede",
		"bgcol":  "#121212",
		"fgcol":  "#dedede",
		"errcol": "#ff0000",
	},

	"jk": map[string]string{
		"errcol": "#ff0000",
	},

	"jup": map[string]string{
		"hicol2": "#dd006f",
		"hicol3": "#dd006f",
		"hicol":  "#f2f2f2",
		"bgcol":  "#758480",
		"fgcol":  "#23476a",
		"errcol": "#ff0000",
	},

	"kibble": map[string]string{
		"hicol2": "#c70031",
		"hicol3": "#c70031",
		"hicol":  "#e2d1e3",
		"bgcol":  "#0e100a",
		"fgcol":  "#f7f7f7",
		"errcol": "#ff0000",
	},

	"later-this-evening": map[string]string{
		"hicol2": "#d45a60",
		"hicol3": "#d45a60",
		"hicol":  "#3c3d3d",
		"bgcol":  "#222222",
		"fgcol":  "#959595",
		"errcol": "#ff0000",
	},

	"lavandula": map[string]string{
		"hicol2": "#7d1625",
		"hicol3": "#7d1625",
		"hicol":  "#736e7d",
		"bgcol":  "#050014",
		"fgcol":  "#736e7d",
		"errcol": "#ff0000",
	},

	"liquid-carbon-transparent": map[string]string{
		"hicol2": "#ff3030",
		"hicol3": "#ff3030",
		"hicol":  "#bccccc",
		"bgcol":  "#000000",
		"fgcol":  "#afc2c2",
		"errcol": "#ff0000",
	},

	"liquid-carbon": map[string]string{
		"hicol2": "#ff3030",
		"hicol3": "#ff3030",
		"hicol":  "#bccccc",
		"bgcol":  "#303030",
		"fgcol":  "#afc2c2",
		"errcol": "#ff0000",
	},

	"maia": map[string]string{
		"hicol2": "#BA2922",
		"hicol3": "#BA2922",
		"hicol":  "#E0E0E0",
		"bgcol":  "#31363B",
		"fgcol":  "#BDX3C7",
		"errcol": "#ff0000",
	},

	"man-page": map[string]string{
		"hicol2": "#cc0000",
		"hicol3": "#cc0000",
		"hicol":  "#cccccc",
		"bgcol":  "#fef49c",
		"fgcol":  "#000000",
		"errcol": "#ff0000",
	},

	"mar": map[string]string{
		"hicol2": "#b5407b",
		"hicol3": "#b5407b",
		"hicol":  "#f8f8f8",
		"bgcol":  "#ffffff",
		"fgcol":  "#23476a",
		"errcol": "#ff0000",
	},

	"material": map[string]string{
		"hicol2": "#EB606B",
		"hicol3": "#EB606B",
		"hicol":  "#FFFFFF",
		"bgcol":  "#1E282C",
		"fgcol":  "#C3C7D1",
		"errcol": "#ff0000",
	},

	"mathias": map[string]string{
		"hicol2": "#e52222",
		"hicol3": "#e52222",
		"hicol":  "#f2f2f2",
		"bgcol":  "#000000",
		"fgcol":  "#bbbbbb",
		"errcol": "#ff0000",
	},

	"medallion": map[string]string{
		"hicol2": "#b64c00",
		"hicol3": "#b64c00",
		"hicol":  "#cac29a",
		"bgcol":  "#1d1908",
		"fgcol":  "#cac296",
		"errcol": "#ff0000",
	},

	"misterioso": map[string]string{
		"hicol2": "#ff4242",
		"hicol3": "#ff4242",
		"hicol":  "#e1e1e0",
		"bgcol":  "#2d3743",
		"fgcol":  "#e1e1e0",
		"errcol": "#ff0000",
	},

	"miu": map[string]string{
		"hicol2": "#b87a7a",
		"hicol3": "#b87a7a",
		"hicol":  "#d9d9d9",
		"bgcol":  "#0d1926",
		"fgcol":  "#d9e6f2",
		"errcol": "#ff0000",
	},

	"molokai": map[string]string{
		"hicol2": "#7325FA",
		"hicol3": "#7325FA",
		"hicol":  "#BBBBBB",
		"bgcol":  "#1b1d1e",
		"fgcol":  "#BBBBBB",
		"errcol": "#ff0000",
	},

	"mona-lisa": map[string]string{
		"hicol2": "#9b291c",
		"hicol3": "#9b291c",
		"hicol":  "#f7d75c",
		"bgcol":  "#120b0d",
		"fgcol":  "#f7d66a",
		"errcol": "#ff0000",
	},

	"mono-amber": map[string]string{
		"hicol2": "#FF9400",
		"hicol3": "#FF9400",
		"hicol":  "#FF9400",
		"bgcol":  "#2B1900",
		"fgcol":  "#FF9400",
		"errcol": "#ff0000",
	},

	"mono-cyan": map[string]string{
		"hicol2": "#00CCFF",
		"hicol3": "#00CCFF",
		"hicol":  "#00CCFF",
		"bgcol":  "#00222B",
		"fgcol":  "#00CCFF",
		"errcol": "#ff0000",
	},

	"mono-green": map[string]string{
		"hicol2": "#0BFF00",
		"hicol3": "#0BFF00",
		"hicol":  "#0BFF00",
		"bgcol":  "#022B00",
		"fgcol":  "#0BFF00",
		"errcol": "#ff0000",
	},

	"monokai-dark": map[string]string{
		"hicol2": "#f92672",
		"hicol3": "#f92672",
		"hicol":  "#f9f8f5",
		"bgcol":  "#272822",
		"fgcol":  "#f8f8f2",
		"errcol": "#ff0000",
	},

	"monokai-soda": map[string]string{
		"hicol2": "#f4005f",
		"hicol3": "#f4005f",
		"hicol":  "#c4c5b5",
		"bgcol":  "#1a1a1a",
		"fgcol":  "#c4c5b5",
		"errcol": "#ff0000",
	},

	"mono-red": map[string]string{
		"hicol2": "#FF3600",
		"hicol3": "#FF3600",
		"hicol":  "#FF3600",
		"bgcol":  "#2B0C00",
		"fgcol":  "#FF3600",
		"errcol": "#ff0000",
	},

	"mono-white": map[string]string{
		"hicol2": "#FAFAFA",
		"hicol3": "#FAFAFA",
		"hicol":  "#FAFAFA",
		"bgcol":  "#262626",
		"fgcol":  "#FAFAFA",
		"errcol": "#ff0000",
	},

	"mono-yellow": map[string]string{
		"hicol2": "#FFD300",
		"hicol3": "#FFD300",
		"hicol":  "#FFD300",
		"bgcol":  "#2B2400",
		"fgcol":  "#FFD300",
		"errcol": "#ff0000",
	},

	"n0tch2k": map[string]string{
		"hicol2": "#a95551",
		"hicol3": "#a95551",
		"hicol":  "#d0b8a3",
		"bgcol":  "#222222",
		"fgcol":  "#a0a0a0",
		"errcol": "#ff0000",
	},

	"neon-night": map[string]string{
		"hicol2": "#FF8E8E",
		"hicol3": "#FF8E8E",
		"hicol":  "#C9CCCD",
		"bgcol":  "#20242d",
		"fgcol":  "#C7C8FF",
		"errcol": "#ff0000",
	},

	"neopolitan": map[string]string{
		"hicol2": "#800000",
		"hicol3": "#800000",
		"hicol":  "#f8f8f8",
		"bgcol":  "#271f19",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"nep": map[string]string{
		"hicol2": "#dd6f00",
		"hicol3": "#dd6f00",
		"hicol":  "#f2f2f2",
		"bgcol":  "#758480",
		"fgcol":  "#23476a",
		"errcol": "#ff0000",
	},

	"neutron": map[string]string{
		"hicol2": "#b54036",
		"hicol3": "#b54036",
		"hicol":  "#e6e8ef",
		"bgcol":  "#1c1e22",
		"fgcol":  "#e6e8ef",
		"errcol": "#ff0000",
	},

	"nightlion-v1": map[string]string{
		"hicol2": "#bb0000",
		"hicol3": "#bb0000",
		"hicol":  "#bbbbbb",
		"bgcol":  "#000000",
		"fgcol":  "#bbbbbb",
		"errcol": "#ff0000",
	},

	"nightlion-v2": map[string]string{
		"hicol2": "#bb0000",
		"hicol3": "#bb0000",
		"hicol":  "#bbbbbb",
		"bgcol":  "#171717",
		"fgcol":  "#bbbbbb",
		"errcol": "#ff0000",
	},

	"nighty": map[string]string{
		"hicol2": "#9B3E46",
		"hicol3": "#9B3E46",
		"hicol":  "#828282",
		"bgcol":  "#2F2F2F",
		"fgcol":  "#DFDFDF",
		"errcol": "#ff0000",
	},

	"nord-light": map[string]string{
		"hicol2": "#E64569",
		"hicol3": "#E64569",
		"hicol":  "#B3B3B3",
		"bgcol":  "#ebeaf2",
		"fgcol":  "#004f7c",
		"errcol": "#ff0000",
	},

	"nord": map[string]string{
		"hicol2": "#E64569",
		"hicol3": "#E64569",
		"hicol":  "#B3B3B3",
		"errcol": "#ff0000",
	},

	"novel": map[string]string{
		"hicol2": "#cc0000",
		"hicol3": "#cc0000",
		"hicol":  "#cccccc",
		"bgcol":  "#dfdbc3",
		"fgcol":  "#3b2322",
		"errcol": "#ff0000",
	},

	"obsidian": map[string]string{
		"hicol2": "#a60001",
		"hicol3": "#a60001",
		"hicol":  "#bbbbbb",
		"bgcol":  "#283033",
		"fgcol":  "#cdcdcd",
		"errcol": "#ff0000",
	},

	"ocean-dark": map[string]string{
		"hicol2": "#AF4B57",
		"hicol3": "#AF4B57",
		"hicol":  "#EEEDEE",
		"bgcol":  "#1C1F27",
		"fgcol":  "#979CAC",
		"errcol": "#ff0000",
	},

	"oceanic-next": map[string]string{
		"hicol2": "#E44754",
		"hicol3": "#E44754",
		"hicol":  "#FFFFFF",
		"bgcol":  "#121b21",
		"fgcol":  "#b3b8c3",
		"errcol": "#ff0000",
	},

	"ocean": map[string]string{
		"hicol2": "#990000",
		"hicol3": "#990000",
		"hicol":  "#bfbfbf",
		"bgcol":  "#224fbc",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"ollie": map[string]string{
		"hicol2": "#ac2e31",
		"hicol3": "#ac2e31",
		"hicol":  "#8a8eac",
		"bgcol":  "#222125",
		"fgcol":  "#8a8dae",
		"errcol": "#ff0000",
	},

	"one-dark": map[string]string{
		"hicol2": "#E06C75",
		"hicol3": "#E06C75",
		"hicol":  "#ABB2BF",
		"bgcol":  "#1E2127",
		"fgcol":  "#5C6370",
		"errcol": "#ff0000",
	},

	"one-half-black": map[string]string{
		"hicol2": "#e06c75",
		"hicol3": "#e06c75",
		"hicol":  "#dcdfe4",
		"bgcol":  "#000000",
		"fgcol":  "#dcdfe4",
		"errcol": "#ff0000",
	},

	"one-light": map[string]string{
		"hicol2": "#DA3E39",
		"hicol3": "#DA3E39",
		"hicol":  "#8E8F96",
		"bgcol":  "#F8F8F8",
		"fgcol":  "#2A2B32",
		"errcol": "#ff0000",
	},

	"pali": map[string]string{
		"hicol2": "#ab8f74",
		"hicol3": "#ab8f74",
		"hicol":  "#F2F2F2",
		"bgcol":  "#232E37",
		"fgcol":  "#d9e6f2",
		"errcol": "#ff0000",
	},

	"papercolor-dark": map[string]string{
		"hicol2": "#AF005F",
		"hicol3": "#AF005F",
		"hicol":  "#D0D0D0",
		"bgcol":  "#1C1C1C",
		"fgcol":  "#D0D0D0",
		"errcol": "#ff0000",
	},

	"papercolor-light": map[string]string{
		"hicol2": "#AF0000",
		"hicol3": "#AF0000",
		"hicol":  "#444444",
		"bgcol":  "#EEEEEE",
		"fgcol":  "#444444",
		"errcol": "#ff0000",
	},

	"paraiso-dark": map[string]string{
		"hicol2": "#ef6155",
		"hicol3": "#ef6155",
		"hicol":  "#a39e9b",
		"bgcol":  "#2f1e2e",
		"fgcol":  "#a39e9b",
		"errcol": "#ff0000",
	},

	"paul-millr": map[string]string{
		"hicol2": "#ff0000",
		"hicol3": "#ff0000",
		"hicol":  "#bbbbbb",
		"bgcol":  "#000000",
		"fgcol":  "#f2f2f2",
		"errcol": "#ff0000",
	},

	"pencil-dark": map[string]string{
		"hicol2": "#c30771",
		"hicol3": "#c30771",
		"hicol":  "#d9d9d9",
		"bgcol":  "#212121",
		"fgcol":  "#f1f1f1",
		"errcol": "#ff0000",
	},

	"pencil-light": map[string]string{
		"hicol2": "#c30771",
		"hicol3": "#c30771",
		"hicol":  "#d9d9d9",
		"bgcol":  "#f1f1f1",
		"fgcol":  "#424242",
		"errcol": "#ff0000",
	},

	"peppermint": map[string]string{
		"hicol2": "#E64569",
		"hicol3": "#E64569",
		"hicol":  "#B3B3B3",
		"bgcol":  "#000000",
		"fgcol":  "#C7C7C7",
		"errcol": "#ff0000",
	},

	"pnevma": map[string]string{
		"hicol2": "#a36666",
		"hicol3": "#a36666",
		"hicol":  "#d0d0d0",
		"bgcol":  "#1c1c1c",
		"fgcol":  "#d0d0d0",
		"errcol": "#ff0000",
	},

	"pro": map[string]string{
		"hicol2": "#990000",
		"hicol3": "#990000",
		"hicol":  "#bfbfbf",
		"bgcol":  "#000000",
		"fgcol":  "#f2f2f2",
		"errcol": "#ff0000",
	},

	"README": map[string]string{
		"errcol": "#ff0000",
	},

	"red-alert": map[string]string{
		"hicol2": "#d62e4e",
		"hicol3": "#d62e4e",
		"hicol":  "#d6d6d6",
		"bgcol":  "#762423",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"red-sands": map[string]string{
		"hicol2": "#ff3f00",
		"hicol3": "#ff3f00",
		"hicol":  "#bbbbbb",
		"bgcol":  "#7a251e",
		"fgcol":  "#d7c9a7",
		"errcol": "#ff0000",
	},

	"rippedcasts": map[string]string{
		"hicol2": "#cdaf95",
		"hicol3": "#cdaf95",
		"hicol":  "#bfbfbf",
		"bgcol":  "#2b2b2b",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"royal": map[string]string{
		"hicol2": "#91284c",
		"hicol3": "#91284c",
		"hicol":  "#524966",
		"bgcol":  "#100815",
		"fgcol":  "#514968",
		"errcol": "#ff0000",
	},

	"sat": map[string]string{
		"hicol2": "#dd0007",
		"hicol3": "#dd0007",
		"hicol":  "#f2f2f2",
		"bgcol":  "#758480",
		"fgcol":  "#23476a",
		"errcol": "#ff0000",
	},

	"seafoam-pastel": map[string]string{
		"hicol2": "#825d4d",
		"hicol3": "#825d4d",
		"hicol":  "#e0e0e0",
		"bgcol":  "#243435",
		"fgcol":  "#d4e7d4",
		"errcol": "#ff0000",
	},

	"sea-shells": map[string]string{
		"hicol2": "#d15123",
		"hicol3": "#d15123",
		"hicol":  "#deb88d",
		"bgcol":  "#09141b",
		"fgcol":  "#deb88d",
		"errcol": "#ff0000",
	},

	"seti": map[string]string{
		"hicol2": "#c22832",
		"hicol3": "#c22832",
		"hicol":  "#eeeeee",
		"bgcol":  "#111213",
		"fgcol":  "#cacecd",
		"errcol": "#ff0000",
	},

	"shaman": map[string]string{
		"hicol2": "#b2302d",
		"hicol3": "#b2302d",
		"hicol":  "#405555",
		"bgcol":  "#001015",
		"fgcol":  "#405555",
		"errcol": "#ff0000",
	},

	"shel": map[string]string{
		"hicol2": "#ab2463",
		"hicol3": "#ab2463",
		"hicol":  "#918988",
		"bgcol":  "#2a201f",
		"fgcol":  "#4882cd",
		"errcol": "#ff0000",
	},

	"slate": map[string]string{
		"hicol2": "#e2a8bf",
		"hicol3": "#e2a8bf",
		"hicol":  "#02c5e0",
		"bgcol":  "#222222",
		"fgcol":  "#35b1d2",
		"errcol": "#ff0000",
	},

	"smyck": map[string]string{
		"hicol2": "#C75646",
		"hicol3": "#C75646",
		"hicol":  "#B0B0B0",
		"bgcol":  "#242424",
		"fgcol":  "#F7F7F7",
		"errcol": "#ff0000",
	},

	"snazzy": map[string]string{
		"hicol2": "#FF5C57",
		"hicol3": "#FF5C57",
		"hicol":  "#F1F1F0",
		"errcol": "#ff0000",
	},

	"soft-server": map[string]string{
		"hicol2": "#a2686a",
		"hicol3": "#a2686a",
		"hicol":  "#99a3a2",
		"bgcol":  "#242626",
		"fgcol":  "#99a3a2",
		"errcol": "#ff0000",
	},

	"solarized-darcula": map[string]string{
		"hicol2": "#f24840",
		"hicol3": "#f24840",
		"hicol":  "#d2d8d9",
		"bgcol":  "#3d3f41",
		"fgcol":  "#d2d8d9",
		"errcol": "#ff0000",
	},

	"solarized-dark-higher-contrast": map[string]string{
		"hicol2": "#d11c24",
		"hicol3": "#d11c24",
		"hicol":  "#eae3cb",
		"bgcol":  "#001e27",
		"fgcol":  "#9cc2c3",
		"errcol": "#ff0000",
	},

	"solarized-dark": map[string]string{
		"hicol2": "#DC322F",
		"hicol3": "#DC322F",
		"hicol":  "#EEE8D5",
		"bgcol":  "#002B36",
		"fgcol":  "#839496",
		"errcol": "#ff0000",
	},

	"solarized-light": map[string]string{
		"hicol2": "#859900",
		"hicol3": "#859900",
		"hicol":  "#002B36",
		"bgcol":  "#FDF6E3",
		"fgcol":  "#657B83",
		"errcol": "#ff0000",
	},

	"spacedust": map[string]string{
		"hicol2": "#e35b00",
		"hicol3": "#e35b00",
		"hicol":  "#f0f1ce",
		"bgcol":  "#0a1e24",
		"fgcol":  "#ecf0c1",
		"errcol": "#ff0000",
	},

	"spacegray-eighties-dull": map[string]string{
		"hicol2": "#b24a56",
		"hicol3": "#b24a56",
		"hicol":  "#b3b8c3",
		"bgcol":  "#222222",
		"fgcol":  "#c9c6bc",
		"errcol": "#ff0000",
	},

	"spacegray-eighties": map[string]string{
		"hicol2": "#ec5f67",
		"hicol3": "#ec5f67",
		"hicol":  "#efece7",
		"bgcol":  "#222222",
		"fgcol":  "#bdbaae",
		"errcol": "#ff0000",
	},

	"spacegray": map[string]string{
		"hicol2": "#b04b57",
		"hicol3": "#b04b57",
		"hicol":  "#b3b8c3",
		"bgcol":  "#20242d",
		"fgcol":  "#b3b8c3",
		"errcol": "#ff0000",
	},

	"spring": map[string]string{
		"hicol2": "#ff4d83",
		"hicol3": "#ff4d83",
		"hicol":  "#ffffff",
		"bgcol":  "#0a1e24",
		"fgcol":  "#ecf0c1",
		"errcol": "#ff0000",
	},

	"square": map[string]string{
		"hicol2": "#e9897c",
		"hicol3": "#e9897c",
		"hicol":  "#f2f2f2",
		"bgcol":  "#0a1e24",
		"fgcol":  "#1a1a1a",
		"errcol": "#ff0000",
	},

	"srcery": map[string]string{
		"hicol2": "#FF3128",
		"hicol3": "#FF3128",
		"hicol":  "#918175",
		"bgcol":  "#282828",
		"fgcol":  "#ebdbb2",
		"errcol": "#ff0000",
	},

	"sundried": map[string]string{
		"hicol2": "#a7463d",
		"hicol3": "#a7463d",
		"hicol":  "#c9c9c9",
		"bgcol":  "#1a1818",
		"fgcol":  "#c9c9c9",
		"errcol": "#ff0000",
	},

	"symphonic": map[string]string{
		"hicol2": "#dc322f",
		"hicol3": "#dc322f",
		"hicol":  "#ffffff",
		"bgcol":  "#000000",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"teerb": map[string]string{
		"hicol2": "#d68686",
		"hicol3": "#d68686",
		"hicol":  "#d0d0d0",
		"bgcol":  "#262626",
		"fgcol":  "#d0d0d0",
		"errcol": "#ff0000",
	},

	"terminal-basic": map[string]string{
		"hicol2": "#990000",
		"hicol3": "#990000",
		"hicol":  "#bfbfbf",
		"bgcol":  "#ffffff",
		"fgcol":  "#000000",
		"errcol": "#ff0000",
	},

	"terminix-dark": map[string]string{
		"hicol2": "#a54242",
		"hicol3": "#a54242",
		"hicol":  "#777777",
		"bgcol":  "#091116",
		"fgcol":  "#868A8C",
		"errcol": "#ff0000",
	},

	"thayer-bright": map[string]string{
		"hicol2": "#f92672",
		"hicol3": "#f92672",
		"hicol":  "#ccccc6",
		"bgcol":  "#1b1d1e",
		"fgcol":  "#f8f8f8",
		"errcol": "#ff0000",
	},

	"tin": map[string]string{
		"hicol2": "#8d534e",
		"hicol3": "#8d534e",
		"hicol":  "#ffffff",
		"bgcol":  "#2e2e35",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"tomorrow-night-blue": map[string]string{
		"hicol2": "#FF9DA3",
		"hicol3": "#FF9DA3",
		"hicol":  "#FFFEFE",
		"bgcol":  "#002451",
		"fgcol":  "#FFFEFE",
		"errcol": "#ff0000",
	},

	"tomorrow-night-bright": map[string]string{
		"hicol2": "#D54E53",
		"hicol3": "#D54E53",
		"hicol":  "#FFFEFE",
		"bgcol":  "#000000",
		"fgcol":  "#E9E9E9",
		"errcol": "#ff0000",
	},

	"tomorrow-night-eighties": map[string]string{
		"hicol2": "#F27779",
		"hicol3": "#F27779",
		"hicol":  "#FFFEFE",
		"bgcol":  "#2C2C2C",
		"fgcol":  "#CCCCCC",
		"errcol": "#ff0000",
	},

	"tomorrow-night": map[string]string{
		"hicol2": "#CC6666",
		"hicol3": "#CC6666",
		"hicol":  "#FFFEFE",
		"bgcol":  "#1D1F21",
		"fgcol":  "#C5C8C6",
		"errcol": "#ff0000",
	},

	"tomorrow": map[string]string{
		"hicol2": "#C82828",
		"hicol3": "#C82828",
		"hicol":  "#FFFEFE",
		"bgcol":  "#FFFFFF",
		"fgcol":  "#4D4D4C",
		"errcol": "#ff0000",
	},

	"toy-chest": map[string]string{
		"hicol2": "#be2d26",
		"hicol3": "#be2d26",
		"hicol":  "#23d183",
		"bgcol":  "#24364b",
		"fgcol":  "#31d07b",
		"errcol": "#ff0000",
	},

	"treehouse": map[string]string{
		"hicol2": "#b2270e",
		"hicol3": "#b2270e",
		"hicol":  "#786b53",
		"bgcol":  "#191919",
		"fgcol":  "#786b53",
		"errcol": "#ff0000",
	},

	"twilight": map[string]string{
		"hicol2": "#c06d44",
		"hicol3": "#c06d44",
		"hicol":  "#ffffd4",
		"bgcol":  "#141414",
		"fgcol":  "#ffffd4",
		"errcol": "#ff0000",
	},

	"ura": map[string]string{
		"hicol2": "#c21b6f",
		"hicol3": "#c21b6f",
		"hicol":  "#808080",
		"bgcol":  "#feffee",
		"fgcol":  "#23476a",
		"errcol": "#ff0000",
	},

	"urple": map[string]string{
		"hicol2": "#b0425b",
		"hicol3": "#b0425b",
		"hicol":  "#87799c",
		"bgcol":  "#1b1b23",
		"fgcol":  "#877a9b",
		"errcol": "#ff0000",
	},

	"vag": map[string]string{
		"hicol2": "#a87139",
		"hicol3": "#a87139",
		"hicol":  "#8a8a8a",
		"bgcol":  "#191f1d",
		"fgcol":  "#d9e6f2",
		"errcol": "#ff0000",
	},

	"vaughn": map[string]string{
		"hicol2": "#705050",
		"hicol3": "#705050",
		"hicol":  "#709080",
		"bgcol":  "#25234f",
		"fgcol":  "#dcdccc",
		"errcol": "#ff0000",
	},

	"vibrant-ink": map[string]string{
		"hicol2": "#ff6600",
		"hicol3": "#ff6600",
		"hicol":  "#f5f5f5",
		"bgcol":  "#000000",
		"fgcol":  "#ffffff",
		"errcol": "#ff0000",
	},

	"vs-code-dark-plus": map[string]string{
		"hicol2": "#E9653B",
		"hicol3": "#E9653B",
		"hicol":  "#C3DDE1",
		"bgcol":  "#1E1E1E",
		"fgcol":  "#CCCCCC",
		"errcol": "#ff0000",
	},

	"warm-neon": map[string]string{
		"hicol2": "#e24346",
		"hicol3": "#e24346",
		"hicol":  "#d0b8a3",
		"bgcol":  "#404040",
		"fgcol":  "#afdab6",
		"errcol": "#ff0000",
	},

	"wez": map[string]string{
		"hicol2": "#cc5555",
		"hicol3": "#cc5555",
		"hicol":  "#cccccc",
		"bgcol":  "#000000",
		"fgcol":  "#b3b3b3",
		"errcol": "#ff0000",
	},

	"wild-cherry": map[string]string{
		"hicol2": "#d94085",
		"hicol3": "#d94085",
		"hicol":  "#fff8de",
		"bgcol":  "#1f1726",
		"fgcol":  "#dafaff",
		"errcol": "#ff0000",
	},

	"wombat": map[string]string{
		"hicol2": "#ff615a",
		"hicol3": "#ff615a",
		"hicol":  "#dedacf",
		"bgcol":  "#171717",
		"fgcol":  "#dedacf",
		"errcol": "#ff0000",
	},

	"wryan": map[string]string{
		"hicol2": "#8c4665",
		"hicol3": "#8c4665",
		"hicol":  "#899ca1",
		"bgcol":  "#101010",
		"fgcol":  "#999993",
		"errcol": "#ff0000",
	},

	"zenburn": map[string]string{
		"hicol2": "#705050",
		"hicol3": "#705050",
		"hicol":  "#dcdccc",
		"bgcol":  "#3f3f3f",
		"fgcol":  "#dcdccc",
		"errcol": "#ff0000",
	},
}
