package api

import (
	"market_go_apigateway/api/handlers"
	"market_go_apigateway/config"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()


	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	// Sale
	r.POST("/sale/create", h.CreateSale)
	r.GET("/sale", h.GetAllSale)
	r.DELETE("/sale/delete/:id", h.DeleteSale)
	r.GET("/sale/:id", h.GetById)
	r.PUT("/sale/update/:id", h.UpdateSale)
	
	r.POST("/code_prod", h.CodeProd)

	// SaleProducts
	r.POST("/sale_products/create", h.CreateSaleProducts)
	r.GET("/sale_products", h.GetAllSaleProducts)
	r.DELETE("/sale_products/delete/:id", h.DeleteSaleProducts)
	r.GET("/sale_products/:id", h.GetByIdSaleProducts)
	r.PUT("/sale_products/update/:id", h.UpdateSaleProducts)

	r.POST("/sale_products/paytype", h.Paytype)

	r.POST("/sale_products/prodaja", h.Prodaja)



	//Change
	r.POST("/change/create", h.CreateChange)
	r.GET("/change", h.GetAllChange)
	r.DELETE("/change/delete/:id", h.DeleteChange)
	r.GET("/change/:id", h.GetByIdChange)
	r.PUT("/change/update/:id", h.UpdateChange)
	r.POST("/change/open_change", h.OpenChange)
	r.POST("/change/close_change", h.CloseChange)

	

	//Branch
	r.POST("/branch/create", h.CreateBranch)
	r.GET("/branch", h.GetAllBranch)
	r.DELETE("/branch/delete/:id", h.DeleteBranch)
	r.GET("/branch/:id", h.GetByIdBranch)
	r.PUT("/branch/update/:id", h.UpdateBranch)

	//Product
	r.POST("/product/create", h.CreateProduct)
	r.GET("/product", h.GetAllProduct)
	r.DELETE("/product/delete/:id", h.DeleteProduct)
	r.GET("/product/:id", h.GetByIdProduct)
	r.PUT("/product/update/:id", h.UpdateProduct)

	//Category
	r.POST("/category/create", h.CreateCategory)
	r.GET("/category", h.GetAllCategory)
	r.DELETE("/category/delete/:id", h.DeleteCategory)
	r.GET("/category/:id", h.GetByIdCategory)
	r.PUT("/category/update/:id", h.UpdateCategory)




	// Remainder
	r.POST("/remainder/create", h.CreateRemainder)
	r.GET("/remainder", h.GetAllRemainder)
	r.DELETE("/remainder/delete/:id", h.DeleteRemainder)
	r.GET("/remainder/:id", h.GetByIdRemainder)
	r.PUT("/remainder/update/:id", h.UpdateRemainder)
	

	

	// //Coming
	r.POST("/coming/create", h.CreateComing)
	r.GET("/coming", h.GetAllComing)
	r.DELETE("/coming/delete/:id", h.DeleteComing)
	r.GET("/coming/:id", h.GetByIdComing)
	r.PUT("/coming/update/:id", h.UpdateComing)

	r.POST("coming/sdelat_prihod", h.SdelatPrihod)

	// //ComingProducts
	r.POST("/coming_products/create", h.CreateComingProducts)
	r.GET("/coming_products", h.GetAllComingProducts)
	r.DELETE("/coming_products/delete/:id", h.DeleteComingProducts)
	r.GET("/coming_products/:id", h.GetByIdComingProducts)
	r.PUT("/coming_products/update/:id", h.UpdateComingProducts)


	//Employees
	r.POST("/employees/create", h.CreateEmployees)
	r.GET("/employees", h.GetAllEmployees)
	r.DELETE("/employees/delete/:id", h.DeleteEmployees)
	r.GET("/employees/:id", h.GetByIdEmployees)
	r.PUT("/employees/update/:id", h.UpdateEmployees)

	//Provider
	r.POST("/provider/create", h.CreateProvider)
	r.GET("/provider", h.GetAllProvider)
	r.DELETE("/provider/delete/:id", h.DeleteProvider)
	r.GET("/provider/:id", h.GetByIdProvider)
	r.PUT("/provider/update/:id", h.UpdateProvider)

	//Magazine
	r.POST("/magazine/create", h.CreateMagazine)
	r.GET("/magazine", h.GetAllMagazine)
	r.DELETE("/magazine/delete/:id", h.DeleteMagazine)
	r.GET("/magazine/:id", h.GetByIdMagazine)
	r.PUT("/magazine/update/:id", h.UpdateMagazine)


	



	return
}
