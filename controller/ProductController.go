package controllers

import(
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/NutchaPK/test/ent"
	"github.com/NutchaPK/test/ent/gender"
	"github.com/NutchaPK/test/ent/product"
	"github.com/NutchaPK/test/ent/size"
	"github.com/NutchaPK/test/ent/style"
)

type ProductController struct {
	client *ent.Client
	router gin.IRouter
}

type Product struct {
	Gender 	int
	Style 	int
	Size 	int
	Price 	int

}

type Gender struct {
	Gender 	string

}
type Style struct {

	Style 	string

}
type Size struct {

	Size 	string


}

// @Summary List a product entity
// @Description get product
// @ID get-product
// @Produce  json
// @Param limit  path int true "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} ent.Product
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products/limit/{limit} [get]
func (ctl *ProductController) ListProductWithLimit (c *gin.Context) {
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		c.JSON(400,gin.H{
			"error":err.Error(),
		})
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	products, err := ctl.client.Product.
	Query().
	WithGender().
	WithStyle().
	WithSize().
	Limit(limit).
	Offset(offset).
	All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, products)
	
}


// @Summary List a product entity
// @Description get product
// @ID get-product
// @Produce  json
// @Param page  path int true "Page"
// @Success 200 {object} ent.Product
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products/page/{page} [get]
func (ctl *ProductController) ListProductWithPage (c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.JSON(400,gin.H{
			"error":err.Error(),
		})
	}
	limit :=10
	if page != 0 {
		limit = page * 10
	}

	offset := limit - 10
	products, err := ctl.client.Product.
	Query().
	WithGender().
	WithStyle().
	WithSize().
	Limit(limit).
	Offset(offset).
	All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, products)
	
}


// @Summary Get a product entity
// @Description get product
// @ID get-product
// @Produce  json
// @Param style  path string true "style"
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} ent.Product
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products/style/{style} [get]
func (ctl *ProductController) FilterByStyle (c *gin.Context) {
	
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}
	styleParam:=c.Param("style")
	if styleParam == "" {
		products, err := ctl.client.Product.
		Query().
		WithGender().
		WithStyle().
		WithSize().
		Limit(limit).
		Offset(offset).
	
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
		}
		c.JSON(200, products)
	} else {
		products, err := ctl.client.Product.
		Query().
		WithGender().
		WithStyle().
		WithSize().
		Limit(limit).
		Offset(offset).
		Where(product.HasStyleWith(style.StyleContains(styleParam))).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
		}
		c.JSON(200, products)
	}
}
	
// @Summary Get a product entity
// @Description get product
// @ID get-product
// @Produce  json
// @Param gender  path string true "gender"
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} ent.Product
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products/gender/{gender} [get]
func (ctl *ProductController) FilterByGender (c *gin.Context){
	
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}
	genderParam:=c.Param("gender")
	if genderParam == "" {
		products, err := ctl.client.Product.
		Query().
		WithGender().
		WithStyle().
		WithSize().
		Limit(limit).
		Offset(offset).
	
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
		}
		c.JSON(200, products)
	} else {
		products, err := ctl.client.Product.
		Query().
		WithGender().
		WithStyle().
		WithSize().
		Limit(limit).
		Offset(offset).
		Where(product.HasGenderWith(gender.GenderContains(genderParam))).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
		}
		c.JSON(200, products)
	}
	
	
}

// @Summary Get a product entity
// @Description get product
// @ID get-product
// @Produce  json
// @Param size  path string true "size"
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} ent.Product
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products/size/{size} [get]
func (ctl *ProductController) FilterBySize (c *gin.Context) {
	
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}
	sizeParam:=c.Param("size")
	if sizeParam == "" {
		products, err := ctl.client.Product.
		Query().
		WithGender().
		WithStyle().
		WithSize().
		Limit(limit).
		Offset(offset).
	
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
		}
		c.JSON(200, products)
	} else {
		products, err := ctl.client.Product.
		Query().
		WithGender().
		WithStyle().
		WithSize().
		Limit(limit).
		Offset(offset).
		Where(product.HasSizeWith(size.SizeContains(sizeParam))).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
		}
		c.JSON(200, products)
	}
	
	
}



func NewProductController(router gin.IRouter, client *ent.Client) *ProductController {
	uc := &ProductController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *ProductController) register() {
	products := ctl.router.Group("/products")
	products .GET("gender/:gender", ctl.FilterByGender)
	products .GET("style/:style", ctl.FilterByStyle)
	products .GET("size/:size", ctl.FilterBySize)
	products .GET("limit/:limit", ctl.ListProductWithLimit)
	products .GET("page/:page", ctl.ListProductWithPage)
}

