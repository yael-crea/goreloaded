# Compteur de mots

Ce projet est une application en ligne de commande écrite en Go. Elle lit un fichier texte et affiche le nombre de mots qu'il contient.

Un mot est ici une suite de caractères séparée par un ou plusieurs espaces, tabulations ou retours à la ligne. La ponctuation accolée à un mot ne crée pas un mot supplémentaire.

## Prérequis

- Go 1.22 ou une version plus récente compatible ;
- un fichier texte à analyser.

Vérifier l'installation de Go :

```bash
go version
```

## Préparation

Se placer dans le dossier du projet :

```bash
cd e1-compter-les-mots-m1-x20-2
```

Aucune dépendance externe n'est nécessaire.

## Utilisation

Lancer le programme en lui donnant exactement un fichier :

```bash
go run . in.txt
```

Avec le fichier fourni, la sortie attendue est :

```text
7
```

Le programme écrit un message explicite sur la sortie d'erreur et retourne le code `1` si :

- l'argument est absent ;
- plusieurs arguments sont fournis ;
- le fichier n'existe pas ou ne peut pas être ouvert ;
- une erreur survient pendant la lecture.

## Fonctionnalités

- lecture d'un fichier texte ;
- comptage des mots séparés par des espaces Unicode ;
- prise en charge des tabulations et retours à la ligne ;
- prise en charge d'un mot pouvant atteindre 1 Mio ;
- messages d'erreur compréhensibles ;
- code de sortie `0` en cas de succès et `1` en cas d'erreur.

Le projet conserve aussi plusieurs fonctions d'exercice testées séparément dans `exercises.go` : conversion vers le décimal, capitalisation, normalisation de ponctuation et passage en majuscules des derniers mots.

## Organisation

- `main.go` : comptage, orchestration de la commande et point d'entrée ;
- `main_test.go` : tests du comptage et de l'application CLI ;
- `exercises.go` : fonctions d'exercice complémentaires ;
- `exercises_test.go` : tests de ces fonctions ;
- `in.txt` : exemple d'entrée ;
- `SUIVI_SEANCE_3.md` : matrice de tests, revue qualité et trame de démonstration.

## Tests et outils de qualité

Exécuter tous les tests :

```bash
go test ./...
```

Afficher le détail des tests :

```bash
go test -v ./...
```

Formater et analyser le code :

```bash
gofmt -w .
go vet ./...
```

Les tests couvrent les cas nominaux, les cas limites, les erreurs et l'exécution réelle de la commande.

## Choix techniques

- `bufio.Scanner` avec `bufio.ScanWords` fournit une règle de découpage simple et adaptée au besoin.
- La fonction `countWords` reçoit un `io.Reader` : elle peut donc être testée sans créer de fichier.
- La fonction `run` reçoit les arguments et les sorties : les erreurs et codes de sortie sont vérifiables directement dans les tests.
- `main` reste volontairement court et délègue le travail à `run`.
- La taille maximale d'un mot est portée à 1 Mio pour éviter la limite par défaut de 64 Kio du scanner.

## Limite connue

Le programme compte les mots selon les espaces. Ainsi, `bonjour,monde` compte comme un seul mot, car aucune espace ne sépare les deux parties.


