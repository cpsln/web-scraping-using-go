package main

import (
    "fmt"
    "log"
    "strings"
    "reflect"
    "strconv"
    "net/http"
    "io/ioutil"
	"encoding/json"
    "github.com/gocolly/colly"
    "github.com/beego/beego/v2/client/orm"
    db "common-go/dbconnection"
    model "common-go/model"
    res "common-go/response"
)


// type DetalsOfProduct struct {
//     Url string `json:"url"`
//     Product ProductD `json:"product"`
// }

func scrapData(w http.ResponseWriter, r *http.Request) {
    err_res := res.ErrorRespons{}

    if r.Method == "POST" {

        var product model.ProductDetails
        c := colly.NewCollector()

        url := res.Data{}
        if body, err := ioutil.ReadAll(r.Body); err == nil {
            if err := json.Unmarshal([]byte(body), &url); err != nil{
                fmt.Println("Unmarshal Error!...",err)
                err_res.StatusCode = 404
                err_res.Message = "Can not able to unmarshal"
                json.NewEncoder(w).Encode(err_res)
                return
            }
        }
        
        product.Url = url.Url

        c.OnHTML("#ppd", func(e *colly.HTMLElement) {

            product.ProductName = e.ChildText("#productTitle")
            product.ProductImageUrl = e.ChildAttr("#landingImage", "src")
            var text string
            var count = 0
            e.ForEach("div#featurebullets_feature_div > div#feature-bullets > ul.a-unordered-list> li > span.a-list-item", func(_ int, elem *colly.HTMLElement) {
                if count > 0 {
                    text += strings.TrimPrefix(strings.TrimSuffix(elem.Text, "\n\n"), "\n") + ". "
                } 
                count +=1
            })
            product.ProductDescription = text
            price := strings.Split(e.ChildText("span#edition_0_price > span.a-size-mini"),"$")
            product.ProductPrice = "$"+price[len(price)-1]
            reviws := strings.Split(e.ChildText("#acrCustomerReviewLink > span#acrCustomerReviewText")," ")
            product.NumberOfReviews = reviws[0]
        })
        
	    c.OnScraped(func(r *colly.Response) {
            err, id := model.InsertProductDetails(&product)
            if err != nil {
                fmt.Println("Insert Error!...",err)
                err_res.StatusCode = 404
                err_res.Message = "Details Not save"
                json.NewEncoder(w).Encode(err_res)
                return
            } else {
                fmt.Println("Data Save", id)
            }
            
            err_res.StatusCode = 200
            err_res.Message = fmt.Sprintf("Data saved, product id is: %d ", id);
            json.NewEncoder(w).Encode(err_res)
        })
        c.Visit(url.Url)
        return
	} 

    err_res.StatusCode = 404
    err_res.Message = r.Method + "Method not allow"
    json.NewEncoder(w).Encode(err_res)
    return
}

func GetProductsDetails (w http.ResponseWriter, r *http.Request){
    err_res := res.ErrorRespons{}
    if r.Method == "GET" {

        query := r.URL.Query()
        filters, err := strconv.ParseInt(query.Get("pid"),10,64)

        if err != nil {
            fmt.Println("Parse Error!...",err)
            err_res.StatusCode = 404
            err_res.Message = "Can't parse to int product id:"+query.Get("pid")
            json.NewEncoder(w).Encode(err_res)
            return
        }
        fmt.Println(reflect.TypeOf(filters))

        p, err := model.GetProducts(filters)

        if err != nil {
            fmt.Println("Error!...",err)
            err_res.StatusCode = 404
            err_res.Message = "Can't read product id:"+query.Get("pid")
            json.NewEncoder(w).Encode(err_res)
            return
        } 
        
        product := res.Data{}

        product.Url                        = p.Url
        product.Product.Id                 = p.Id
        product.Product.ProductName        = p.ProductName
        product.Product.ProductImageUrl    = p.ProductImageUrl
        product.Product.ProductDescription = p.ProductDescription
        product.Product.ProductPrice       = p.ProductPrice
        product.Product.NumberOfReviews    = p.NumberOfReviews

        json.NewEncoder(w).Encode(product)
        return
    }

    err_res.StatusCode = 404
    err_res.Message = r.Method + "Method not allow"
    json.NewEncoder(w).Encode(err_res)
    return
}

func UpdateProduct (w http.ResponseWriter, r *http.Request){

    err_res := res.ErrorRespons{}

    if r.Method == "PUT" {
        var data res.Data
        // var product model.ProductDetails

        maps := make(map[string]interface{})

        if body, err := ioutil.ReadAll(r.Body); err == nil {
            if err := json.Unmarshal([]byte(body), &data); err != nil{
                fmt.Println("Unmarshal Error!...",err)
                err_res.StatusCode = 404
                err_res.Message = "Can not able to unmarshal"
                json.NewEncoder(w).Encode(err_res)
                return
            }
        }

        cond := orm.NewCondition().And("id", data.Product.Id)

        maps["product_name"]    = data.Product.ProductName
        maps["image_url"]       = data.Product.ProductImageUrl
        maps["description"]     = data.Product.ProductDescription
        maps["price"]           = data.Product.ProductPrice
        maps["total_review"]    = data.Product.NumberOfReviews
        maps["url"]             = data.Url

        err := model.UpdateProduct(cond, maps)

        if err != nil {
            fmt.Println("Update Error!...",err)
            err_res.StatusCode = 404
            err_res.Message = "Can not update"
            json.NewEncoder(w).Encode(err_res)
            return
        }

        err_res.StatusCode = 200
        err_res.Message = "Updated"
        json.NewEncoder(w).Encode(err_res)
        return
    }

    err_res.StatusCode = 404
    err_res.Message = r.Method + " Method not allow"
    json.NewEncoder(w).Encode(err_res)
    return
}


func handleRequests() {
    http.HandleFunc("/scrap", scrapData)
	http.HandleFunc("/get", GetProductsDetails)
    http.HandleFunc("/update", UpdateProduct)
    // err := db.ConnectDatabase()

    fmt.Println("running on: http://localhost:8000/")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    db.ConnectDatabase()
    // db.Close()
    handleRequests()
}


