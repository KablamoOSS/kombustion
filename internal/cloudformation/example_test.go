package cloudformation

import "log"

func Example_GenerateYamlStack() {
	params := map[string]string{"size": "5"}
	out, err := GenerateYamlStack(GenerateParams{
		Filename: "api",
		Env:      "prod",
		ParamMap: params,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out)
}
