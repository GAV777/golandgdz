package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"golandgdz/storage"
	"log"
	"os"
)

const (
	storageFile = "bins.json"
	version     = "1.0.0"
)

func main() {
	// Создаем хранилище
	storageManager := storage.NewFileStorage()

	// Загружаем существующие данные
	if err := storageManager.LoadFromFile(storageFile); err != nil {
		log.Printf("Warning: failed to load storage: %v", err)
	}

	// Сохраняем данные при выходе
	defer func() {
		if err := storageManager.SaveToFile(storageFile); err != nil {
			log.Printf("Failed to save storage: %v", err)
		}
	}()

	// Определяем флаги командной строки
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createName := createCmd.String("name", "", "Bin name")
	createPrivate := createCmd.Bool("private", false, "Make bin private")
	createFile := createCmd.String("file", "", "JSON file to upload")

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getID := getCmd.String("id", "", "Bin ID")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateID := updateCmd.String("id", "", "Bin ID")
	updateName := updateCmd.String("name", "", "New bin name")
	updatePrivate := updateCmd.Bool("private", false, "Set private flag")
	updateFile := updateCmd.String("file", "", "JSON file to update content")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.String("id", "", "Bin ID to delete")

	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	// Обрабатываем команды
	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		handleCreate(storageManager, *createName, *createPrivate, *createFile)
	case "get":
		getCmd.Parse(os.Args[2:])
		handleGet(storageManager, *getID)
	case "list":
		listCmd.Parse(os.Args[2:])
		handleList(storageManager)
	case "update":
		updateCmd.Parse(os.Args[2:])
		handleUpdate(storageManager, *updateID, *updateName, *updatePrivate, *updateFile)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		handleDelete(storageManager, *deleteID)
	case "version":
		versionCmd.Parse(os.Args[2:])
		fmt.Printf("CloudBins CLI v%s\n", version)
	case "help":
		printHelp()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printHelp()
		os.Exit(1)
	}
}

func handleCreate(sm storage.StorageManager, name string, private bool, filePath string) {
	if name == "" {
		log.Fatal("Name is required for create command")
	}

	var content json.RawMessage
	if filePath != "" {
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}

		// Валидируем JSON
		if !json.Valid(data) {
			log.Fatal("File does not contain valid JSON")
		}

		content = json.RawMessage(data)
	} else {
		content = json.RawMessage("{}")
	}

	bin, err := sm.CreateBin(name, content, private)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bin created successfully!\n")
	fmt.Printf("ID: %s\n", bin.ID)
	fmt.Printf("Name: %s\n", bin.Name)
	fmt.Printf("Private: %v\n", bin.Private)
	fmt.Printf("Created at: %s\n", bin.CreatedAt.Format("2006-01-02 15:04:05"))
}

func handleGet(sm storage.StorageManager, id string) {
	if id == "" {
		log.Fatal("ID is required for get command")
	}

	bin, err := sm.GetBin(id)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(bin, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func handleList(sm storage.StorageManager) {
	binList, err := sm.ListBins()
	if err != nil {
		log.Fatal(err)
	}

	if binList.Total == 0 {
		fmt.Println("No bins found")
		return
	}

	fmt.Printf("Total bins: %d\n\n", binList.Total)
	for _, bin := range binList.Bins {
		fmt.Printf("ID: %s\n", bin.ID)
		fmt.Printf("Name: %s\n", bin.Name)
		fmt.Printf("Private: %v\n", bin.Private)
		fmt.Printf("Created: %s\n", bin.CreatedAt.Format("2006-01-02"))
		fmt.Println("---")
	}
}

func handleUpdate(sm storage.StorageManager, id, name string, private bool, filePath string) {
	if id == "" {
		log.Fatal("ID is required for update command")
	}

	updates := make(map[string]interface{})

	if name != "" {
		updates["name"] = name
	}

	// Проверяем, был ли передан флаг private
	for _, arg := range os.Args {
		if arg == "-private" || arg == "--private" {
			updates["private"] = private
			break
		}
	}

	if filePath != "" {
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}

		if !json.Valid(data) {
			log.Fatal("File does not contain valid JSON")
		}

		updates["content"] = json.RawMessage(data)
	}

	if len(updates) == 0 {
		log.Fatal("No updates specified")
	}

	bin, err := sm.UpdateBin(id, updates)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bin updated successfully!\n")
	fmt.Printf("ID: %s\n", bin.ID)
	fmt.Printf("Name: %s\n", bin.Name)
	fmt.Printf("Private: %v\n", bin.Private)
}

func handleDelete(sm storage.StorageManager, id string) {
	if id == "" {
		log.Fatal("ID is required for delete command")
	}

	if err := sm.DeleteBin(id); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bin deleted successfully")
}

func printHelp() {
	helpText := `CloudBins CLI - Manage JSON bins in cloud storage

Usage:
  cloudbins <command> [flags]

Commands:
  create    Create a new bin
  get       Get a bin by ID
  list      List all bins
  update    Update a bin
  delete    Delete a bin
  version   Show version
  help      Show this help message

Examples:
  cloudbins create --name "My Data" --file data.json
  cloudbins get --id "bin-id-here"
  cloudbins list
  cloudbins update --id "bin-id" --name "New Name" --private
  cloudbins delete --id "bin-id"

For command-specific flags, run: cloudbins <command> --help`

	fmt.Println(helpText)
}
