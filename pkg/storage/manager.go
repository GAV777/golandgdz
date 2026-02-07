package storage

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
	"time"
)

// StorageManager управляет операциями с Bins
type StorageManager interface {
	CreateBin(name string, content json.RawMessage, private bool) (*types.Bin, error)
	GetBin(id string) (*types.Bin, error)
	ListBins() (*types.BinList, error)
	UpdateBin(id string, updates map[string]interface{}) (*types.Bin, error)
	DeleteBin(id string) error
	LoadFromFile(filename string) error
	SaveToFile(filename string) error
}

// FileStorage реализация хранилища в файле
type FileStorage struct {
	bins map[string]types.Bin
}

// NewFileStorage создает новое файловое хранилище
func NewFileStorage() *FileStorage {
	return &FileStorage{
		bins: make(map[string]types.Bin),
	}
}

// CreateBin создает новый Bin
func (fs *FileStorage) CreateBin(name string, content json.RawMessage, private bool) (*types.Bin, error) {
	if name == "" {
		return nil, fmt.Errorf("bin name cannot be empty")
	}

	// Генерируем уникальный ID
	id := uuid.New().String()

	bin := types.Bin{
		ID:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
		Content:   content,
	}

	fs.bins[id] = bin
	return &bin, nil
}

// GetBin получает Bin по ID
func (fs *FileStorage) GetBin(id string) (*types.Bin, error) {
	bin, exists := fs.bins[id]
	if !exists {
		return nil, fmt.Errorf("bin with id %s not found", id)
	}
	return &bin, nil
}

// ListBins возвращает список всех Bins
func (fs *FileStorage) ListBins() (*types.BinList, error) {
	bins := make([]types.Bin, 0, len(fs.bins))

	for _, bin := range fs.bins {
		bins = append(bins, bin)
	}

	return &types.BinList{
		Bins:  bins,
		Total: len(bins),
	}, nil
}

// UpdateBin обновляет Bin
func (fs *FileStorage) UpdateBin(id string, updates map[string]interface{}) (*types.Bin, error) {
	bin, exists := fs.bins[id]
	if !exists {
		return nil, fmt.Errorf("bin with id %s not found", id)
	}

	// Обновляем поля
	if name, ok := updates["name"].(string); ok {
		bin.Name = name
	}

	if private, ok := updates["private"].(bool); ok {
		bin.Private = private
	}

	if content, ok := updates["content"].(json.RawMessage); ok {
		bin.Content = content
	}

	fs.bins[id] = bin
	return &bin, nil
}

// DeleteBin удаляет Bin
func (fs *FileStorage) DeleteBin(id string) error {
	if _, exists := fs.bins[id]; !exists {
		return fmt.Errorf("bin with id %s not found", id)
	}

	delete(fs.bins, id)
	return nil
}

// LoadFromFile загружает данные из файла
func (fs *FileStorage) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// Файл не существует, начинаем с пустого хранилища
			return nil
		}
		return fmt.Errorf("failed to read file: %w", err)
	}

	var binList types.BinList
	if err := json.Unmarshal(data, &binList); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	// Загружаем bins в map
	fs.bins = make(map[string]types.Bin)
	for _, bin := range binList.Bins {
		fs.bins[bin.ID] = bin
	}

	return nil
}

// SaveToFile сохраняет данные в файл
func (fs *FileStorage) SaveToFile(filename string) error {
	binList, err := fs.ListBins()
	if err != nil {
		return fmt.Errorf("failed to list bins: %w", err)
	}

	data, err := json.MarshalIndent(binList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
