package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["opms/controllers/order:CustomerController"] = append(beego.GlobalControllerRouter["opms/controllers/order:CustomerController"],
        beego.ControllerComments{
            Method: "AddCustomerPost",
            Router: `/customer/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:CustomerController"] = append(beego.GlobalControllerRouter["opms/controllers/order:CustomerController"],
        beego.ControllerComments{
            Method: "AddCustomerGet",
            Router: `/customer/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:CustomerController"] = append(beego.GlobalControllerRouter["opms/controllers/order:CustomerController"],
        beego.ControllerComments{
            Method: "AjaxStatus",
            Router: `/customer/ajax/status`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:CustomerController"] = append(beego.GlobalControllerRouter["opms/controllers/order:CustomerController"],
        beego.ControllerComments{
            Method: "EditCustomerPost",
            Router: `/customer/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:CustomerController"] = append(beego.GlobalControllerRouter["opms/controllers/order:CustomerController"],
        beego.ControllerComments{
            Method: "EditCustomerGet",
            Router: `/customer/edit/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:CustomerController"] = append(beego.GlobalControllerRouter["opms/controllers/order:CustomerController"],
        beego.ControllerComments{
            Method: "ListAllCustomers",
            Router: `/customer/manage`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:CustomerController"] = append(beego.GlobalControllerRouter["opms/controllers/order:CustomerController"],
        beego.ControllerComments{
            Method: "RemoveOrder",
            Router: `/customer/remove/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:CustomerController"] = append(beego.GlobalControllerRouter["opms/controllers/order:CustomerController"],
        beego.ControllerComments{
            Method: "ShowOrder",
            Router: `/customer/showOrder`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:OrderController"] = append(beego.GlobalControllerRouter["opms/controllers/order:OrderController"],
        beego.ControllerComments{
            Method: "AddOrder",
            Router: `/order/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:OrderController"] = append(beego.GlobalControllerRouter["opms/controllers/order:OrderController"],
        beego.ControllerComments{
            Method: "AddOrderPost",
            Router: `/order/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:OrderController"] = append(beego.GlobalControllerRouter["opms/controllers/order:OrderController"],
        beego.ControllerComments{
            Method: "AjaxStatus",
            Router: `/order/ajax/status`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:OrderController"] = append(beego.GlobalControllerRouter["opms/controllers/order:OrderController"],
        beego.ControllerComments{
            Method: "EditOrder",
            Router: `/order/edit/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:OrderController"] = append(beego.GlobalControllerRouter["opms/controllers/order:OrderController"],
        beego.ControllerComments{
            Method: "ListAllOrders",
            Router: `/order/manage`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:OrderController"] = append(beego.GlobalControllerRouter["opms/controllers/order:OrderController"],
        beego.ControllerComments{
            Method: "RemoveOrder",
            Router: `/order/remove/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["opms/controllers/order:OrderController"] = append(beego.GlobalControllerRouter["opms/controllers/order:OrderController"],
        beego.ControllerComments{
            Method: "ShowOrder",
            Router: `/order/showOrder`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
