package main


import ("fmt" 
"io/ioutil"
"os"
"bufio"
)

func main(){
    boucle()
}


func boucle(){
    var choix int
    var tmp = true
    for tmp {
        fmt.Println("1. Lire le fichier texte.txt")
        fmt.Println("2. Ecrire dans le fichier texte.txt")
        fmt.Println("3. Vider le fichier texte.txt")
        fmt.Println("4. Ajouter du texte dans le fichier texte.txt")
        fmt.Println("5. Quitter")
        fmt.Print("Votre choix : ")
        fmt.Scanln(&choix)
        switch choix {
        case 1:
            lire()
        case 2:
            ecrire()
        case 3:
            vider()
        case 4:
            ajouter_texte()
        case 5:
            tmp = false
        default:
            fmt.Println("Choix invalide")
            fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n")
            main()
        }
    }
}





func lire(){
	data, err := ioutil.ReadFile("texte.txt")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Contenu du fichier texte.txt : ")
    fmt.Println("----------------------------")
    fmt.Println(string(data))
    fmt.Println("----------------------------")
}

func ecrire(){
    file, err := os.OpenFile("texte.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(os.Stdin) 
    fmt.Print("Entrez quelque chose : ")
    scanner.Scan()                      
    entreeUtilisateur := scanner.Text() 
    data := []byte(entreeUtilisateur)
    errr := ioutil.WriteFile("texte.txt", data, 0644)
    if errr != nil {
        fmt.Println(errr)
    }
}

func vider(){
    file, err := os.OpenFile("texte.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
}

func ajouter_texte(){
    file, err := os.OpenFile("texte.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(os.Stdin) 
    fmt.Print("Entrez quelque chose : ")
    scanner.Scan()                      
    entreeUtilisateur := scanner.Text() 
    file, errr := os.OpenFile("texte.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
    if errr != nil {
        fmt.Println(errr)
    }
    defer file.Close()

    if err != nil {
        panic(err)
    }

    _, err = file.WriteString(entreeUtilisateur)
    if err != nil {
        panic(err)
    }

}