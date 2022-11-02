package front

import (
	"github.com/gin-gonic/gin"
	c "github.com/wujiyu98/ginframe/app/front/internal/controller"
)

func Init(e *gin.Engine) {

	r := e.Group("/")
	r.GET("/", c.Index)
	r.GET("/contact", c.Contact)
	r.GET("/about/:pathname", c.About)
	r.GET("/info/:pathname", c.Info)
	r.GET("/manufacturers", c.Manufacturers)
	r.GET("/manufacturers/:pathname", c.Manufacturer)
	r.GET("/category", c.Categories)
	r.GET("/category/:pathname", c.Category)
	r.GET("/product/:pathname", c.Product)
	r.GET("/news-category/*pathname", c.NewsCategory)
	r.GET("/news/:pathname", c.News)
	r.GET("/enquiry", c.Enquiry)

}
