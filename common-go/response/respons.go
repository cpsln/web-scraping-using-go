package response


type ProductDetails struct {
	Id						int		`json:"id"`
	ProductName				string	`json:"product_name"`
	ProductImageUrl			string	`json:"image_url"`
	ProductDescription		string	`json:"description"`
	ProductPrice			string	`json:"price"`
	NumberOfReviews			string	`json:"total_review"`
}

type Data struct {
    Url  		string `json:"url"`
	Product 	ProductDetails `json:"product"`
}


type ErrorRespons struct {
	StatusCode	int64	`json:"status_code"`
	Message  	string	`json:"message"`
}

