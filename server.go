package main

//import modules

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	validation "github.com/go-ozzo/ozzo-validation" //specific data validation
	"github.com/gofiber/fiber/v2"                   //web framework
	"github.com/gofiber/template/html"
	"github.com/lithammer/shortuuid" //to generate ids
)

//define base customer struct
type customer struct {
	User_id string
	First_Name string
	Last_Name string
	Birth_Date string
	Gender string
	Email string
	Address string
 }

//define handlers below
func indexHandler(c *fiber.Ctx, db *sql.DB) error {

   rows, err := db.Query("SELECT id, first_name, last_name, birth_date, gender, email, address FROM customers")
   defer rows.Close()
   if err != nil {
       log.Fatalln(err)
       c.JSON("An error occured")
   }
       type row struct {
			User_id string
            First_name  string
            Last_name string
			Birth_date string
			Gender string
			Email string
			Address string
    }
	customerList := []row{}
	    for rows.Next() {
            var r row
            err = rows.Scan(&r.User_id, &r.First_name, &r.Last_name, &r.Birth_date, &r.Gender, &r.Email, &r.Address)
            if err != nil {
				log.Fatalf("Scan: %v", err)
            }
            customerList = append(customerList, r)
    }
   return c.Render("index", fiber.Map{
       "Customers": customerList,
   })
}

func createPageHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var customers []string
	rows, err := db.Query("SELECT * FROM customers")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		customers = append(customers, res)
	}
	return c.Render("create", fiber.Map{
		"Customers": customers,
	})
 }

func editPageHandler(c *fiber.Ctx, db *sql.DB, selectedUser string) error {
	rows, err := db.Query("SELECT * FROM customers WHERE id = ($1)", selectedUser)
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	type row struct {
		User_id string
		First_name  string
		Last_name string
		Birth_date string
		Gender string
		Email string
		Address string
}
customerList := []row{}
	for rows.Next() {
		var r row
		err = rows.Scan(&r.User_id, &r.First_name, &r.Last_name, &r.Birth_date, &r.Gender, &r.Email, &r.Address)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		customerList = append(customerList, r)
}
	return c.Render("edit", fiber.Map{
		"Customers": customerList,
	})
 }

 func searchPageHandler(c *fiber.Ctx, db *sql.DB) error {

	rows, err := db.Query("SELECT id, first_name, last_name, birth_date, gender, email, address FROM customers")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
		type row struct {
			User_id string
			First_name  string
			Last_name string
			Birth_date string
			Gender string
			Email string
			Address string
	}
	customerList := []row{}
		for rows.Next() {
			var r row
			err = rows.Scan(&r.User_id, &r.First_name, &r.Last_name, &r.Birth_date, &r.Gender, &r.Email, &r.Address)
			if err != nil {
				log.Fatalf("Scan: %v", err)
			}
			customerList = append(customerList, r)
	}

	return c.Render("search", fiber.Map{
		"Customers": customerList,
	})
 }

 func showPageHandler(c *fiber.Ctx, db *sql.DB, selectedUser string) error {

	rows, err := db.Query("SELECT * FROM customers WHERE id = ($1)", selectedUser)
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	type row struct {
		User_id string
		First_name  string
		Last_name string
		Birth_date string
		Gender string
		Email string
		Address string
}
customerList := []row{}
	for rows.Next() {
		var r row
		err = rows.Scan(&r.User_id, &r.First_name, &r.Last_name, &r.Birth_date, &r.Gender, &r.Email, &r.Address)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		customerList = append(customerList, r)
}
	return c.Render("show", fiber.Map{
		"Customers": customerList,
	})
 }

