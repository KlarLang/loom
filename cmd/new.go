package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type BackendInfo struct {
	Name        string
	Command     string
	VersionFlag string
	Installed   bool
	Version     string
}

/*
	func detectBackends(log Log) map[string]BackendInfo {
		backends := map[string]BackendInfo{
			"java": {
				Name:        "Java",
				Command:     "java",
				VersionFlag: "-version",
			},
			"c": {
				Name:        "C (GCC)",
				Command:     "gcc",
				VersionFlag: "--version",
			},
			"python": {
				Name:        "Python",
				Command:     "python3",
				VersionFlag: "--version",
			},
			"rust": {
				Name:        "Rust",
				Command:     "rustc",
				VersionFlag: "--version",
			},
		}

		fmt.Printf("%s◉%s Detecting installed backends...\n", log.PRIMARY_COLOR, log.RESET_COLOR)

		for key, backend := range backends {
			cmd := exec.Command(backend.Command, backend.VersionFlag)
			output, err := cmd.CombinedOutput()

			if err == nil {
				backend.Installed = true
				// Pega a primeira linha da versão
				lines := strings.Split(string(output), "\n")
				if len(lines) > 0 {
					backend.Version = strings.TrimSpace(lines[0])
				}
				fmt.Printf("  %s✔%s %s detected: %s\n",
					log.SUCESS_COLOR, log.RESET_COLOR, backend.Name, backend.Version)
			} else {
				backend.Installed = false
				fmt.Printf("  %s✖%s %s not found\n",
					log.ERROR_COLOR, log.RESET_COLOR, backend.Name)
			}

			backends[key] = backend
		}

		return backends
	}

	func selectDefaultBackend(backends map[string]BackendInfo, log Log) string {
		installedBackends := []string{}

		for key, backend := range backends {
			if backend.Installed {
				installedBackends = append(installedBackends, key)
			}
		}

		if len(installedBackends) == 0 {
			fmt.Printf("\n%s⚠ Warning:%s No backends detected. Using Java as default.\n",
			log.WARNING_COLOR, log.RESET_COLOR)
			return "java"
		}

		if len(installedBackends) == 1 {
			fmt.Printf("\n%s◉%s Only one backend found, using %s as default.\n",
			log.PRIMARY_COLOR, log.RESET_COLOR, installedBackends[0])
			return installedBackends[0]
		}

		fmt.Printf("\n%s◉%s Multiple backends available. Choose default:\n",
		log.PRIMARY_COLOR, log.RESET_COLOR)

		for i, backend := range installedBackends {
			fmt.Printf("  %s[%d]%s %s\n", log.ACCENT_COLOR, i+1, log.RESET_COLOR, backends[backend].Name)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nSelect [1]: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			return installedBackends[0]
		}

		var selection int
		fmt.Sscanf(input, "%d", &selection)

		if selection > 0 && selection <= len(installedBackends) {
			return installedBackends[selection-1]
		}

		return installedBackends[0]
	}

	func generateBackendConfig(backends map[string]BackendInfo, defaultBackend string) string {
		config := "[targets]\n"

		order := []string{"java", "c", "cpp", "python", "rust", "go", "node"}

		for _, key := range order {
			backend, exists := backends[key]
			if !exists {
				continue
			}

			padding := strings.Repeat(" ", 7-len(key))
			config += fmt.Sprintf("%s%s = %v\n", key, padding, backend.Installed)
		}

		return config
	}
*/
func newCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: loom new <project_name>")
		return
	}

	log := NewLog()
	name := os.Args[2]

	log.Header()
	log.finalizeBottomheader()
	fmt.Println()
	log.Line()

	// backends := detectBackends(log)
	fmt.Println()

	// defaultBackend := selectDefaultBackend(backends, log)
	// fmt.Printf("%s◉%s Default backend set to: %s\n", log.PRIMARY_COLOR, log.RESET_COLOR, defaultBackend)
	// fmt.Println()

	readme := question("Do you want to create a base README for " + name)
	author := askAuthor(log)

	fmt.Printf("%s◉%s Ok, creating project structure...\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	createFoldersTree(name, log)

	fmt.Printf("%s◉%s Almost there, writing manifest...\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	createManifest(name, log, readme, author)

	fmt.Printf("%s◈%s Finalizing...\n", log.ACCENT_COLOR, log.RESET_COLOR)
	fmt.Printf("\n%s✔%s Project '%s' created.\n", log.SUCESS_COLOR, log.RESET_COLOR, name)

	fmt.Printf("\n%s◉%s Next steps:%s\n", log.PRIMARY_COLOR, log.RESET_COLOR, log.RESET_COLOR)
	fmt.Printf("  cd %s\n", name)
	fmt.Printf("  kc run src/main.k\n")

	log.Line()
}

func createManifest(name string, log Log, readme bool, author string) {
	SRC := filepath.Join(name, "src")
	// TESTS := filepath.Join(name, "tests")

	// backendConfig := generateBackendConfig(backends, defaultBackend)

	// 	tomlContent := fmt.Sprintf(`[project]
	// name = "%s"
	// version = "0.1.0"
	// author = "%s"

	// [build]
	// default-backend = "%s"

	// %s`, name, author, defaultBackend, backendConfig)

	readmeContent := fmt.Sprintf(`# Project: %s
Project created with Loom %s

# Build
%skl
kc build src/main.kl
%s

## Run 
%skl
kc run src/main.kl 
%s
`, name, log.LoomVersion, "```", "```", "```", "```")

	// for _, backend := range backends {
	// 	if backend.Installed {
	// 		readmeContent += fmt.Sprintf("- **%s**: %s\n", backend.Name, backend.Version)
	// 	}
	// }

	files := []string{
		filepath.Join(SRC, "main.kl"),
		// filepath.Join(TESTS, "main_test.k"),
		// filepath.Join(name, "loom.toml"),
		filepath.Join(name, "README.md"),
	}

	filesContents := [][]byte{
		[]byte("@Use(\"java\")\npublic void main(){\n\tprintln(\"Hello from Klar!\");\n\treturn null;\n}"),
		// []byte("test \"project boots\" {\n\tassert(true);\n}"),
		// []byte(tomlContent),
		[]byte(readmeContent),
	}

	for i := 0; i < len(files); i++ {
		file := files[i]
		content := filesContents[i]

		if strings.HasSuffix(file, "README.md") && !readme {
			continue
		}

		if err := makeFile(content, file); err != nil {
			fmt.Printf("%s✖ Error:%s %s\n", log.ERROR_COLOR, log.RESET_COLOR, err)
			return
		}
	}
}

func createFoldersTree(name string, log Log) {
	BASE_FOLDER := name

	SRC := filepath.Join(BASE_FOLDER, "src")
	// TESTS := filepath.Join(BASE_FOLDER, "tests")
	// BUILD := filepath.Join(BASE_FOLDER, "build")
	// BUILD_CACHE := filepath.Join(BUILD, "cache")
	// BUILD_BACK := filepath.Join(BUILD, "backends")
	// KLANG_CACHES := filepath.Join(BASE_FOLDER, ".klar")

	folders := []string{SRC}

	fmt.Printf("%s◉%s Creating folder tree...\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)

		if err != nil {
			fmt.Printf("%s✖ Error:%s %s\n", log.ERROR_COLOR, log.RESET_COLOR, err)
			return
		}
	}
}

var respostasSim = map[string]bool{
	// --- PORTUGUÊS (Variações, Gírias e Formais) ---
	"sim": true, "s": true, "si": true, "sin": true, "sii": true, "siim": true,
	"claro": true, "com certeza": true, "certeza": true, "sem dúvida": true,
	"positivo": true, "afirmativo": true, "confirmo": true, "confirmado": true,
	"verdade": true, "verdadeiro": true, "v": true,
	"ok": true, "okay": true, "okey": true, "k": true, "kk": true,
	"tá": true, "ta": true, "tá bom": true, "tabom": true, "tá certo": true,
	"certo": true, "correto": true, "exato": true, "isso": true, "é isso": true,
	"é": true, "eh": true, "aham": true, "aha": true, "humrum": true,
	"beleza": true, "blz": true, "belê": true,
	"pode": true, "pode ser": true, "pode crer": true,
	"bora": true, "vamos": true, "vai": true, "manda ver": true, "demorou": true,
	"fechou": true, "fechado": true, "combinado": true,
	"aceito": true, "topo": true, "quero": true,
	"uai": true, // Mineiro way

	// --- INGLÊS (Slang, Formal, Internet) ---
	"yes": true, "y": true, "ye": true, "ya": true, "yah": true, "yeh": true,
	"yeah": true, "yep": true, "yup": true, "yess": true, "yas": true, "yass": true,
	"sure": true, "sure thing": true, "for sure": true,
	"ok": true, "okay": true, "okie": true, "k": true, "kk": true, "kay": true,
	"alright": true, "all right": true, "aight": true, "right": true, "righto": true,
	"correct": true, "accurate": true, "positive": true, "affirmative": true,
	"absolutely": true, "definitely": true, "certainly": true, "undoubtedly": true,
	"indeed": true, "agreed": true, "granted": true,
	"aye": true, "aye aye": true,
	"roger": true, "roger that": true, "copy": true, "copy that": true,
	"bet": true, "you bet": true, "totally": true, "totes": true,
	"fine": true, "sounds good": true, "good": true,
	"go": true, "go ahead": true, "proceed": true, "continue": true,
	"enable": true, "enabled": true, "on": true, "active": true,

	// --- ESPANHOL (Variações Regionais) ---
	"sí": true, "si": true, "sip": true, "síp": true,
	"claro": true, "claro que sí": true, "claro que si": true,
	"vale": true, "ya": true, "venga": true,
	"bueno": true, "bue": true,
	"correcto": true, "exacto": true, "cierto": true,
	"por supuesto": true, "desde luego": true, "obvio": true,
	"de acuerdo": true, "está bien": true,
	"dale": true, "va": true, "arre": true, // México/Argentina/etc
	"simón": true, "simon": true, // Gíria Méx
	"okey": true, "sale": true, "listo": true,

	// --- FRANCÊS ---
	"oui": true, "ouais": true, "ouaip": true,
	"d'accord": true, "dac": true, "ok": true,
	"bien sur": true, "bien sûr": true, "absolument": true,
	"exact": true, "effectivement": true, "certes": true,
	"c'est ça": true, "entendu": true, "allez": true, "vas-y": true,

	// --- ALEMÃO ---
	"ja": true, "jo": true, "jep": true, "jupp": true,
	"sicher": true, "sicherlich": true, "klar": true, "alles klar": true,
	"genau": true, "stimmt": true, "richtig": true,
	"einverstanden": true, "ok": true, "okay": true,

	// --- ITALIANO ---
	"sì": true, "si": true, "già": true,
	"certo": true, "certamente": true, "sicuro": true,
	"va bene": true, "vabene": true, "ok": true, "d'accordo": true,
	"giusto": true, "esatto": true, "perfetto": true,

	// --- OUTRAS LÍNGUAS (Principais e Representativas) ---
	// Russo
	"da": true, "да": true, "konechno": true, "aga": true,
	// Japonês (Romaji + Kanji/Kana)
	"hai": true, "ha": true, "ee": true, "sou": true, "sou desu": true,
	"はい": true, "ええ": true, "そうです": true,
	// Chinês (Pinyin + Hanzi - Simplificado/Tradicional)
	"shi": true, "dui": true, "hao": true, "xing": true, "ok": true,
	"是": true, "是的": true, "对": true, "好": true, "行": true,
	// Coreano
	"ne": true, "ye": true, "ung": true, "eung": true,
	"네": true, "예": true, "응": true,
	// Árabe (Transliterado + Script)
	"na'am": true, "naam": true, "aiwa": true, "yani": true, "n": true, // n muitas vezes mapeado para naam em sistemas
	"نعم": true, "ايوه": true,
	// Hindi
	"haan": true, "ha": true, "ji": true, "sahi": true,
	"हाँ": true, "जी": true,
	// Holandês
	"ja": true, "jawel": true, "jep": true, "oké": true,
	// Polonês
	"tak": true, "dobrze": true, "jasne": true,
	// Turco
	"evet": true, "he": true, "tamam": true, "peki": true,
	// Sueco/Norueguês/Dinamarquês
	"ja": true, "jo": true, "joo": true, "javisst": true,
	// Grego
	"ne": true, "nai": true, "ναι": true,
	// Tcheco
	"ano": true, "jo": true,

	// --- UNIVERSAIS / TÉCNICOS / SÍMBOLOS ---
	"1": true, "true": true, "t": true,
	"+": true, "plus": true,
	"✓": true, "✔": true, "☑": true,
	"👍": true, "👌": true, "🙆": true, "🙆‍♂️": true, "🙆‍♀️": true,
	"infinite": true, "always": true,
	"success": true, "pass": true,
}

func question(quest string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(quest + " (s/n) [default = n]: ")

	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	return respostasSim[input]
}

func askAuthor(log Log) string {
	reader := bufio.NewReader(os.Stdin)
	input := "outsider"
	qtd := 0

	for true {
		qtd++
		fmt.Printf("%s◉%s Author name: ", log.PRIMARY_COLOR, log.RESET_COLOR)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Printf("%s✖ Error:%s your name cannot be empty!\n", log.ERROR_COLOR, log.RESET_COLOR)

			if qtd%5 == 0 {
				ok := false

				for true {
					fmt.Print("...Do you want to follow your nameless course? (s/n): ")

					input, _ := reader.ReadString('\n')
					input = strings.TrimSpace(input)

					if input != "" {
						if respostasSim[input] {
							input = "outsider"
							ok = true
						}

						break
					}
				}

				if ok {
					break
				}

				continue
			}

			continue
		}

		break
	}

	return input
}
