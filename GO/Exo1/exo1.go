package main

import "fmt"

func main() {
    allumettes := 13
    maxPrises := 3
    joueurActuel := 1

    for allumettes > 0 {
        fmt.Printf("Il reste %d allumettes\n", allumettes)

        fmt.Printf("Joueur %d, combien d'allumettes voulez-vous prendre ? ", joueurActuel)
        var prises int
        fmt.Scanln(&prises)

        if prises < 1 || prises > maxPrises || prises > allumettes {
            fmt.Println("Nombre d'allumettes pris invalide !")
            continue
        }

        allumettes -= prises

        if joueurActuel == 1 {
            joueurActuel = 2
        } else {
            joueurActuel = 1
        }
    }

    fmt.Printf("Le joueur %d a gagné !\n", joueurActuel)
}