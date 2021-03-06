package main

import (
	"strings"

	"github.com/xyproto/vt100"
)

func (e *Editor) gitHighlight(line string) string {
	var coloredString string
	if strings.HasPrefix(line, "#") {
		filenameColor := vt100.Red
		renameColor := vt100.Magenta
		if strings.HasPrefix(line, "# On branch ") {
			coloredString = vt100.DarkGray.Get(line[:12]) + vt100.LightCyan.Get(line[12:])
		} else if strings.HasPrefix(line, "# Your branch is up to date with '") && strings.Count(line, "'") == 2 {
			parts := strings.SplitN(line, "'", 3)
			coloredString = vt100.DarkGray.Get(parts[0]+"'") + vt100.LightGreen.Get(parts[1]) + vt100.DarkGray.Get("'"+parts[2])
		} else if line == "# Changes to be committed:" {
			coloredString = vt100.DarkGray.Get("# ") + vt100.LightBlue.Get("Changes to be committed:")
		} else if line == "# Untracked files:" {
			coloredString = vt100.DarkGray.Get("# ") + vt100.LightBlue.Get("Untracked files:")
		} else if strings.Contains(line, "new file:") {
			parts := strings.SplitN(line[1:], ":", 2)
			coloredString = vt100.DarkGray.Get("#") + vt100.LightYellow.Get(parts[0]) + vt100.DarkGray.Get(":") + filenameColor.Get(parts[1])
		} else if strings.Contains(line, "modified:") {
			parts := strings.SplitN(line[1:], ":", 2)
			coloredString = vt100.DarkGray.Get("#") + vt100.LightYellow.Get(parts[0]) + vt100.DarkGray.Get(":") + filenameColor.Get(parts[1])
		} else if strings.Contains(line, "deleted:") {
			parts := strings.SplitN(line[1:], ":", 2)
			coloredString = vt100.DarkGray.Get("#") + vt100.LightYellow.Get(parts[0]) + vt100.DarkGray.Get(":") + filenameColor.Get(parts[1])
		} else if strings.Contains(line, "renamed:") {
			parts := strings.SplitN(line[1:], ":", 2)
			if strings.Contains(parts[1], "->") {
				filenames := strings.SplitN(parts[1], "->", 2)
				coloredString = vt100.DarkGray.Get("#") + vt100.LightYellow.Get(parts[0]) + vt100.DarkGray.Get(":") + renameColor.Get(filenames[0]) + vt100.White.Get("->") + renameColor.Get(filenames[1])
			} else {
				coloredString = vt100.DarkGray.Get("#") + vt100.LightYellow.Get(parts[0]) + vt100.DarkGray.Get(":") + filenameColor.Get(parts[1])
			}
		} else if fields := strings.Fields(line); strings.HasPrefix(line, "# Rebase ") && len(fields) >= 5 && strings.Contains(fields[2], "..") {
			textColor := vt100.LightGray
			commitRange := strings.SplitN(fields[2], "..", 2)
			coloredString = vt100.DarkGray.Get("# ") + textColor.Get(fields[1]) + " " + vt100.LightBlue.Get(commitRange[0]) + textColor.Get("..") + vt100.LightBlue.Get(commitRange[1]) + " " + textColor.Get(fields[3]) + " " + vt100.LightBlue.Get(fields[4]) + " " + textColor.Get(strings.Join(fields[5:], " "))
		} else {
			coloredString = vt100.DarkGray.Get(line)
		}
	} else if fields := strings.Fields(line); len(fields) >= 3 && hasAnyPrefixWord(line, []string{"p", "pick", "r", "reword", "e", "edit", "s", "squash", "f", "fixup", "x", "exec", "b", "break", "d", "drop", "l", "label", "t", "reset", "m", "merge"}) {
		coloredString = vt100.Red.Get(fields[0]) + " " + vt100.LightBlue.Get(fields[1]) + " " + vt100.LightGray.Get(strings.Join(fields[2:], " "))
	} else {
		coloredString = e.gitColor.Get(line)
	}
	return coloredString

}
