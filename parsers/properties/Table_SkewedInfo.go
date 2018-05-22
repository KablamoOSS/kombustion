package properties


type Table_SkewedInfo struct {
	
	
	
	SkewedColumnValueLocationMaps interface{} `yaml:"SkewedColumnValueLocationMaps,omitempty"`
	SkewedColumnNames interface{} `yaml:"SkewedColumnNames,omitempty"`
	SkewedColumnValues interface{} `yaml:"SkewedColumnValues,omitempty"`
}

func (resource Table_SkewedInfo) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
