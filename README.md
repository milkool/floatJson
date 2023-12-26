# jfloat
Specify the format of the decimal point when serializing json

## example
```
type JsonTest struct {
	S1 string        `json:"s1"`
	J1 jfloat.JFloat `json:"j1"`
}

func main() {
	jt := JsonTest{
		S1: "exsample",
		J1: jfloat.JFloat{
			Value:  100,
			Digits: 2,
		},
	}

	jsonOutput, _ := json.Marshal(jt)
	fmt.Printf("%s\n", jsonOutput)
}
```

## output
```
{"s1":"exsample","j1":100.00}
```
