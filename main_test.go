package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "phrase simple", input: "un deux trois quatre", want: 4},
		{name: "plusieurs lignes", input: "un deux\ntrois\tquatre", want: 4},
		{name: "chaîne vide", input: "", want: 0},
		{name: "espaces seulement", input: "  \n\t ", want: 0},
		{name: "un seul mot", input: "bonjour", want: 1},
		{name: "ponctuation", input: "Bonjour, tout le monde !", want: 5},
		{name: "apostrophe", input: "Aujourd'hui l'été arrive", want: 3},
		{name: "caractères Unicode", input: "été déjà là", want: 3},
		{name: "mot supérieur à 64 Kio", input: strings.Repeat("a", 70*1024), want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := countWords(strings.NewReader(test.input))
			if err != nil {
				t.Fatalf("countWords() retourne une erreur inattendue : %v", err)
			}
			if got != test.want {
				t.Fatalf("countWords(%q) = %d, attendu %d", test.input, got, test.want)
			}
		})
	}
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) {
	return 0, errors.New("lecture impossible")
}

func TestCountWordsReadError(t *testing.T) {
	if _, err := countWords(failingReader{}); err == nil {
		t.Fatal("countWords() devait transmettre l'erreur de lecture")
	}
}

func TestRun(t *testing.T) {
	tmp := t.TempDir()
	input := filepath.Join(tmp, "in.txt")
	if err := os.WriteFile(input, []byte("un deux trois"), 0o600); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr string
	}{
		{
			name:       "fichier valide",
			args:       []string{input},
			wantCode:   0,
			wantStdout: "3\n",
		},
		{
			name:       "argument manquant",
			args:       nil,
			wantCode:   1,
			wantStderr: "usage:",
		},
		{
			name:       "trop d'arguments",
			args:       []string{input, "autre.txt"},
			wantCode:   1,
			wantStderr: "usage:",
		},
		{
			name:       "fichier absent",
			args:       []string{filepath.Join(tmp, "absent.txt")},
			wantCode:   1,
			wantStderr: "impossible d'ouvrir",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			gotCode := run(test.args, &stdout, &stderr)
			if gotCode != test.wantCode {
				t.Fatalf("run() = %d, attendu %d", gotCode, test.wantCode)
			}
			if got := stdout.String(); got != test.wantStdout {
				t.Fatalf("stdout = %q, attendu %q", got, test.wantStdout)
			}
			if !strings.Contains(stderr.String(), test.wantStderr) {
				t.Fatalf("stderr = %q, doit contenir %q", stderr.String(), test.wantStderr)
			}
		})
	}
}

func TestCommand(t *testing.T) {
	tmp := t.TempDir()
	binary := filepath.Join(tmp, "compter-mots")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	build := exec.Command("go", "build", "-o", binary, ".")
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, output)
	}

	input := filepath.Join(tmp, "in.txt")
	if err := os.WriteFile(input, []byte("un deux\n\ttrois   quatre\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	t.Run("compte les mots", func(t *testing.T) {
		cmd := exec.Command(binary, input)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("commande: %v\n%s", err, output)
		}
		if got := strings.TrimSpace(string(output)); got != "4" {
			t.Fatalf("résultat = %q, attendu 4", got)
		}
	})

	t.Run("mauvais nombre d'arguments", func(t *testing.T) {
		cmd := exec.Command(binary)
		if err := cmd.Run(); err == nil {
			t.Fatal("la commande devait échouer")
		} else if exitErr, ok := err.(*exec.ExitError); !ok || exitErr.ExitCode() != 1 {
			t.Fatalf("code de sortie = %v, attendu 1", err)
		}
	})

	t.Run("fichier absent", func(t *testing.T) {
		cmd := exec.Command(binary, filepath.Join(tmp, "absent.txt"))
		if err := cmd.Run(); err == nil {
			t.Fatal("la commande devait échouer")
		} else if exitErr, ok := err.(*exec.ExitError); !ok || exitErr.ExitCode() != 1 {
			t.Fatalf("code de sortie = %v, attendu 1", err)
		}
	})
}

