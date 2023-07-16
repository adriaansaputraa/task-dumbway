class Car{

    #name = ""
    #model = ""
    #duit = 0

    constructor(name,model,duit){
        this.#name = name
        this.#model = model
        this.#duit = duit
    }

    //getter

    get name(){
        return this.#name
    }

    get model(){
        return this.#model
    }

    get duit(){
        return this.#duit
    }

    //Setter

    set duit(value){
        if(value < 100){
            this.#duit = this.#duit + value
        }else{
            this.#duit = value
        }
    }
}
    let myCar = new Car("red","fortuner",20000)
    myCar.duit = 9999

    console.log(myCar.name)
    console.log(myCar.model)
    console.log(myCar.duit)
    