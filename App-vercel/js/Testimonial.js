//fake data -> array of object
const TestimonialData = [
    {
        user:"Surya Elidanto",
        quote:"Keren banget jasanya!",
        image:"https://images.unsplash.com/photo-1501196354995-cbb51c65aaea?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1171&q=80",
        rating: 5
    },
    {
        user:"Surya Elz",
        quote:"Keren lah pokoknya!",
        image:"https://images.unsplash.com/photo-1463453091185-61582044d556?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1170&q=80",
        rating:4
    },
    {
        user:"Surya Gans",
        quote:"The best pelayanannya!",
        image:"https://images.unsplash.com/photo-1534528741775-53994a69daeb?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=464&q=80",
        rating:4
    },
    {
        user:"Suryaaa",
        quote:"Oke lah!",
        image:"https://images.unsplash.com/photo-1537511446984-935f663eb1f4?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1170&q=80",
        rating:3
    },
    {
        user:"Suryeah",
        quote:"Apa apaan ini!",
        image:"https://images.unsplash.com/photo-1522529599102-193c0d76b5b6?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1170&q=80",
        rating:1
    },

]

let TestimonialHTML = ""

TestimonialData.forEach((card)=>{
    TestimonialHTML += 
    `<div class="testimonial">
        <img src="${card.image}" class="profile-testimonial" />
        <p class="quote">"${card.quote}"</p>
        <p class="author">- ${card.user}</p>
        <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
    </div>`
})

document.getElementById("testimonials").innerHTML = TestimonialHTML

//Filter testimonial by rating

function filterTestimonial(rating){
    
    let FilteredTestimonialHTML = ""

    const filteredData = TestimonialData.filter((card)=>{
        return card.rating === rating
    })

    filteredData.forEach((card)=>{
        FilteredTestimonialHTML += 
        `<div class="testimonial">
            <img src="${card.image}" class="profile-testimonial" />
            <p class="quote">"${card.quote}"</p>
            <p class="author">- ${card.user}</p>
            <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
        </div>`
    })

    document.getElementById("testimonials").innerHTML = FilteredTestimonialHTML
}

//Show all the data

function allTestimonial(){
    
    let FilteredTestimonialHTML = ""

    TestimonialData.forEach((card)=>{
        FilteredTestimonialHTML += 
        `<div class="testimonial">
            <img src="${card.image}" class="profile-testimonial" />
            <p class="quote">"${card.quote}"</p>
            <p class="author">- ${card.user}</p>
            <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
        </div>`
    })

    document.getElementById("testimonials").innerHTML = FilteredTestimonialHTML
}