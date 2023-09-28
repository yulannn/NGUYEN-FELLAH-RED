package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func ClearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin", "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (p *Personnage) DisplayInfo() {
	fmt.Printf("Nom: %s\n", p.nom)
	fmt.Printf("Class: %s\n", p.class)
	fmt.Printf("Niveau: %d\n", p.level)
	fmt.Printf("Point_vie_max: %d\n", p.pointVieMax)
	fmt.Printf("Point_vie_actuel: %d\n", p.pointVieActual)
	fmt.Printf("Coins: %d\n", p.coins)
	fmt.Printf("------------------ \n")

	fmt.Println("\n")
}

func Inputint() (int, error) {
	fmt.Print(">> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	chiffre, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return chiffre, nil
}
