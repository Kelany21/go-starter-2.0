package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang-ddd-starter/app/validations/user"
	"golang-ddd-starter/domian/models"
	"golang-ddd-starter/domian/repos"
	"golang-ddd-starter/helpers"
	"golang-ddd-starter/infrastructure"
)

type UserApp struct {
	repo repos.UserRepository
}

var UserApplication *UserApp = nil

func NewUserApplication(db *gorm.DB) {
	UserApplication = &UserApp{infrastructure.NewUserRepository(db)}
}

//UserApp implements the UserApplication interface
var _ ApplicationInterface = &UserApp{}

func (a *UserApp) Show(g *gin.Context) {
	// find this row or return 404
	row, find := a.repo.FindOrFail(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneGetItem(g), row)
}

func (a *UserApp) List(g *gin.Context) {
	rows, err := a.repo.GetAll()
	if err != nil {
		helpers.ReturnForbidden(g, err.Error())
		return
	}
	helpers.OkResponse(g, helpers.DoneGetItem(g), rows)
}

func (a *UserApp) Paginate(g *gin.Context) {
	paginator, err := a.repo.Paginate(&helpers.Param{
		Page:    Page(g),
		Limit:   Limit(g),
		OrderBy: Order(g, "id desc"),
		Filters: filter(g),
		Preload: preload(),
		ShowSQL: true,
	})
	if err != nil {
		helpers.ReturnForbidden(g, err.Error())
		return
	}
	// return response
	helpers.OkResponseWithPaging(g, helpers.DoneGetAllItems(g), paginator)
}

func (a *UserApp) Create(g *gin.Context) {
	// init struct to validate request
	row := new(models.User)
	// check if request valid
	if !user.Store(g, row) {
		return
	}
	/// check if this email exists
	count, err := infrastructure.RecordCount(models.User{}, "email = ? ", row.Email)
	if err != nil {
		helpers.ReturnForbidden(g, err.Error())
		return
	}
	if count > 0 {
		helpers.ReturnDuplicateData(g, "email")
		return
	}
	// create new row
	err = a.repo.Create(row)
	if err != nil {
		helpers.ReturnForbidden(g, err.Error())
		return
	}
	//now return row data after transformers
	helpers.OkResponse(g, helpers.DoneCreateItem(g), row)
}

func (a *UserApp) Update(g *gin.Context) {
	// init struct to validate request
	row := new(models.User)
	// check if request valid
	if !user.Update(g, row) {
		return
	}
	// find this row or return 404
	oldRow, find := a.repo.FindOrFail(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	/// check if this email exists
	count, err := infrastructure.RecordCount(models.User{}, "email = ? AND email != ?", row.Email, oldRow.Email)
	if err != nil {
		helpers.ReturnForbidden(g, err.Error())
		return
	}
	if count > 0 {
		helpers.ReturnDuplicateData(g, "email")
		return
	}
	/// update allow columns
	err = a.repo.Update(row, oldRow.ID)
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneUpdate(g), row)
}

func (a *UserApp) Delete(g *gin.Context) {
	// find this row or return 404
	row, find := a.repo.FindOrFail(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	err := a.repo.Delete(row.ID)
	if err != nil {
		helpers.ReturnForbidden(g, err.Error())
		return
	}
	// now return row data after transformers
	helpers.OkResponseWithOutData(g, helpers.DoneDelete(g))
}

/**
* filter module with some columns
 */
func filter(g *gin.Context) []string {
	var filter []string
	if g.Query("block") != "" {
		filter = append(filter, "block = "+g.Query("block"))
	}
	if g.Query("name") != "" {
		filter = append(filter, `name like "%`+g.Query("name")+`%"`)
	}
	if g.Query("email") != "" {
		filter = append(filter, `email like "%`+g.Query("email")+`%"`)
	}
	if g.Query("role") != "" {
		filter = append(filter, `role like "%`+g.Query("role")+`%"`)
	}
	return filter
}

/**
* preload module with some preload conditions
 */
func preload() []string {
	return []string{}
}
