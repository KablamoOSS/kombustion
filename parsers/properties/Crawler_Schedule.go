package properties


type Crawler_Schedule struct {
	
	ScheduleExpression interface{} `yaml:"ScheduleExpression,omitempty"`
}

func (resource Crawler_Schedule) Validate() []error {
	errs := []error{}
	
	return errs
}
