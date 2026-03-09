package main

import "fmt"

type User struct {
	NickName string
	Password string
	RealName RealName
}
type RealName struct {
	Name  string
	Age   *int
	Tels  []string
	Email map[string]string
}

func main() {
	age := 18
	var u = &User{
		NickName: "sxy",
		Password: "123456",
		RealName: RealName{
			Name: "RayeLi",
			Age:  &age,
			Tels: []string{"tel1", "tel2", "tel3"},
			Email: map[string]string{
				"email1": "email1",
				"email2": "email2",
			},
		},
	}
	fmt.Println(u)
	fmt.Println(*u.RealName.Age)
	fmt.Println("*******************")
	u1 := *u
	//age1 := 1999
	u1.NickName = "sxy1"
	u1.RealName.Name = "RayeLi11"
	*u1.RealName.Age = 19999
	u1.RealName.Tels[0] = "tel11"
	u1.RealName.Email["email1"] = "email11"
	fmt.Println(u)
	fmt.Println(*u.RealName.Age)

}
