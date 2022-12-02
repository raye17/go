package main

type kelvin float64

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v' K\n", k)
		time.Sleep(time.Second)
	}
}
