package main

//Themes defined here take precedence over their generated counterparts and should be manually curated.

var themes = map[string]map[string]string{
	"default": {
		"fgcol": "#8C8C8C",
		"bgcol": "#282828",

		"hicol2": "#805b13",
		"hicol3": "#b4801b",
		"hicol":  "#ffffff",
		"errcol": "#a10705",
	},
}

func init() {
	for k, v := range generatedThemes {
		if _, exists := themes[k]; !exists {
			themes[k] = v
		}
	}
}
