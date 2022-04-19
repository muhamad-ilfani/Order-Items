# Get Started
#
clone : git clone https://github.com/muhamad-ilfani/Order-Items.git

running : go run main.go
# Setting Environment

const (

	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "April2022"
	dbname   = "orders_by"
)

# Model

type Order struct {

	gorm.Model
	Customer_name string    `json:"customer_name"`
	Order_id      uint      `json:"order_id"`
	Ordered_at    time.Time `json:"ordered_at"`
	Item_user     []Item    `json:"item_user" gorm:"foreignKey:ID_ref;references:ID"`
}

type Item struct {

	gorm.Model
	Item_code   string `json:"item_code" gorm:"primaryKey"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	Order_id    uint   `json:"order_id"`
	ID_ref      uint
}


# API LIST

### PORT ":8080"
#

Create order : http://localhost:[PORT]/orders method : POST

Read order : http://localhost:[PORT]/orders method : GET

Update Order/ Add Items : http://localhost:[PORT]/orders/:id method : PUT

Delete Order : http://localhost:[PORT]/orders/:id method : DELETE