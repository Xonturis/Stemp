Contexte
========

Ce programme a été créé dans le cadre de mon Projet de Recherche et Développement durant mon cursus ingénieur à l'école Polytechnique de l'université de Nantes. Il s'agit d'un programme qui récupère la température du système parle biais du fichier spécial `/sys/class/thermal/thermal_zoneX/temp`.

En outre, il est possible de récolter la consommation énergétique du CPU avec quelques modifications mineures, le code qui permet cette récolte est disponible dans le répertoire `rapl`.

Compilation
===========

Afin de compiler ce programme il faut installer [Go](https://go.dev/doc/install).

En suite, pour compiler il suffit de lancer cette commande :
```
go build
```

Utilisation
===========

Après avoir compilé, un fichier `main` est créé qui peut être directement lancé depuis la console. La sonde une fois lancée va créer un fichier csv. Ce fichier csv nommé suivant le format `ddMMYYYY-Stemp.csv` comporte deux colonnes : `timestamp` et `temperature`.

La sonde récolte et écrit les données relatives à la température du système à un rythme régulier de 10 secondes. Ce rythmes est actuellement constant et ne peut être changé que par le biais du code dans la fonction `main` du fichier `main.go`.

Les données collectées sont, pour le timestamp, au format Unix, c'est-à-dire le nombre de secondes depuis le 1er janvier 1970, la température est en °C.

Licence
=======

Ce code est sous licence BSD (3-Clause BSD License). Voir le fichier `LICENCE`.