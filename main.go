package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

// FileWatcher - структура для инкапсуляции логики отслеживания файла
type FileWatcher struct {	//инкапсуляция
	FilePath string //путь файла
	Watcher  *fsnotify.Watcher //объект для отслеживания изменений
	LogFile  *os.File //файл логов
	StopChan chan bool // Канал для сигнала остановки
}

// NewFileWatcher - конструктор для FileWatcher
func NewFileWatcher(filePath string) (*FileWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
// Открытие файла для записи логов
	logFile, err := os.OpenFile("file_changes.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия лог файла: %v", err)
	}

	return &FileWatcher{
		FilePath: filePath,
		Watcher:  watcher,
		LogFile:  logFile,
		StopChan: make(chan bool),
	}, nil
}

// Start - метод для запуска отслеживания файла, Запускает горутину
func (fw *FileWatcher) Start() error {
	err := fw.Watcher.Add(fw.FilePath)
	if err != nil {
		return err
	}

	log.SetOutput(fw.LogFile)
	fmt.Printf("Начато отслеживание файла: %s\n", fw.FilePath)
// Запуск горутины для обработки событий от Watcher
	go func() {
		for {
			select {
			case event, ok := <-fw.Watcher.Events:	// Ожидание события от Watcher
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Printf("Файл изменен: %s\n", event.Name)
				}
			case err, ok := <-fw.Watcher.Errors:	// Обработка ошибок
				if !ok {
					return
				}
				log.Println("Ошибка:", err)
			case <-fw.StopChan:	// Обработка сигнала остановки
				fmt.Println("Отслеживание остановлено.")
				return
			}
		}
	}()

	return nil
}

// Stop - метод для остановки отслеживания файла
func (fw *FileWatcher) Stop() {
	fw.StopChan <- true	// Отправка сигнала остановки в канал
	fw.Watcher.Close()
	fw.LogFile.Close()
}
// Вспомогательная функция для получения пути к файлу из диалогового окна.
func getFilePathFromDialog() (string, error) {
	cmd := exec.Command("python", "dialog.py")
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("ошибка запуска диалога: %v", err)
	}

	reader := bufio.NewReader(stdout)
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", fmt.Errorf("ошибка чтения из диалога: %v", err)
	}

	return string(line), nil
}

func main() {
	filePath, err := getFilePathFromDialog()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Выбранный файл:", filePath)

	watcher, err := NewFileWatcher(filePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := watcher.Start(); err != nil {
		log.Fatal(err)
	}

	// Ожидание нажатия клавиши Enter (OK)
	var input rune
	fmt.Println("Нажмите Enter для выхода...")
	fmt.Scanf("%c", &input)

	watcher.Stop()
	fmt.Println("Программа завершена.")
}