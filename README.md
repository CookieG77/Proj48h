# Proj48h

Proj48h est une application web développée en Go, conçue pour analyser et sécuriser des sites web. Elle offre une interface utilisateur intuitive et prend en charge plusieurs langues.

## Table des matières

- [Fonctionnalités](#fonctionnalités)
- [Technologies Utilisées](#technologies-utilisées)
- [Installation](#installation)
- [Configuration](#configuration)
- [Utilisation](#utilisation)
- [Contribuer](#contribuer)
- [Licence](#licence)

## Fonctionnalités

- **Analyse de Sites Web** : L'application permet d'analyser des sites web pour détecter des vulnérabilités.
- **Envoi de Rapports** : Les utilisateurs peuvent recevoir des rapports par email après l'analyse.
- **Interface Utilisateur** : Une interface responsive avec des thèmes clairs et sombres.
- **Support Multilingue** : L'application prend en charge plusieurs langues, notamment le français et l'anglais.

## Technologies Utilisées

- **Langage de Programmation** : Go (Golang)
- **Frameworks** : Utilisation de bibliothèques comme `chromedp` pour l'analyse de sites.
- **Docker** : Conteneurisation de l'application pour faciliter le déploiement et la gestion des dépendances.

## Installation

1. Clonez le dépôt :

   ```bash
   git clone https://github.com/votre-utilisateur/Proj48h.git
   cd Proj48h
   ```

2. Installez les dépendances avec `pnpm` :

   ```bash
   pnpm install
   ```

3. Créez le répertoire temporaire :

   ```bash
   go run Projet/functions/TempDir.go
   ```

## Configuration

Avant de lancer l'application, assurez-vous de configurer le fichier `MailConfig.json` pour les paramètres d'envoi d'email.

## Utilisation

Pour démarrer l'application, exécutez la commande suivante :

```bash
   docker-compose up --build
   ```

L'application sera accessible à l'adresse suivante : [http://localhost:8080](http://localhost:8080).

## Contribuer

Les contributions sont les bienvenues ! Si vous souhaitez contribuer, veuillez suivre ces étapes :

1. Forkez le projet.
2. Créez une nouvelle branche (`git checkout -b feature/YourFeature`).
3. Apportez vos modifications et validez (`git commit -m 'Add some feature'`).
4. Poussez vos modifications (`git push origin feature/YourFeature`).
5. Ouvrez une Pull Request.

## Licence

Ce projet est sous licence MIT. Consultez le fichier [LICENSE](LICENSE) pour plus de détails.
