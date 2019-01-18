package main

import (
  "fmt"
)

type vehicle struct{
  doors int
  color string
}

type truck struct{
  vehicle
  fourWheel bool
}

type sedan struct{
  vehicle
  luxury bool
}

func main() {

  car1 := truck{
    vehicle: vehicle{
    doors: 2,
    color: "red",
    },
    fourWheel: true,
  }

  car2 := sedan{
    vehicle: vehicle{
    doors: 4,
    color: "white",
    },
    luxury: true,
  }

  fmt.Println(car1)
  fmt.Println(car2)
  fmt.Println(car1.doors, car1.color, car1.fourWheel)
  fmt.Println(car2.doors, car2.color, car2.luxury)
}