func createCustomerHandler(c *fiber.Ctx, db *sql.DB) error {
   newCustomer := customer{}
   if err := c.BodyParser(&newCustomer); err != nil {
       log.Printf("An error occured: %v", err)
       return c.SendString(err.Error())
   }
   
   fmt.Printf("New customer data: %v", newCustomer)
   err := newCustomer.Validate()
	
	if err != nil {
		fmt.Println(err)
		return c.Render("notice", fiber.Map{
			"errorMessage": err,
		})
	}

	customerId := shortuuid.New()
	
   if newCustomer.First_Name != "" {
       _, err := db.Exec("INSERT into customers VALUES ($1, $2, $3, $4, $5, $6, $7)", customerId, newCustomer.First_Name, newCustomer.Last_Name, newCustomer.Birth_Date, newCustomer.Gender, newCustomer.Email, newCustomer.Address)
       if err != nil {
           log.Fatalf("An error occured while executing query: %v", err)
       }
   }

   return c.Redirect("/")
}

func editCustomerHandler(c *fiber.Ctx, db *sql.DB) error {
	
	 newCustomer := customer{}
	if err := c.BodyParser(&newCustomer); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}

	fmt.Printf("New customer data: %v", newCustomer)
	validationErr := newCustomer.Validate()
	 
	 if validationErr != nil {
		 fmt.Println(validationErr)
		 return c.Render("notice", fiber.Map{
			 "errorMessage": validationErr,
		 })
	 }

	if newCustomer.First_Name != "" {
		_, err := db.Exec("UPDATE customers SET first_name = $2, last_name = $3, birth_date = $4, gender = $5, email = $6, address = $7 WHERE id = $1", newCustomer.User_id, newCustomer.First_Name, newCustomer.Last_Name, newCustomer.Birth_Date, newCustomer.Gender, newCustomer.Email, newCustomer.Address)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}
 
	return c.Redirect("/search")
 }

func deleteCustomerHandler(c *fiber.Ctx, db *sql.DB) error {
   userId := c.Query("item")
   fmt.Printf("Customer to delete: %v", userId)
   db.Exec("DELETE from customers WHERE id=$1", userId)
   return c.SendString("deleted")
}

func (a customer) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.First_Name, validation.Required, validation.Length(0, 100)),
		validation.Field(&a.Last_Name, validation.Required, validation.Length(0, 100)),
		validation.Field(&a.Birth_Date, validation.Required),
		validation.Field(&a.Gender, validation.Required),
		validation.Field(&a.Email, validation.Required),
	)
}

func main() {
	//perform database configs
   connStr := "postgresql://postgres:wern234nfdsaCk%23tyu@localhost/crud_app_data?sslmode=disable"

   db, err := sql.Open("postgres", connStr)
   if err != nil {
       log.Fatal(err)
   }
   //define engine
   engine := html.New("./views", ".html")
   //define fiber app
   app := fiber.New(fiber.Config{
       Views: engine,
   })
   //define routes
   app.Get("/", func(c *fiber.Ctx) error {
       return indexHandler(c, db)
   })

   app.Get("/hello", func(c *fiber.Ctx) error {

	return c.SendString("Hello, World!")
  })

   app.Get("/create", func(c *fiber.Ctx) error {
		return createPageHandler(c, db);
   })

   app.Get("/edit/user/:userid", func(c *fiber.Ctx) error {
	selectedUser := c.Params("userid")
	return editPageHandler(c, db, selectedUser);
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		return searchPageHandler(c, db);
	})

	app.Get("/show/user/:userid", func(c *fiber.Ctx) error {
		selectedUser := c.Params("userid")
		return showPageHandler(c, db, selectedUser);
	})

   app.Post("/createCustomer", func(c *fiber.Ctx) error {
       return createCustomerHandler(c, db)
   })

   app.Post("/editCustomer", func(c *fiber.Ctx) error {
	return editCustomerHandler(c, db)
	})

   app.Delete("/delete", func(c *fiber.Ctx) error {
       return deleteCustomerHandler(c, db)
   })
   //define port
   port := os.Getenv("PORT")
   if port == "" {
       port = "3001"
   }
   app.Static("/", "./public")
   log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}