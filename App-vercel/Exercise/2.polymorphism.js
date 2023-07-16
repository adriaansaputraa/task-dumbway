class vehicle{
    drive(){
        return "The vehicle is driving"
    }
}

class car extends vehicle{
    drive(){
        return "the car is driving"
    }
}

class electriCar extends vehicle{
    drive(){
        return "the electric car is driving"
    }
}

const  myvehicle = new vehicle()
const myCar = new car()
const myelectriCar = new electriCar()


console.log(myvehicle.drive())
console.log(myCar.drive())
console.log(myelectriCar.drive())