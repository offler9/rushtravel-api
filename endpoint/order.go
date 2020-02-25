package endpoint

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rush-travel/util"

	//"github.com/rs/xid"
	"github.com/rush-travel/config"
	"github.com/rush-travel/model"
)

type Coordinate struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}
type Coordinate1 struct {
	Latitude  float64 `json:"lati"`
	Longitude float64 `json:"long"`
}
type Order struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time     `sql:"index"`
	OrderId        string         `gorm:"column:order_id" json:"orderId"`
	Username       string         `gorm:"column:username" json:"username"`
	PickUpLoc      pq.StringArray `gorm:"type:string;column:pick_up_loc"`
	DropOffLoc     pq.StringArray `gorm:"type:string;column:drop_off_loc"`
	Date           string         `gorm:"column:date" json:"date"`
	Days           string         `gorm:"column:days" json:"days"`
	NoOfPassangers int            `gorm:"column:no_off_passangers" json:"no_off_passangers"`
	Price          int            `gorm:"column:price" json:"price"`
	VehicleType    string         `gorm:"column:vehicle_type" json:"vehicle_type"`
	Status string `json:"status"`
	Note string `json:"note"`
}

func CreateOrder(c *gin.Context) {
	var ord Order
	var buffer bytes.Buffer
	var buf bytes.Buffer
	//var coord []Coordinate
	//cor := Coordinate{}
	var pick []byte
	var drop []byte
	var err error
	//var coord1 []Coordinate1

	lat, _ := strconv.ParseFloat(c.PostForm("lat"), 64)
	lng, _ := strconv.ParseFloat(c.PostForm("lng"), 64)
	lati, _ := strconv.ParseFloat(c.PostForm("lati"), 64)
	long, _ := strconv.ParseFloat(c.PostForm("long"), 64)
	Date, _ := time.Parse("2006-01-02", c.PostForm("date"))
	NoOffPass, _ := strconv.Atoi(c.PostForm("no_off_passangers"))
	Price, _ := strconv.Atoi(c.PostForm("price"))

	//
	pickUpLoc := map[string]interface{}{
		"Latitude":  lat,
		"Longitude": lng,
	}

	for _, pickUpLoc := range pickUpLoc {
		pick, err = json.Marshal(pickUpLoc)
		fmt.Println("pick", pick)
		if err != nil {
			fmt.Println(err)
		}
		buffer.WriteString(string(pick) + " ")
	}
	s := strings.TrimSpace(buffer.String())
	fmt.Println("s",s)
	ss := []string{s}
	// json.Unmarshal([]byte(s), &cor)
	fmt.Println("s", ss)
	//
	dropOffLoc := map[string]interface{}{
		"Latitude":  lati,
		"Longitude": long,
	}
	fmt.Println(" d ",dropOffLoc)
	for _, dropOffLoc := range dropOffLoc {
		drop, err = json.Marshal(dropOffLoc)
		if err != nil {
			fmt.Println(err)
		}
		buf.WriteString(string(drop) + " ")
	}

	d := strings.TrimSpace(buf.String())
	dd := []string{d}

	fmt.Println("dd", d)
	//drop := []string{d}
	n := 5
	unique := make([]byte, n)

	if _, err := rand.Read(unique); err != nil {
		panic(err)
	}

	ID := fmt.Sprintf("%X", unique)
	//fmt.Println(ID)

	createorder := Order{
		OrderId:        ID,
		Username:       "tes",
		PickUpLoc:      ss,
		DropOffLoc:     dd,
		Date:           Date.Format("2006-01-01"),
		Days:           c.PostForm("days"),
		Price:          Price,
		NoOfPassangers: NoOffPass,
		VehicleType:    c.PostForm("vehicle_type"),
		Status: "waiting",
	}
	err = config.DB.Model(&ord).Save(&createorder).Error
	if err != nil {
		fmt.Println("err", err)
		util.CallServerError(c, "failed", nil)
		return
	}
	util.CallSuccessOK(c, "Successfully Add Order", nil)
}

func FetchOrder(c *gin.Context) {
	var order []Order
	var order1 []Order

	config.DB.Model(&order).Find(&order)
	if len(order) < 0 {
		util.CallErrorNotFound(c, "No Order Data", nil)
	}

	for _, item := range order {
		order1 = append(order1, Order{
			ID:             item.ID,
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
			DeletedAt:      item.DeletedAt,
			OrderId:        item.OrderId,
			Username:       item.Username,
			PickUpLoc:      item.PickUpLoc,
			DropOffLoc:     item.DropOffLoc,
			Date:           item.Date,
			Days:           item.Days,
			Price:          item.Price,
			NoOfPassangers: item.NoOfPassangers,
			VehicleType:    item.VehicleType,
		})
	}
	util.CallSuccessOK(c, "Fetch All Orders Data ", order1)
}

func UpdateOrderStatus(c *gin.Context){
	var order model.Order
	id := c.Param("id")
	note := c.PostForm("note")
	status :=c.PostForm("status")

	update := model.Order{
		Note: note,
		Status: status,
	}
	err:= config.DB.Model(&order).Where("ID = ? ", id).Update(&update).Error
	if err != nil {
		util.CallServerError(c,"failed when to try update order status", nil)
	}
	util.CallSuccessOK(c,"Successfully Update Order Status",nil)
}