package yaml

import (
	"fmt"

	"sxy/demo/model"

	"github.com/spf13/viper"
)

func LoadYaml01(name string, types string, path string) {
	viper.SetConfigName(name)
	viper.SetConfigType(types)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// fmt.Println("All settings: ", viper.AllSettings())
	// fmt.Println("person: ", viper.Get("person"))
	var person model.Person

	err = viper.UnmarshalKey("person", &person)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	// 打印个人信息
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Gender: %s\n", person.Gender)
	fmt.Printf("Email: %s\n", person.Email)
	fmt.Printf("Phone: %s\n", person.Phone)
	fmt.Printf("Address: %s, %s, %s, %s\n", person.Address.Street, person.Address.City, person.Address.State, person.Address.Zip)
	fmt.Printf("Education:\n")
	for _, edu := range person.Education {
		fmt.Printf("  - %s in %s from %s, Year: %d\n", edu.Degree, edu.Major, edu.School, edu.Year)
	}
	fmt.Printf("Experience:\n")
	for _, exp := range person.Experience {
		fmt.Printf("  - %s as %s from %d to %d\n", exp.Company, exp.Position, exp.StartYear, exp.EndYear)
	}
	fmt.Printf("Skills:\n")
	for k, v := range person.Skills {
		fmt.Printf("  - %s: %s\n", k, v)
	}
}
