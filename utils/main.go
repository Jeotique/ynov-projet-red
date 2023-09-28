package utils

import (
	"bufio"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/mattn/go-tty"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func ListenInput() string {
	println("listening for input")
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		println("aaa")
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		return string(r)
	}
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func Write(text string, args ...any) {
	if args != nil && len(args) > 0 {
		fmt.Printf(text, args)
	} else {
		print(text)
	}
}

func Writeln(text string, args ...any) {
	if args != nil && len(args) > 0 {
		fmt.Printf(text, args)
		println()
	} else {
		println(text)
	}
}

func Writeanim(text string, timee time.Duration, args ...any) {
	for _, character := range text {
		print(string(rune(character)))
		time.Sleep(timee * time.Millisecond)
	}
}

func WaitForInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

func RandomNumber(min int, max int) int {
	return rand.Intn((max - min + 1) + min)
}

func WaitForNumberInput() (bool, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	if !IsNumeric(scanner.Text()) {
		return false, -1
	}
	number, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
		return false, -1
	}
	return true, number
}

func IsNumeric(s string) bool {
	for _, char := range s {
		if !('0' <= char && char <= '9') {
			return false
		}
	}
	return true
}

func PlaySound(path string, rate ...int) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	if len(rate) < 1 {
		rate = []int{1}
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	//speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Init(format.SampleRate*beep.SampleRate(rate[0]), format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
