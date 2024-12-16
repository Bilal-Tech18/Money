# Spid3r - Website Crawler

**Spid3r** est un outil puissant conçu pour crawler les sites web, en extrayant chaque page, champ de saisie (`input`, `textarea`), et bouton (`button`, `a`, `div`, `span` avec des rôles interactifs). Il offre des options de sortie flexibles, permettant de sauvegarder les résultats aux formats Excel, texte ou JSON, ce qui le rend idéal pour diverses applications telles que la collecte de données et l'analyse de sites web.

## Installation

Avant d'utiliser Spid3r, assurez-vous que toutes les bibliothèques requises sont installées en exécutant le script d'installation :

```
./Spid3r_installer.sh
```
## Compilation de l'outil

Une fois les dépendances installées, compilez l'outil en utilisant la commande suivante :

```
go build crawling.go createExcelFile.go createTxtFile.go saveJson.go main.go
```
## Utilisation

Vous pouvez exécuter Spid3r directement avec la commande suivante :

```
./crawling [--excel|--txt|--json|--verbose|--help] <url>
```
### Options :

--excel : Sauvegarde la sortie dans un fichier Excel.
--txt : Sauvegarde la sortie dans un fichier texte.
--json : Sauvegarde la sortie dans un fichier JSON.
--verbose : Affiche chaque URL visitée dans le terminal ainsi que les informations sur les éléments trouvés (inputs, boutons, etc.).
--help : Affiche les informations d'utilisation.

### Fonctionnalités

Détection d'éléments de saisie : Identifie tous les éléments input et textarea.
Détection des boutons :
    Repère tous les boutons HTML (<button>).
    Identifie les liens <a> avec role="button".
    Détecte les div et span ayant role="button" ou des attributs d'accessibilité comme aria-pressed ou aria-haspopup.
Options de sortie : Enregistre les résultats sous forme de fichier Excel, texte ou JSON, permettant une analyse approfondie.

Exemple

```
./crawling --excel "https://example.com"
```

Cette commande crawl le site web et sauvegarde les résultats dans un fichier Excel.




---

Ce fichier `README.md` fournit une explication claire et détaillée des fonctionnalités, de l’installation, de la compilation et de l’utilisation de l'outil **Spid3r**.

