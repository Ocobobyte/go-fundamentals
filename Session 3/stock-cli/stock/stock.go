package stock

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	maxQuantity = 100
	lowStock    = 5
)

type productName string

type Product struct {
	Name     string
	Quantity int
}

// Stock holds products indexed by a normalized product name.
type Stock map[productName]Product

// Manager encapsulates stock data and the file used for persistence.
type StockManager struct {
	stock    Stock
	filePath string
}

// NewManager initializes a new Manager and loads existing stock data.
func NewStockManager(filePath string) (*StockManager, error) {
	m := &StockManager{
		stock:    make(Stock),
		filePath: filePath,
	}

	// If the file doesn't exist, create it.
	if _, err := os.Stat(filePath); err != nil {
		fmt.Println("Stock file not found, creating new one")
		if err := os.WriteFile(filePath, []byte("{}"), 0644); err != nil {
			return nil, fmt.Errorf("failed to create stock file: %w", err)
		}
	}

	if err := m.load(); err != nil {
		return nil, fmt.Errorf("failed to load stock: %w", err)
	}

	return m, nil
}

func (m *StockManager) load() error {
	data, err := os.ReadFile(m.filePath)
	if err != nil {
		return fmt.Errorf("error reading stock file: %w", err)
	}
	var s Stock
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("error unmarshalling stock data: %w", err)
	}
	m.stock = s
	return nil
}

func (m *StockManager) save() error {
	data, err := json.MarshalIndent(m.stock, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling stock data: %w", err)
	}
	if err := os.WriteFile(m.filePath, data, 0644); err != nil {
		return fmt.Errorf("error writing stock file: %w", err)
	}
	return nil
}

func (m *StockManager) productKey(name string) productName {
	return productName(strings.ToLower(name))
}

func (m *StockManager) Add(p Product) error {
	if p.Quantity < 0 {
		return fmt.Errorf("cannot add product with negative quantity")
	}

	key := m.productKey(p.Name)
	if _, exists := m.stock[key]; exists {
		return fmt.Errorf("product %s already exists", p.Name)
	}
	if p.Quantity > maxQuantity {
		return fmt.Errorf("quantity %d exceeds maximum stock of %d", p.Quantity, maxQuantity)
	}
	m.stock[key] = p
	return m.save()
}

func (m *StockManager) Update(p Product) error {
	key := m.productKey(p.Name)
	existing, exists := m.stock[key]
	if !exists {
		return fmt.Errorf("product %s does not exist", p.Name)
	}
	if existing.Quantity+p.Quantity > maxQuantity {
		return fmt.Errorf("update would exceed maximum stock of %d: current %d, attempted addition %d", maxQuantity, existing.Quantity, p.Quantity)
	}
	existing.Quantity += p.Quantity
	m.stock[key] = existing
	return m.save()
}

func (m *StockManager) Delete(name string) error {
	key := m.productKey(name)
	if _, exists := m.stock[key]; !exists {
		return fmt.Errorf("product %s does not exist", name)
	}
	delete(m.stock, key)
	return m.save()
}

func (m *StockManager) LowStock() []Product {
	var lowStockProducts []Product
	for _, p := range m.stock {
		if p.Quantity <= lowStock {
			lowStockProducts = append(lowStockProducts, p)
		}
	}
	return lowStockProducts
}

func (m *StockManager) List(limit int) []Product {
	var products []Product
	for _, p := range m.stock {
		if limit > 0 && len(products) >= limit {
			break
		}
		products = append(products, p)
	}
	return products
}
