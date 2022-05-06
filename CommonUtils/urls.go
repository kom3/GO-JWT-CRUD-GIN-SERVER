package CommonUtils

import (
	"gomodules/APIS/CRUDAPIS"

	"github.com/gin-gonic/gin"
)

// New urls and APIs can be mapped below
func routerCollector(router *gin.Engine) {
	router.POST("/samplepost", CRUDAPIS.SamplePostAPI)
	router.GET("/sampleget", CRUDAPIS.SampleGETAPI)
	router.GET("/getallemployees", CRUDAPIS.GetAllEmployeesList)
	router.GET("/getemployeebyid", CRUDAPIS.GetEmployeeById)
	router.GET("/getemployeebyid/:empid", CRUDAPIS.GetEmployeeById)
	router.GET("/addemployee", CRUDAPIS.AddEmployee)
	router.GET("/addemployees", CRUDAPIS.AddEmployees)
	router.GET("/updateemployeebyid", CRUDAPIS.UpdateEmployeeById)
	router.GET("/updateallemployees", CRUDAPIS.UpdateAllEmployees)
	router.GET("/deleteemployeebyid", CRUDAPIS.DeleteEmployeeById)
	router.GET("/deleteallemployees", CRUDAPIS.DeleteAllEmployees)
}
