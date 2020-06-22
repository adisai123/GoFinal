package config

import "gopkg.in/mgo.v2"

func Init() {
	mgo.Dial("")
}
