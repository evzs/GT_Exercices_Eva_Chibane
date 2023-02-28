package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	nbAllumettes := 77
	joueurs := []string{"Joueur 1", "Joueur 2"}
	joueur := 0
	scanner := bufio.NewScanner(os.Stdin)

	for nbAllumettes > 0 {
		fmt.Println("Il reste", nbAllumettes, "allumettes")
		fmt.Println(joueurs[joueur], "combien d'allumettes retirez-vous ?")
		scanner.Scan()
		allumettesRetirees, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Veuillez saisir un nombre valide")
			continue
		}

		if allumettesRetirees < 1 || allumettesRetirees > 3 {
			fmt.Println("Vous devez retirer entre 1 et 3 allumettes")
		} else if allumettesRetirees > nbAllumettes {
			fmt.Println("Vous ne pouvez pas retirer plus d'allumettes que celles restantes")
		} else {
			nbAllumettes -= allumettesRetirees
			if nbAllumettes == 0 {
				fmt.Println(joueurs[joueur], "a perdu")
			} else {
				joueur = 1 - joueur
			}
		}
	}
}
