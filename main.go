package main

import "gofr.dev/pkg/gofr"

type Students struct {
	ID           int    `json:"id"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Age          int    `json:"age"`
	Class        int    `json:"class"`
	Gender       string `json:"gender"`
	Address      string `json:"address"`
	Phone_number int    `json:"phone_number"`
}

func main() {
	// initialise gofr object
	app := gofr.New()

	app.GET("/students", func(ctx *gofr.Context) (interface{}, error) {
		var Studs []Students

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM Student_details")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var Stud Students
			if err := rows.Scan(&Stud.ID, &Stud.First_name, &Stud.Last_name, &Stud.Age, &Stud.Class, &Stud.Gender, &Stud.Address, &Stud.Phone_number); err != nil {
				return nil, err
			}

			Studs = append(Studs, Stud)
		}

		return Studs, nil
	})

	app.POST("/students/{first_name}/{last_name}/{age}/{class}/{gender}/{address}/{phone_number}", func(ctx *gofr.Context) (interface{}, error) {
		first_name := ctx.PathParam("first_name")
		last_name := ctx.PathParam("last_name")
		age := ctx.PathParam("age")
		class := ctx.PathParam("class")
		gender := ctx.PathParam("gender")
		address := ctx.PathParam("address")
		phone_number := ctx.PathParam("phone_number")

		_, err := ctx.DB().ExecContext(ctx,
			"INSERT INTO Student_details (first_name,last_name,age,class,gender,address,phone_number) VALUES (?,?,?,?,?,?,?)",
			first_name, last_name, age, class, gender, address, phone_number)

		return nil, err
	})

	app.PUT("/students/{id}/{first_name}/{last_name}/{age}/{class}/{gender}/{address}/{phone_number}", func(ctx *gofr.Context) (interface{}, error) {
		first_name := ctx.PathParam("first_name")
		last_name := ctx.PathParam("last_name")
		age := ctx.PathParam("age")
		class := ctx.PathParam("class")
		gender := ctx.PathParam("gender")
		address := ctx.PathParam("address")
		phone_number := ctx.PathParam("phone_number")

		_, err := ctx.DB().ExecContext(ctx,
			"UPDATE Student_details SET first_name=?, last_name=?,age=?, class=?, gender=?, address=? ,phone_number=? WHERE id=?",
			first_name, last_name, age, class, gender, address, phone_number)

		return nil, err
	})

	app.DELETE("/students/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")

		_, err := ctx.DB().ExecContext(ctx,
			"DELETE FROM Student_details WHERE id=?",
			id)

		return nil, err
	})

	app.Start()
}
