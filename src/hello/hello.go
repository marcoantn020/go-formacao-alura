package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 1
const delay = 5

func main() {
	showIntroduction()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		case 1:
			initialMonitoring()
		case 2:

			showLog()
		default:
			fmt.Println("Comando invalido")
			os.Exit(-1)
		}
	}

}

func showLog() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	fmt.Println("Logs...")
	fmt.Println(string(file))
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	statusSite := "online"
	if status == false {
		statusSite = "offline"
	}

	message := time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - " + statusSite + ": " + strconv.FormatBool(status) + "\n"
	_, err = file.WriteString(message)
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
}

func initialMonitoring() {
	fmt.Println("Monitorando...")

	sites := readSitesFile()

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			verifySites(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func readSitesFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	return sites
}

func verifySites(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readCommand() int {
	var command int
	_, err := fmt.Scan(&command)
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	return command
}

func showMenu() {
	fmt.Println("\n1. Iniciar Monitoramento")
	fmt.Println("2. Exibir Logs")
	fmt.Println("0. Sair do Programa")
}

func showIntroduction() {
	name := "Marco"
	version := 1.1
	fmt.Println("Ola sr. ", name)
	fmt.Println("Este programa esta na versao ", version)
}
