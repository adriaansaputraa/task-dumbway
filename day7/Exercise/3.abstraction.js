class Car{

    #name = ""
    #model = ""
    #duit = 0

    constructor(name,model){
        this.#name = name
        this.#model = model
        this.#duit = duit
    }

    //getter

    get make(){
        return this.#name
    }

    get mdoel(){
        return this.#model
    }

    get model(){
        return this.#duit
    }
}
    let myCar = new Car("red","fortuner",20000)

    console.log(myCar.name)
    console.log(myCar.model)
    console.log(myCar.duit)
    