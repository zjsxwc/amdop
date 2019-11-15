package main

import (
	"fmt"
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20191115_151429 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20191115_151429{}
	m.Created = "20191115_151429"

	migration.Register("User_20191115_151429", m)
}

// Run the migrations
func (m *User_20191115_151429) Up() {
	fmt.Println("uuuuuuuuuuuuuuuu")
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *User_20191115_151429) Down() {
	fmt.Println("dddddddddddddddd")
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
