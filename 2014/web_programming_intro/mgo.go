type Person struct {
	Name  string
	Phone string `bson:"PhOnE,omitempty"`
}

session, _ := mgo.Dial("server1.example.com,server2.example.com")
defer session.Close()

c := session.DB("test").C("people")
c.Insert(&Person{"Ale", "+55 53 8116 9639"},
         &Person{"Cla", "+55 53 8402 8510"})

var result Person
c.Find(bson.M{"name": "Ale"}).One(&result)
