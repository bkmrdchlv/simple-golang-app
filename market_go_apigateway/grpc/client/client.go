package client

import (
	"market_go_apigateway/config"
	"market_go_apigateway/genproto/branch"
	"market_go_apigateway/genproto/category"
	"market_go_apigateway/genproto/change"
	"market_go_apigateway/genproto/coming"
	"market_go_apigateway/genproto/employees"
	"market_go_apigateway/genproto/magazine"
	"market_go_apigateway/genproto/product"
	"market_go_apigateway/genproto/provider"
	"market_go_apigateway/genproto/remainder"
	"market_go_apigateway/genproto/sale"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	BranchService() branch.BranchServiceClient
	CategoryService() category.CategoryServiceClient
	ChangeService() change.ChangeServiceClient

	ComingService() coming.ComingServiceClient
	ComingProductsService() coming.ComingProductsServiceClient

	EmployeesService() employees.EmployeesServiceClient
	MagazineService() magazine.MagazineServiceClient
	ProductService() product.ProductServiceClient
	ProvicderService() provider.ProviderServiceClient
	RemainderService() remainder.RemainderServiceClient

	SaleService() sale.SaleServiceClient
	SaleProductService() sale.SaleProductsServiceClient
}

type grpcClients struct {
	branchService   branch.BranchServiceClient
	categoryService category.CategoryServiceClient
	changeService   change.ChangeServiceClient

	comingService         coming.ComingServiceClient
	comingProductsService coming.ComingProductsServiceClient

	employeesService employees.EmployeesServiceClient
	magazineService  magazine.MagazineServiceClient
	productService   product.ProductServiceClient
	providerService  provider.ProviderServiceClient
	remainderService remainder.RemainderServiceClient

	saleService         sale.SaleServiceClient
	saleProductsService sale.SaleProductsServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	connContentService, err := grpc.Dial(
		cfg.ContentServiceHost+cfg.ContentGRPCPort,
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, err
	}

	connStockService, err := grpc.Dial(
		cfg.StocktentServiceHost+cfg.StockGRPCPort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	connSaleService, err := grpc.Dial(
		cfg.SaleServiceHost+cfg.SaleGRPCPort,
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, err
	}

	return &grpcClients{
		branchService:         branch.NewBranchServiceClient(connContentService),
		categoryService:       category.NewCategoryServiceClient(connContentService),
		changeService:         change.NewChangeServiceClient(connSaleService),
		comingService:         coming.NewComingServiceClient(connStockService),
		comingProductsService: coming.NewComingProductsServiceClient(connStockService),
		employeesService:      employees.NewEmployeesServiceClient(connContentService),
		magazineService:       magazine.NewMagazineServiceClient(connContentService),
		productService:        product.NewProductServiceClient(connContentService),
		providerService:       provider.NewProviderServiceClient(connContentService),
		remainderService:      remainder.NewRemainderServiceClient(connStockService),
		saleService:           sale.NewSaleServiceClient(connSaleService),
		saleProductsService:   sale.NewSaleProductsServiceClient(connSaleService),
	}, nil

}

func (g *grpcClients) BranchService() branch.BranchServiceClient {
	return g.branchService
}
func (g *grpcClients) CategoryService() category.CategoryServiceClient {
	return g.categoryService
}
func (g *grpcClients) ChangeService() change.ChangeServiceClient {
	return g.changeService
}

func (g *grpcClients) ComingService() coming.ComingServiceClient {
	return g.comingService
}
func (g *grpcClients) ComingProductsService() coming.ComingProductsServiceClient {
	return g.comingProductsService
}

func (g *grpcClients) EmployeesService() employees.EmployeesServiceClient {
	return g.employeesService
}
func (g *grpcClients) MagazineService() magazine.MagazineServiceClient {
	return g.magazineService
}
func (g *grpcClients) ProductService() product.ProductServiceClient {
	return g.productService
}
func (g *grpcClients) ProvicderService() provider.ProviderServiceClient {
	return g.providerService
}
func (g *grpcClients) RemainderService() remainder.RemainderServiceClient {
	return g.remainderService
}

func (g *grpcClients) SaleService() sale.SaleServiceClient {
	return g.saleService
}
func (g *grpcClients) SaleProductService() sale.SaleProductsServiceClient {
	return g.saleProductsService
}
