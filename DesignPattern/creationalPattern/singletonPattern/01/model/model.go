package model

import "sync"

type stuMsg struct {
	name string
	sid  int
	age  int
}

var stuMessage *stuMsg
var once sync.Once

func GetStuMsg() *stuMsg {
	if stuMessage == nil {
		once.Do(
			func() {
				stuMessage = &stuMsg{}
			})
	}
	//} else {
	//	return stuMessage
	//}
	return stuMessage
}
func (m *stuMsg) SetMsg(name string, sid int, age int) {
	m.name = name
	m.sid = sid
	m.age = age
}
func (m *stuMsg) GetMsg() (string, int, int) {
	return m.name, m.sid, m.sid
}
