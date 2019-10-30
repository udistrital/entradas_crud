// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/entradas_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/entrada_elemento",
			beego.NSInclude(
				&controllers.EntradaElementoController{},
			),
		),

		beego.NSNamespace("/tipo_entrada",
			beego.NSInclude(
				&controllers.TipoEntradaController{},
			),
		),

		beego.NSNamespace("/soporte_entrada",
			beego.NSInclude(
				&controllers.SoporteEntradaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
