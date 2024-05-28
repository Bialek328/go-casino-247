package boards

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	// "os"
	// "os/exec"
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

const (
	RedColor   = "\x1b[31m"
	BlueColor  = "\x1b[34m"
	GreenColor = "\x1b[32m"
	ResetColor = "\x1b[0m"
)

func showBanner(color string) {
	str := color + baner + ResetColor
	fmt.Print(str)
}

func clearBanner() {
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
	clearBanner()
	showBanner(color)
	time.Sleep(interval)
}
