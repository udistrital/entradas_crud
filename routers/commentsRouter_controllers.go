package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:EntradaElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:EntradaElementoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:EntradaElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:EntradaElementoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:EntradaElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:EntradaElementoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:EntradaElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:EntradaElementoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:EntradaElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:EntradaElementoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:SoporteEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:SoporteEntradaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:SoporteEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:SoporteEntradaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:SoporteEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:SoporteEntradaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:SoporteEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:SoporteEntradaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:SoporteEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:SoporteEntradaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:TipoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:TipoEntradaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:TipoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:TipoEntradaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:TipoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:TipoEntradaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:TipoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:TipoEntradaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:TipoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/entradas_crud/controllers:TipoEntradaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
