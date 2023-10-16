package routes

import (
	api "github.com/CocaineCong/gin-mall/api/v1"
	"github.com/CocaineCong/gin-mall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1") //抽离出的公共域名
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		//轮播图
		v1.GET("carousels", api.ListCarousel)

		// 商品操作
		v1.GET("products", api.ListProduct)
		v1.GET("products/:id", api.ShowProduct) //todo 这种属于是路径参数，不需要从表格中提取.其形式多为:{{url}}products/1 {{url}}products/id
		v1.GET("products/imgs/:id", api.ListProductImg)
		v1.GET("category/list", api.ListCategory)

		//放在这个Group外面的表示为不需要登录验证，即token的操作.反之，放在Group内的则需要进行token的验证.
		authed := v1.Group("/") //嵌套式登录，需要登录保护
		authed.Use(middleware.JWT())
		{
			// 	用户操作
			authed.PUT("user/:id", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			//显示金额
			authed.POST("money", api.ShowMoney)

			//商品操作
			authed.POST("product", api.CreateProduct)
			authed.POST("products", api.SearchProduct)

			//收藏夹操作
			authed.GET("favorites/list", api.ListFavorites)
			authed.POST("favorites/create", api.CreateFavorites)
			authed.DELETE("favorites/delete/:id", api.DeleteFavorites)

			//地址操作
			authed.POST("addresses", api.CreateAddress)
			authed.GET("addresses/:id", api.ShowAddress)
			authed.GET("addresses", api.ListAddress)
			authed.PUT("addresses/:id", api.UpdateAddress)
			authed.DELETE("addresses/:id", api.DeleteAddress)

			//	购物车操作
			authed.POST("carts", api.CreateCart)
			authed.GET("carts", api.ListCart)
			authed.PUT("carts/:id", api.UpdateCart)
			authed.DELETE("carts/:id", api.DeleteCart)

			// 订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("orders", api.ListOrder)
			authed.PUT("orders/:id", api.ShowOrder)
			authed.DELETE("orders/:id", api.DeleteOrder)

			// 支付功能
			authed.POST("pay", api.OrderPay)

		}
	}
	return r
}
