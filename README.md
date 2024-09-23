# Hangman Game (Jeu du Pendu)

## Description

Ce projet implémente le célèbre jeu du pendu ("Hangman") en GoLang. Le but du jeu est de deviner un mot aléatoire en proposant des lettres, avec un nombre limité de tentatives. Si le joueur propose une lettre incorrecte, un personnage commence à être dessiné, représentant la progression vers la défaite. Le jeu se termine lorsque le mot est découvert ou lorsque le joueur a épuisé toutes ses tentatives.

## Fonctionnalités

> Le programme choisit un mot au hasard depuis un fichier words.txt.
> Certaines lettres du mot sont révélées aléatoirement pour aider le joueur.
> Le joueur dispose de 10 tentatives pour deviner le mot complet.
> À chaque mauvaise proposition, une partie du dessin du pendu est affichée.
> Le jeu s'exécute dans la console, avec des interactions via des entrées clavier.