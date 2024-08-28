package boards

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const baner = `
                              ______   ______   ______                          
                             |____  | |____  | |____  |                         
                                 / /      / /      / /                          
                                / /      / /      / /                           
                               / /      / /      / /                            
                              /_/      /_/      /_/                             				  
 ______ ______ ______ ______ ______ ______ ______ ______ ______ ______ ______ 
|______|______|______|______|______|______|______|______|______|______|______|
 ______ ______ ______ ______ ______ ______ ______ ______ ______ ______ ______ 
|______|______|______|______|______|______|______|______|______|______|______|

   _____   .______    _______ .__   __.     ___     _  _         __ ______  
 /  __  \  |   _  \  |   ____||  \ |  |    |__ \   | || |      /  /____  | 
|  |  |  | |  |_)  | |  |__   |   \|  |       ) | | || |_     /  /    / /  
|  |  |  | |   ___/  |   __|  |  .    |      / /  |__   _|   /  /    / /   
|  '--'  | |  |      |  |____ |  |\   |     / /_     | |    /  /    / /    
 \______/  | _|      |_______||__| \__|    |____|    |_|   /__/    /_/

`
const instruction = `
                            Press < any key > to continue
`

const (
	RedColor   = "\x1b[31m"
	BlueColor  = "\x1b[34m"
	GreenColor = "\x1b[32m"
	ResetColor = "\x1b[0m"
)

func showBanner(color string) {
	str := color + baner + ResetColor
	fmt.Print(str)
	str2 := ResetColor + "\n" + instruction
	fmt.Print(str2)
}

func ClearBanner() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print("\033[H")
}

func DisplayBanner() {
	red := RedColor
	green := GreenColor
	blue := BlueColor
	interval := 500 * time.Millisecond

	blinkBanner(red, interval)
	blinkBanner(green, interval)
	blinkBanner(blue, interval)
}

func blinkBanner(color string, interval time.Duration) {
	showBanner(color)
	time.Sleep(interval)
	ClearBanner()

}
