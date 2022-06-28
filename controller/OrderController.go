package controllers
 
import (
	"context"
	"strconv"
	"time"

	"github.com/NutchaPK/test/ent"
	"github.com/NutchaPK/test/ent/order"
	"github.com/NutchaPK/test/ent/product"
	"github.com/gin-gonic/gin"

)

type OrderController struct{
	client *ent.Client
	router gin.IRouter
}

type Order struct {
	Product int
	Status	string
	ShippingAddress string
	Date	string
}


// Createorder handles POST requests for adding order entities
// @Summary Create order
// @Description Create order
// @ID create-order
// @Accept   json
// @Produce  json
// @Param order body order true "order entity"
// @Success 200 {object} ent.order
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders [post]
func (ctl *OrderController) Createorder(c *gin.Context) {
	obj := Order{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "order binding failed",
		})
		return
	}
	product, err := ctl.client.Product.
		Query().
		Where(product.IDEQ(int(obj.Product))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "order not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.Date)

	corder, err := ctl.client.Order.
		Create().
		SetProduct(product).
		SetShippingAddress(obj.ShippingAddress).
		SetStatus(obj.Status).
		SetPaidDate(times).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}


	c.JSON(200, corder)
}


// @Summary List a order entity
// @Description get order
// @ID get-order
// @Produce  json
// @Param limit  path int true "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} ent.Order
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders/limit/{limit} [get]
func (ctl *OrderController) ListOrderWithLimit (c *gin.Context) {
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

	orders, err := ctl.client.Order.
	Query().
	WithProduct().
	Limit(limit).
	Offset(offset).
	All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, orders)
	
}


// @Summary List a order entity
// @Description get order
// @ID get-product
// @Produce  json
// @Param page  path int true "Page"
// @Success 200 {object} ent.Order
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders/page/{page} [get]
func (ctl *OrderController) ListOrderWithPage (c *gin.Context) {
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
	orders, err := ctl.client.Order.
	Query().
	WithProduct().
	Limit(limit).
	Offset(offset).
	All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, orders)
	
}

// @Summary Get a order entity
// @Description get order
// @ID get-order
// @Produce  json
// @Param status  path string true "size"
// @Success 200 {object} ent.Order
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders/status/{status} [get]
func (ctl *OrderController) FilterByStatus (c *gin.Context) {
	
	statusParam:=c.Param("status")
	if statusParam == "" {
		orders, err := ctl.client.Order.
		Query().
		WithProduct().
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
		}
		c.JSON(200, orders)
	} else {
		orders, err := ctl.client.Order.
		Query().
		WithProduct().
		Where(order.StatusContains(statusParam)).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
		}
		c.JSON(200, orders)
	}
	
	
}


func NewOrderController(router gin.IRouter, client *ent.Client) *OrderController {
	uc := &OrderController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *OrderController) register() {
	orders := ctl.router.Group("/products")
	orders .POST("", ctl.Createorder)
	orders .GET("status/:status", ctl.FilterByStatus)
	orders .GET("limit/:limit", ctl.ListOrderWithLimit)
	orders .GET("page/:page", ctl.ListOrderWithPage)
}