package model


import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type ProductDetails struct {
	Id						int		`json:"id" orm:"column(id);pk"`
	ProductName				string	`json:"product_name" orm:"column(product_name)"`
	ProductImageUrl			string	`json:"image_url" orm:"column(image_url)"`
	ProductDescription		string	`json:"description" orm:"column(description)"`
	ProductPrice			string	`json:"price" orm:"column(price)"`
	NumberOfReviews			string	`json:"total_review" orm:"column(total_review)"`
	Url						string  `json:"url" orm:"column(url)"`
}


func (p *ProductDetails) TableName() string {
	return "product_details"
}

func init() {
	orm.RegisterModel(new(ProductDetails))
}


func InsertProductDetails(product *ProductDetails) (error, int64){
	o := orm.NewOrm()
	id, err := o.Insert(product)

	if err != nil {
		fmt.Println("Error!...",err)
		return err, -1
	}

	return nil, id
}

func GetProducts(id int64) (* ProductDetails, error) {
	var product ProductDetails

	o := orm.NewOrm()
	cond := orm.NewCondition()
	condition := cond.And("id",id)
	err := o.QueryTable("product_details").SetCond(condition).RelatedSel().One(&product)
	if err != nil {
		fmt.Println("failed ->",err)
		return nil, err
	}
	return &product, nil
}

func UpdateProduct(cond *orm.Condition, maps map[string]interface{}) error {
	o :=orm.NewOrm()
	_, err := o.QueryTable("product_details").SetCond(cond).Update(maps)

	if err != nil {
		return err
	}
	return nil
}