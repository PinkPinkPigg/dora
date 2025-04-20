package kits

import "testing"

func TestDiscoverMysqlService(t *testing.T) {
	DiscoverMService("localhost:8500", "mysql")
	//if err != nil {
	//	t.Error(err)
	//}
}
