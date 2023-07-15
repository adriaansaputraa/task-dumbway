class car{
    constructor(color,price){
        this.color = color
        this.price = price
    }

    getInfo(){
        return `I have a car with color ${this.color}, buy it in ${this.price}`
    }

}

class electricCar extends car{
    constructor(color,price,batteryCapacity){
        super(color,price)
    }
}

const mobil1 = new electricCar("blue",20000,200);
console.log(mobil1.getInfo())