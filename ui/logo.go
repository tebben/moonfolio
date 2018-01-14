package ui

import (
	"fmt"
)

// GetLogo returns the strings to construct the Moonfolio logo
func GetLogo() []string {
	return []string{
		fmt.Sprintf("%s       _.._    %s%s__  __  ____   ____  _   _ ______ ____  _      _____ ____", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s    .' .-'`   %s%s|  \\/  |/ __ \\ / __ \\| \\ | |  ____/ __ \\| |    |_   _/ __ \\", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s   /  /       %s%s| \\  / | |  | | |  | |  \\| | |__ | |  | | |      | || |  | |", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s   |  |       %s%s| |\\/| | |  | | |  | | . ` |  __|| |  | | |      | || |  | |", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s   \\  \\       %s%s| |  | | |__| | |__| | |\\  | |   | |__| | |____ _| || |__| |", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s    '._'-._   %s%s|_|  |_|\\____/ \\____/|_| \\_|_|    \\____/|______|_____\\____/ ", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s       ```    %s", ColorWhite, BoldEnd),
	}
}

// GetLogoV2 returns the strings to construct the Moonfolio logo
func GetLogoV2() []string {
	return []string{
		fmt.Sprintf("%s%s       _.._   %s%s", BoldStart, ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s    .' .-'`   %s%s███╗   ███╗ ██████╗  ██████╗ ███╗   ██╗███████╗ ██████╗ ██╗     ██╗ ██████╗ ", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s   /  /       %s%s████╗ ████║██╔═══██╗██╔═══██╗████╗  ██║██╔════╝██╔═══██╗██║     ██║██╔═══██╗", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s   |  |       %s%s██╔████╔██║██║   ██║██║   ██║██╔██╗ ██║█████╗  ██║   ██║██║     ██║██║   ██║", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s   \\  \\       %s%s██║╚██╔╝██║██║   ██║██║   ██║██║╚██╗██║██╔══╝  ██║   ██║██║     ██║██║   ██║", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s    '._'-._   %s%s██║ ╚═╝ ██║╚██████╔╝╚██████╔╝██║ ╚████║██║     ╚██████╔╝███████╗██║╚██████╔╝", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s       ```    %s%s╚═╝     ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚═╝      ╚═════╝ ╚══════╝╚═╝ ╚═════╝  ", ColorWhite, ColorRed, BoldStart),
		fmt.Sprintf("%s", BoldEnd),
	}
}
