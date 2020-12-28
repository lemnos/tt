import os
import json

# Generates themes from terminal_themes.json (itself generated from some long forgotten github repo [possibly guake]).

f = os.path.join(os.path.dirname(os.path.abspath(__file__)),
                 'terminal_themes.json')

themes = json.loads(open(f).read())


def blend(src, dest, opacity):
    src = src[1:]
    dest = dest[1:]

    dr, dg, db = [int(dest[i:i + 2], 16) for i in (0, 2, 4)]
    sr, sg, sb = [int(src[i:i + 2], 16) for i in (0, 2, 4)]

    sr = sr * opacity + dr * (1 - opacity)
    sg = sg * opacity + dg * (1 - opacity)
    sb = sb * opacity + db * (1 - opacity)

    return "#%.2x%.2x%.2x" % (int(sr), int(sg), int(sb))


print("//GENERATED CODE, DO NOT EDIT BY HAND (see themegen.py)\n\n")
print("package main\n")
print("var generatedThemes = map[string]map[string]string{")

for name, t in themes.items():
    print('\t"%s": map[string]string{' % name)

    # Meat (alter these to taste)
    mapping = {
        "bgcol": t['background'],
        "fgcol": t['foreground'],
        "hicol": t['color7'],
        "hicol2": blend(t['background'], t['color9'], .3),
        "hicol3": t['color9'],
        "errcol": t['color1'],
    }

    for k, v in mapping.items():
        print('\t\t"%s": "%s",' % (k, v))

    print("\t},")

print("}")
