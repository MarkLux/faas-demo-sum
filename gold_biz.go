package gold

import (
	"github.com/MarkLux/GOLD/serving/common"
	"log"
)

/**
  * function service example
  * show usage of rpc, db & cache
 */


// would be invoked when the service launch at first time.
func (s *GoldService) OnInit() {
	// do something here..
}

// the biz function
func (s *GoldService) OnHandle(req *common.GoldRequest, rsp *common.GoldResponse) error {
	// get data from request
	a := req.Data["a"].(float64)
	b := req.Data["b"].(float64)
	
	sum := a+b
	rsp.Data = make(map[string]interface{})

	rsp.Data["sum"] = sum

	return nil
}

// when an error is throwed, this function would be called
// if you don't want to break the process, return true
// otherwise you should return false
func (s *GoldService) OnError(err error) bool {
	log.Println(err)
	return false
}
