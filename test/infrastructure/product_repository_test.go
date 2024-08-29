package infrastructure

import (
	"example.com/product-api/common/postgresql"
	"example.com/product-api/domain"
	"example.com/product-api/persistance"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"os"
	"testing"
)

var productRepository persistance.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "5432",
		UserName:              "postgres",
		Password:              "153515",
		DbName:                "workshops",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "10s",
	})

	productRepository = persistance.NewProductRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}

func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {
	setup(ctx, dbPool)

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Iron",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Washing Machine",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Floor Lamp",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Decoration Palace",
		},
	}

	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAll()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool)
}

func TestGetById(t *testing.T) {
	setup(ctx, dbPool)

	expectedProduct := domain.Product{
		Id:       1,
		Name:     "AirFryer",
		Price:    3000.0,
		Discount: 22.0,
		Store:    "ABC TECH",
	}

	t.Run("GetById", func(t *testing.T) {
		actualProduct, _ := productRepository.GetById(1)
		assert.Equal(t, expectedProduct, actualProduct)
	})

	clear(ctx, dbPool)
}

func TestGetAllProductsByStore(t *testing.T) {
	setup(ctx, dbPool)

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Iron",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Washing Machine",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
	}

	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProducts := productRepository.GetAllByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool)
}

func TestAddProduct(t *testing.T) {
	expectedProduct := []domain.Product{
		{
			Id:       1,
			Name:     "Telephone",
			Price:    20000.0,
			Discount: 10.0,
			Store:    "Samsung",
		},
	}
	newProduct := domain.Product{
		Name:     "Telephone",
		Price:    20000.0,
		Discount: 10.0,
		Store:    "Samsung",
	}

	t.Run("AddProduct", func(t *testing.T) {
		productRepository.Add(newProduct)
		actualProducts := productRepository.GetAll()
		assert.Equal(t, 1, len(actualProducts))
		assert.Equal(t, expectedProduct, actualProducts)
	})

	clear(ctx, dbPool)
}