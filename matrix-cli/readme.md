# Matrix-Op CLI

Matrix-Op CLI est une application de ligne de commande (CLI) légère et puissante pour effectuer des opérations sur les matrices. Elle permet de créer, additionner, et multiplier des matrices directement depuis le terminal.

---

## Fonctionnalités

- **Création de matrices** : Créez des matrices avec des dimensions et des valeurs par défaut.
- **Addition de matrices** : Additionnez deux matrices de même taille.
- **Multiplication de matrices** : Multipliez deux matrices compatibles.
- **Astuce de compatibilité** : Assurez-vous que les dimensions des matrices sont correctes avant d'effectuer des calculs.

---

## Utilisation

### Lancer l'application

Sur Linux ou Windows, exécutez simplement le fichier binaire :

- **Linux** :
  ```bash
  ./matrix-cli
  ```

- **Windows** :
  ```cmd
  matrix-cli.exe
  ```

### Menu principal

Après le lancement, vous verrez un menu interactif avec les options suivantes :
```plaintext
╔════════════════════════════════════════════════╗
║                 Matrix-Op CLI                 ║
╠════════════════════════════════════════════════╣
║ Un outil simple pour manipuler les matrices.  ║
╚════════════════════════════════════════════════╝
Bienvenue dans Matrix-Op CLI !
Astuce : Les matrices doivent avoir des dimensions compatibles pour les calculs.
Commandes disponibles :
  1. create [rows] [cols] [value] - Crée une matrice avec une valeur par défaut
  2. add [rows] [cols] - Additionne deux matrices de même taille
  3. multiply [rows1] [cols1] [cols2] - Multiplie deux matrices compatibles
  4. exit - Quitte l'application
```

### Exemples de commandes

1. **Création d'une matrice** :
   ```bash
   create 3 3 0
   ```
   Crée une matrice 3x3 remplie de zéros.

2. **Addition de matrices** :
   ```bash
   add 2 2
   ```
   Additionne deux matrices 2x2.

3. **Multiplication de matrices** :
   ```bash
   multiply 2 3 2
   ```
   Multiplie une matrice 2x3 avec une matrice 3x2.

4. **Quitter l'application** :
   ```bash
   exit
   ```

---

## Compilation

### Multiplateforme

Pour compiler le programme pour Linux et Windows, utilisez les commandes suivantes :

#### Pour Linux :
```bash
GOOS=linux GOARCH=amd64 go build -o matrix-op main.go
```

#### Pour Windows :
```bash
GOOS=windows GOARCH=amd64 go build -o matrix-op.exe main.go
```


