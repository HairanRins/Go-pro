package main 

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"matrix-cli/matrix"
	
)

func main() {
	fmt.Println("\033[1;32m")
	fmt.Println("╔═════════════════════════════════════════════╗")
	fmt.Println("║                 Matrix-Op CLI               ║")
	fmt.Println("╠═════════════════════════════════════════════╣")
	fmt.Println("║ Un outil pour manipuler les matrices        ║")
	fmt.Println("╚═════════════════════════════════════════════╝")
	fmt.Println("Astuce : Les matrices doivent avoir des dimensions compatibles pour les calculs.")
	fmt.Println("Commandes disponibles :")
	fmt.Println("  1. create [rows] [cols] [value] - Crée une matrice avec une valeur par défaut")
	fmt.Println("  2. add [rows] [cols] - Additionne deux matrices de même taille")
	fmt.Println("  3. multiply [rows1] [cols1] [cols2] - Multiplie deux matrices compatibles")
	fmt.Println("  4. exit - Quitte l'application")
	fmt.Println("---------------------------------------------------")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Lire la commande utilisateur
		fmt.Print("Entrez une commande : ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "create":
			if len(args) != 4 {
				fmt.Println("Usage: create [rows] [cols] [value]")
				continue
			}
			rows, _ := strconv.Atoi(args[1])
			cols, _ := strconv.Atoi(args[1])
			value, _ := strconv.Atoi(args[1])
			mat := matrix.New(rows, cols, value)
			fmt.Println("Matrice créée :")
			mat.Print()

		case "add":
			if len(args) != 3 {
				fmt.Println("Usage: add [rows] [cols]")
				continue
			}
			rows, _ := strconv.Atoi(args[1])
			cols, _ := strconv.Atoi(args[2])
			fmt.Println("Entrez les valeurs pour la première matrice :")
			mat1 := readMatrix(rows, cols, scanner)
			fmt.Println("Entrez les valeurs pour la deuxième matrice :")
			mat2 := readMatrix(rows, cols, scanner)

			result, err := matrix.Add(mat1, mat2)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Résultat de l'addition :")
				result.Print()
			}

		case "multiply":
			if len(args) != 4 {
				fmt.Println("Usage: multiply [rows1] [cols1] [cols2]")
				continue
			}
			rows1, _ := strconv.Atoi(args[1])
			cols1, _ := strconv.Atoi(args[2])
			cols2, _ := strconv.Atoi(args[3])
			fmt.Println("Entrez les valeurs pour la première matrice :")
			mat1 := readMatrix(rows1, cols1, scanner)
			fmt.Println("Entrez les valeurs pour la deuxième matrice :")
			mat2 := readMatrix(cols1, cols2, scanner)

			result, err := matrix.Multiply(mat1, mat2)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Résultat de la multiplication :")
				result.Print()
			}
		
		case "exit":
			fmt.Println("Au revoir!")
			return

		default:
			fmt.Println("Commande inconnue. Veuillez réessayer.")
		
		}
	}
}

// Fonction pour lire une matrice à partir de l'entrée utilisateur
func readMatrix(rows, cols int, scanner *bufio.Scanner)  matrix.Matrix {
	data := make([][]int, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]int, cols)
		fmt.Printf("Ligne %d (entrez %d valeurs séparées par des espaces) : ", i+1, cols)
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		for j := 0; j < cols; j++ {
			data[i][j], _ = strconv.Atoi(line[j])
		}
	}
	mat, _ := matrix.CreateWith(data)
	return mat
}