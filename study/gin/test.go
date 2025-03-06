package gin

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func StartGin() {

	r := gin.Default()
	r.POST("/f", func(c *gin.Context) {
		//file, _ := c.FormFile("file2")
		c.Redirect(http.StatusFound, "/add")
		c.JSON(200, gin.H{
			"message": "hello",
			//"ext":     path.Ext(file.Filename),
		})
		//time.Sleep(3 * time.Second)
	})
	r.POST("/add", ADD)
	r.Run(":8899")

}
func Parse01() {
	u, err := url.Parse("https://example.org/?a=1&a=2&b=&c=3&&&&")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	fmt.Println(q["a"])
	fmt.Println("B: ", q.Get("B"))
	fmt.Println(q.Get(""))
	fmt.Println(".....")
	q.Add("sxy", "sss")
	q.Add("name ns", "lcx")
	fmt.Println(q.Encode())
	fmt.Println(u.String())
}

//	func uuidTest() {
//		u1 := uuid.NewV4()
//		fmt.Println(u1)
//		u2, _ := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
//		fmt.Println("u2: ", u2)
//	}

type Arr struct {
	A int `json:"a"  form:"a"`
	B int `json:"b" form:"b"`
}

func ADD(c *gin.Context) {
	var s Arr
	if err := c.Bind(&s); err != nil {
		fmt.Println("sss")
		log.Fatal(err)
	}
	fmt.Println(s.A, s.B)
	res := s.A + s.B
	c.JSON(200, gin.H{
		"result": res,
	})
}
