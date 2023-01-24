package main

import (
	"classic"
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

type game struct {
	Title      string
	Word       string
	WordUser   []string
	LengthWord int
	Attempts   int
	ToFind     []string
	Position   string
	File       string
}

type Language struct {
	Fr []string
	En []string
	Es []string
	Ge []string
}
type Settings struct {
	Language  Language
	Langue    []string
	Scorboard []string
	Pictures  []string
	Sound     []string
}
type UserInfo struct {
	Difficulty string
	Pseudo     string
	Score      int
}

type Board struct {
	Easy Facile

	Medium Moyen

	Hard Difficile
}

type Facile struct {
	Pseudo1 string
	Score1  int
	Pseudo2 string
	Score2  int
	Pseudo3 string
	Score3  int
}
type Moyen struct {
	Pseudo1 string
	Score1  int
	Pseudo2 string
	Score2  int
	Pseudo3 string
	Score3  int
}
type Difficile struct {
	Pseudo1 string
	Score1  int
	Pseudo2 string
	Score2  int
	Pseudo3 string
	Score3  int
}
type base struct {
	Hangman game
	Set     Settings
}

var bd = base{}
var Sb = Board{}
var Joueur = UserInfo{}

func variable() {
	bd.Set.Language.Fr = []string{"New Super Hangman Web", "facile", "moyen", "difficile", "entre un nom", "lancer",
		"Bonne chance ", "Vous avez", "essaies", "entrez une lettre ou un mot", "envoyé", "lettre déja essayer : ", "rejouer",
		"tu as Gagne", "tableaux des scores", "rejoué",
	}
	bd.Set.Language.En = []string{"New Super Hangman Web", "easy", "medium", "hard", "enter a nickname", "start",
		"good luck ",
	}
	bd.Set.Language.Es = []string{"New Super Hangman Web", "fácil", "medio", "difícil", "Introduce un apodo", "iniciar"}
	bd.Set.Langue = bd.Set.Language.En
	var Word = classic.RandomWord("words.txt")
	var data = game{
		Title: "...", Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""),
		LengthWord: len(Word), Position: "https://clipground.com/images/html5-logo-2.png", File: "word3.txt",
	}
	Sb.Easy = Facile{Pseudo1: "N/A", Score1: 0, Pseudo2: "N/A", Score2: 0, Pseudo3: "N/A", Score3: 0}
	Sb.Medium = Moyen{Pseudo1: "N/A", Score1: 0, Pseudo2: "N/A", Score2: 0, Pseudo3: "N/A", Score3: 0}
	Sb.Hard = Difficile{Pseudo1: "N/A", Score1: 0, Pseudo2: "N/A", Score2: 0, Pseudo3: "N/A", Score3: 0}
	bd.Hangman = data
}
func Home(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./Source/Web/" + "menuv2" + ".html")
	if r.FormValue("send") == "submit" {
		if r.FormValue("dif") == "fa" {
			bd.Hangman.File = "words.txt"
		} else if r.FormValue("dif") == "mo" {
			bd.Hangman.File = "words2.txt"
		} else if r.FormValue("dif") == "di" {
			bd.Hangman.File = "words3.txt"
		} else {
			bd.Hangman.File = "words.txt"
		}
		Joueur.Pseudo = r.FormValue("name")
		var Word = classic.RandomWord(bd.Hangman.File)
		bd.Hangman = game{Title: "goodluck " + r.FormValue("name"), Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""), LengthWord: 5, Position: "https://clipground.com/images/html5-logo-2.png"}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if r.FormValue("param") == "submit" {
		http.Redirect(w, r, "/setting", http.StatusSeeOther)
	}
	page.ExecuteTemplate(w, "menuv2.html", bd)
}
func scoreboard(User *UserInfo, Scoreboard *Board) {
	switch User.Difficulty {
	case "Hard":
		if User.Score > Scoreboard.Hard.Score1 {
			Scoreboard.Hard.Score3 = Scoreboard.Hard.Score2
			Scoreboard.Hard.Pseudo3 = Scoreboard.Hard.Pseudo2
			Scoreboard.Hard.Score2 = Scoreboard.Hard.Score1
			Scoreboard.Hard.Pseudo2 = Scoreboard.Hard.Pseudo1
			Scoreboard.Hard.Score1 = User.Score
			Scoreboard.Hard.Pseudo1 = User.Pseudo
			break
		} else if User.Score < Scoreboard.Hard.Score1 && User.Score > Scoreboard.Hard.Score2 {
			Scoreboard.Hard.Pseudo3 = Scoreboard.Hard.Pseudo2
			Scoreboard.Hard.Score3 = Scoreboard.Hard.Score2
			Scoreboard.Hard.Pseudo2 = User.Pseudo
			Scoreboard.Hard.Score2 = User.Score
			break
		} else if User.Score > Scoreboard.Hard.Score3 {
			Scoreboard.Hard.Score3 = User.Score
			Scoreboard.Hard.Pseudo3 = User.Pseudo
		} else {
			break
		}
	case "Medium":
		if User.Score > Scoreboard.Medium.Score1 {
			Scoreboard.Medium.Score3 = Scoreboard.Medium.Score2
			Scoreboard.Medium.Pseudo3 = Scoreboard.Medium.Pseudo2
			Scoreboard.Medium.Score2 = Scoreboard.Medium.Score1
			Scoreboard.Medium.Pseudo2 = Scoreboard.Medium.Pseudo1
			Scoreboard.Medium.Score1 = User.Score
			Scoreboard.Medium.Pseudo1 = User.Pseudo
			break
		} else if User.Score < Scoreboard.Medium.Score2 && User.Score > Scoreboard.Medium.Score3 {
			Scoreboard.Medium.Pseudo3 = Scoreboard.Medium.Pseudo2
			Scoreboard.Medium.Score3 = Scoreboard.Medium.Score2
			Scoreboard.Medium.Pseudo2 = User.Pseudo
			Scoreboard.Medium.Score2 = User.Score
			break
		} else if User.Score > Scoreboard.Medium.Score3 {
			Scoreboard.Medium.Score3 = User.Score
			Scoreboard.Medium.Pseudo3 = User.Pseudo
		} else {
			break
		}
	case "Easy":
		if User.Score > Scoreboard.Easy.Score1 {
			Scoreboard.Easy.Score3 = Scoreboard.Easy.Score2
			Scoreboard.Easy.Pseudo3 = Scoreboard.Easy.Pseudo2
			Scoreboard.Easy.Score2 = Scoreboard.Easy.Score1
			Scoreboard.Easy.Pseudo2 = Scoreboard.Easy.Pseudo1
			Scoreboard.Easy.Score1 = User.Score
			Scoreboard.Easy.Pseudo1 = User.Pseudo
			break
		} else if User.Score < Scoreboard.Easy.Score2 && User.Score > Scoreboard.Easy.Score3 {
			Scoreboard.Easy.Pseudo3 = Scoreboard.Easy.Pseudo2
			Scoreboard.Easy.Score3 = Scoreboard.Easy.Score2
			Scoreboard.Easy.Pseudo2 = User.Pseudo
			Scoreboard.Easy.Score2 = User.Score
			break
		} else if User.Score > Scoreboard.Easy.Score3 {
			Scoreboard.Easy.Score3 = User.Score
			Scoreboard.Easy.Pseudo3 = User.Pseudo
		} else {
			break
		}
	}
	User.Score = 0
	User.Pseudo = "N/A"
	return
}
func Hangman(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./Source/Web/" + "hangman" + ".tmpl")
	if r.FormValue("loser") == "submit" {
		http.Redirect(w, r, "/loser", http.StatusSeeOther)
	}
	if r.FormValue("reset") == "submit" {
		Word := classic.RandomWord("words.txt")
		bd.Hangman = game{
			Title: "Hangman by Léo & Nathan", Word: classic.Upper(Word), WordUser: classic.WordChoice(Word), Attempts: 10, ToFind: classic.StringToList(""),
			LengthWord: len(Word), Position: "https://clipground.com/images/html5-logo-2.png",
		}
	}
	choice := classic.Upper(r.FormValue("wordletter"))
	fmt.Println(choice)
	if len(choice) == 1 {
		index := classic.Verif(bd.Hangman.Word, choice)
		for i := 0; i < len(index); i++ {
			bd.Hangman.WordUser[index[i]] = choice
		}
		if len(index) == 0 {
			bd.Hangman.Attempts -= 1
		} else {
			bd.Hangman.Attempts += doublons(bd.Hangman.ToFind, choice)
		}
	} else {
		if choice == bd.Hangman.Word {
			bd.Hangman.WordUser = classic.StringToList("Congrats !")
			http.Redirect(w, r, "/win", http.StatusSeeOther)
		} else if choice != bd.Hangman.Word && len(choice) > 1 {
			bd.Hangman.Attempts -= 2
		}
	}
	if choice != "" {
		bd.Hangman.ToFind = append(bd.Hangman.ToFind, choice)
	}
	if (len(classic.Verif(classic.ListToString(bd.Hangman.WordUser), "_")) == 0) && (choice != bd.Hangman.Word) {
		http.Redirect(w, r, "/win", http.StatusSeeOther)
	}
	if bd.Hangman.Attempts <= 0 {
		http.Redirect(w, r, "/loser", http.StatusSeeOther)
	}
	t.ExecuteTemplate(w, "hangman.tmpl", bd)
}
func Loser(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./Source/Web/" + "loser" + ".html")
	if r.FormValue("restart") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	page.ExecuteTemplate(w, "loser.html", bd)
}
func Win(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./Source/Web/" + "win" + ".html")
	if r.FormValue("restart") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	page.ExecuteTemplate(w, "win.html", bd)
}
func Scoreb(w http.ResponseWriter, r *http.Request) {
	start, _ := template.ParseFiles("./Web/" + "ScoreBoard.page.tmpl" + ".html")
	start.ExecuteTemplate(w, "ScoreBoard.page.tmpl.html", Sb)
}

func Parameter(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./Source/Web/" + "param" + ".html")
	if r.FormValue("lg") == "en" {
		bd.Set.Langue = bd.Set.Language.En
	}
	if r.FormValue("lg") == "fr" {
		bd.Set.Langue = bd.Set.Language.Fr
	}
	if r.FormValue("lg") == "es" {
		bd.Set.Langue = bd.Set.Language.Es
	}
	if r.FormValue("lg") == "ge" {
		bd.Set.Langue = bd.Set.Language.Ge
	}
	if r.FormValue("send") == "submit" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	page.ExecuteTemplate(w, "param.html", bd)
}
func doublons(liste []string, choice string) int {
	for i := 0; i < len(liste); i++ {
		if liste[i] == choice {
			return -1
		}
	}
	return 0
}
func main() {
	variable()
	http.HandleFunc("/home", Home)
	http.HandleFunc("/setting", Parameter)
	http.HandleFunc("/loser", Loser)
	http.HandleFunc("/win", Win)
	http.HandleFunc("/scoreboard", Scoreb)
	http.HandleFunc("/", Hangman)
	fmt.Println("http://localhost" + port + "/home")
	fs := http.FileServer(http.Dir("Source"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe("localhost:8080", nil)
}
