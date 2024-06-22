package fnetutils

func PrintBanner() {
	for _, line := range FNetBanner() {
		println(line)
	}
}

func FNetBanner() []string {
	return []string{
		"    ______   _   __         __     ",
		"   / ____/  / | / /  ___   / /_    ",
		"  / /_     /  |/ /  / _ \\\\ / __/ ",
		" / __/    / /|  /  /  __// /_      ",
		"/_/      /_/ |_/   \\\\___/ \\__/  ",
	}
}
