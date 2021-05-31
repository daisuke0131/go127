package a

type Person struct{
	name string `json:"name"`
}

type Person2 struct{
	name, name2 string `json:"name"`
}

type Person3 struct{
	Name string `json:"name"`
}

type Person4 struct{
	Name, name2 string `json:"name"`
}

type Person5 struct{
	nick string `json:"nick"`
	Name string `json:"name"`
}