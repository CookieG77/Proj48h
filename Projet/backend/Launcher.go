package backend

import (
	pages "Proj48h/backend/pages"
	"Proj48h/functions"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"strconv"
)

var Port = "8080"

// LaunchWebApp lance l'application web
func LaunchWebApp() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				functions.ClearCmd()
				os.Exit(1)
			}
		}
	}()

	// Gestion des arguments
	Args := functions.GetArgs()
	if slices.Contains(Args, "--port") {
		// Gestion argument '--port [port]'
		if len(Args) > slices.Index(Args, "--port")+1 {
			rawProposedPort := Args[slices.Index(Args, "--port")+1]
			proposedPort, err := strconv.Atoi(rawProposedPort)

			if err != nil || proposedPort < 1 || proposedPort > 65535 {
				fmt.Println("Le port proposé n'est pas un nombre valide.")
				os.Exit(1)
			}
			Port = rawProposedPort
		}
	}

	// Mettre la gestion des fichiers statiques ici
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./statics/css"))))
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./statics/img"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./statics/js"))))

	// Mettre la gestion des routes ici
	http.HandleFunc("/", pages.HomePage)
	http.HandleFunc("/group", pages.RecapPage)

	// Définit le port du serveur local
	finalPort := ""
	if Port != "" {
		finalPort = ":" + Port
	} else {
		finalPort = ":8080"
	}
	println("Serveur lancer sur : http://localhost" + finalPort)

	// Lance le serveur local
	if err := http.ListenAndServe(finalPort, nil); err != nil {
		panic(err)
	}
}
