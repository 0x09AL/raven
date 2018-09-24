package model


type CFG struct {
	Creds struct{
		Username string
		Password string
	}
	Extra struct{
		Searchengine string
	}
}

type Scan struct {
	Scan_id string
	Scan_name string
	Company string
	Domain string
	Pages_Number int
}


type Profile struct{
	Name string
	Company string
	Position string
}


type Export struct{
	Format string
	Domain string
	Output string
}

type Person struct{
	FirstName string
	LastName string
	Email string
}

type Breach struct {
	Name string `json:"Name"`
}