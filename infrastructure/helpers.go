package infrastructure

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	"golang-ddd-starter/helpers"
	"math"
	"os"
	"strconv"
	"strings"
)

/***
* truncate tables
 */
func DbTruncate(tableName ...string) {
	for _, table := range tableName {
		DB.Exec("TRUNCATE " + table)
	}
}

/**
* this function get struct and return with only
* Available column that allow to updated depend on FillAbleColumn function
* this for security
* map struct to update
 */
func UpdateOnlyAllowColumns(structNeedToMap interface{}, fillAble []string) interface{} {
	row := structs.Map(structNeedToMap)
	var data = make(map[string]interface{})
	for _, value := range fillAble {
		if row[strings.Title(value)] != "" {
			data[value] = row[strcase.ToCamel(value)]
		}
	}
	return data
}

/**
* add preload dynamic
* this will allow to add more than one preload
 */
func PreloadD(db *gorm.DB, preload []string) *gorm.DB {
	if len(preload) > 0 {
		for _, p := range preload {
			db = db.Preload(p)
		}
		return db
	}
	return db
}

func RecordCount(model interface{}, query interface{}, args ...interface{}) (count int, err error) {
	err = DB.Model(model).Where(query, args...).Count(&count).Error
	return
}

// Paging 分页
func Paging(p *helpers.Param, result interface{}) (*helpers.Paginator, error) {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}

	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	if len(p.Filters) > 0 {
		for _, o := range p.Filters {
			db = db.Where(o)
		}
	}

	if len(p.Preload) > 0 {
		db = PreloadD(db, p.Preload)
	}

	done := make(chan bool, 1)
	var paginator helpers.Paginator
	var count int
	//var resultCount int
	var offset int

	go countRecords(db, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	//err := db.Model(result).Limit(p.Limit).Offset(offset).Count(&resultCount).Error
	//if err != nil {
	//	return nil, err
	//}

	err := db.Limit(p.Limit).Offset(offset).Find(result).Error
	if err != nil {
		return nil, err
	}
	<-done

	paginator.TotalRecord = count
	//paginator.Records = result
	paginator.Page = p.Page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator, nil
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}

func Page(g *gin.Context) int {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	return page
}

func Limit(g *gin.Context) int {
	page, _ := strconv.Atoi(g.DefaultQuery("limit", os.Getenv("LIMIT")))
	return page
}

func Order(g *gin.Context, orders ...string) []string {
	var o []string

	if g.Query("order") != "" {

		orders := strings.SplitN(g.Query("order"), "|", -1)

		for _, order := range orders {
			o = append(o, order)
		}
		return o

	} else {

		for _, orderBy := range orders {
			o = append(o, orderBy)
		}
		return o

	}
}

//func Order(order ...string)[]string   {
//	var o []string
//	for _ , orderBy := range order{
//		o = append(o , orderBy)
//	}
//	return o
//}
