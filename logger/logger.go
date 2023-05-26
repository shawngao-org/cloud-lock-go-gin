package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	Black   = "\033[1;30m"
	Red     = "\033[1;31m"
	Green   = "\033[1;32m"
	Yellow  = "\033[1;33m"
	Blue    = "\033[1;34m"
	Magenta = "\033[1;35m"
	Cyan    = "\033[1;36m"
	White   = "\033[1;37m"
	Reset   = "\033[0m"
)

func LogErr(pack string, format string, a ...any) {
	logTypeImpl("error", pack, format, a...)
}

func LogWarn(pack string, format string, a ...any) {
	logTypeImpl("warn", pack, format, a...)
}

func LogInfo(pack string, format string, a ...any) {
	logTypeImpl("info", pack, format, a...)
}

func LogSuccess(pack string, format string, a ...any) {
	logTypeImpl("success", pack, format, a...)
}

func getNowTimeString() string {
	return time.Now().String()[0:19]
}

func LogStarted() {
	fmt.Println("                                              _ .-') _                                          .-. .-')   ")
	fmt.Println("                                             ( (  OO) )                                         \\  ( OO )  ")
	fmt.Println("   .-----. ,--.      .-'),-----.  ,--. ,--.   \\     .'_        ,--.      .-'),-----.    .-----. ,--. ,--.  ")
	fmt.Println("  '  .--./ |  |.-') ( OO'  .-.  ' |  | |  |   ,`'--..._)       |  |.-') ( OO'  .-.  '  '  .--./ |  .'   /  ")
	fmt.Println("  |  |('-. |  | OO )/   |  | |  | |  | | .-') |  |  \\  '       |  | OO )/   |  | |  |  |  |('-. |      /,  ")
	fmt.Println(" /_) |OO  )|  |`-' |\\_) |  |\\|  | |  |_|( OO )|  |   ' |       |  |`-' |\\_) |  |\\|  | /_) |OO  )|     ' _) ")
	fmt.Println(" ||  |`-'|(|  '---.'  \\ |  | |  | |  | | `-' /|  |   / :      (|  '---.'  \\ |  | |  | ||  |`-'| |  .   \\   ")
	fmt.Println("(_'  '--'\\ |      |    `'  '-'  '('  '-'(_.-' |  '--'  /       |      |    `'  '-'  '(_'  '--'\\ |  |\\   \\  ")
	fmt.Println("   `-----' `------'      `-----'   `-----'    `-------'        `------'      `-----'    `-----' `--' '--'  ")
}

func LogRestarting() {
	fmt.Println(" _  .-')     ('-.    .-')    .-') _      ('-.     _  .-')   .-') _               .-') _")
	fmt.Println("( \\( -O )  _(  OO)  ( OO ). (  OO) )    ( OO ).-.( \\( -O ) (  OO) )             ( OO ) )")
	fmt.Println(" ,------. (,------.(_)---\\_)/     '._   / . --. / ,------. /     '._ ,-.-') ,--./ ,--,'  ,----.")
	fmt.Println(" |   /`. ' |  .---'/    _ | |'--...__)  | \\-.  \\  |   /`. '|'--...__)|  |OO)|   \\ |  |\\ '  .-./-')")
	fmt.Println(" |  /  | | |  |    \\  :` `. '--.  .--'.-'-'  |  | |  /  | |'--.  .--'|  |  \\|    \\|  | )|  |_( O- )")
	fmt.Println(" |  |_.' |(|  '--.  '..`''.)   |  |    \\| |_.'  | |  |_.' |   |  |   |  |(_/|  .     |/ |  | .--, \\")
	fmt.Println(" |  .  '.' |  .--' .-._)   \\   |  |     |  .-.  | |  .  '.'   |  |  ,|  |_.'|  |\\    | (|  | '. (_/")
	fmt.Println(" |  |\\  \\  |  `---.\\       /   |  |     |  | |  | |  |\\  \\    |  | (_|  |   |  | \\   |  |  '--'  |.-..-..-.")
	fmt.Println(" `--' '--' `------' `-----'    `--'     `--' `--' `--' '--'   `--'   `--'   `--'  `--'   `------' `-'`-'`-'")
}

func logTypeImpl(typeStr string, pack string, format string, a ...any) {
	color := Blue
	tip := "INFO"
	switch typeStr {
	case "error":
		color = Red
		tip = "ERROR"
		break
	case "warn":
		color = Yellow
		tip = "WARN"
		break
	case "success":
		color = Green
		break
	}
	logImpl(color+"["+getNowTimeString()+"] ["+tip+"]: ["+pack+"] "+format+"\n"+Reset, a...)
}

func logImpl(format string, a ...any) {
	_, err := fmt.Fprintf(os.Stdout, format, a...)
	if err != nil {
		return
	}
}
