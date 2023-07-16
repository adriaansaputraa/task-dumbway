class Testimonial{ //class
    #quote = "" //encapsulation
    #image = "" //encapsulation

    constructor (quote,image){
        this.#quote = quote
        this.#image = image
    }

    get quote(){
        return this.#quote
    }

    get image(){
        return this.#image
    }

    get User(){
        throw new Error ("there is must be user to make testimonials")
    }

    get testimonialHTMLclass(){ //polymorphism
        return `<div class="testimonial">
                    <img src="${this.image}" class="profile-testimonial" />
                    <p class="quote">"${this.quote}"</p>
                    <p class="author">- ${this.User}</p>
                </div>` 
    }
}

class UserTestimonial extends Testimonial{
    #user = ""
    
    constructor (user,quote,image){
        super(quote,image) // inheritance
        this.#user = user
    }

    get User(){ //polymorphism & abstraction
        return "user : "+this.#user
    }
}

class CompanyTestimonial extends Testimonial{
    #company = ""
    
    constructor (company,quote,image){
        super(quote,image) // inheritance
        this.#company = company
    }
    
    get User(){ //polymorphism & abstraction
        return "company : "+this.#company
    }
}
const testimonial = new UserTestimonial("Surya Ellidanto","Mantab sekali jasanya!","https://images.unsplash.com/photo-1501196354995-cbb51c65aaea?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1171&q=80") //object

const testimonial2 = new UserTestimonial("Surya Elz","Keren lah pokoknya!","https://images.unsplash.com/photo-1463453091185-61582044d556?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1170&q=80") //object


const testimonial3 = new CompanyTestimonial("ABC company","wuhuu keren lah!","https://images.unsplash.com/photo-1534528741775-53994a69daeb?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=464&q=80") //object


let testimonialData = [testimonial,testimonial2,testimonial3]
let testimonialHTML = ""

for(let i = 0; i<testimonialData.length; i++){
    testimonialHTML += testimonialData[i].testimonialHTMLclass
    
}

document.getElementById("testimonials").innerHTML = testimonialHTML