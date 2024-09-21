package main

func main() {
	type spec struct {
		milege int
		fuel   string
	}
	type car struct {
		price   int
		name    string
		carSpec spec
	}

	alto := car{
		price: 500000,
		name:  "altolxi",
		carSpec: spec{
			milege: 35,
			fuel:   "petrol",
		},
	}
	println("details of car ")
	println(alto.carSpec.fuel, alto.carSpec.milege, alto.name, alto.price)

	my := struct {
		name string
		age  int
	}{
		name: "sachin",
		age:  24,
	}

	println("my anonymus struct values are ", my.name, my.age)

	//Nested structs require you to access fields through the nested struct name, providing more explicit field access.
	//add Address
	//person.add.city

	//Embedded structs promote their fields to the outer struct, allowing more convenient access but potentially causing name conflicts.
	//Address
	//person.city

}
