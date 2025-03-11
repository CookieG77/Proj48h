# Utiliser une image de base Go compatible
FROM golang:1.24 AS builder

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers go.mod et go.sum
COPY Projet/go.mod Projet/go.sum ./

# Télécharger les dépendances
RUN go mod download

# Copier le reste des fichiers du projet
COPY Projet/ ./

# Compiler l'application
RUN go build -o main .

# Étape finale
FROM alpine:latest

# Installer les dépendances nécessaires pour les binaires Go
RUN apk --no-cache add ca-certificates libc6-compat

# Définir le répertoire de travail
WORKDIR /app

# Copier l'exécutable depuis l'étape de construction
COPY --from=builder /app/main .

# Copier le fichier de configuration
COPY --from=builder /app/MailConfig.json .

# Copier les ressources statiques et templates
COPY --from=builder /app/statics ./statics
COPY --from=builder /app/templates ./templates

# Assurer que l'exécutable a les bonnes permissions
RUN chmod +x ./main

# Exposer le port
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["./main"]
