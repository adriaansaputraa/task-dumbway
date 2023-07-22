const promise = new Promise((resolve,reject)=>{
    const xhttp = new XMLHttpRequest();
    xhttp.open('GET',"https://api.npoint.io/f334577a286bc694ca02",true)
    
    xhttp.onload = function(){
        if(xhttp.status === 200){
            resolve(JSON.parse(xhttp.response))
        }else if(xhttp.status >= 400){
            reject("Error loading data")
        }
    }//mengecek status didalam server

    xhttp.onerror = function(){
        reject("Network error")

    }//menampilkan error ketika request ke server   
    xhttp.send()
})

//promise-chaining
// promise.then((value)=>{
//     console.log(value)
//     let TestimonialHTML = ""
    
//     value.forEach((card)=>{
//     TestimonialHTML += 
//     `<div class="testimonial">
//         <img src="${card.image}" class="profile-testimonial" />
//         <p class="quote">"${card.quote}"</p>
//         <p class="author">- ${card.user}</p>
//         <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
//     </div>`
//     })

//     document.getElementById("testimonials").innerHTML = TestimonialHTML

// }).catch((err)=>{
//     console.log(err)
// })

// function filterTestimonial(rating){
    
//     promise.then((value)=>{
//     let FilteredTestimonialHTML = ""

//     const filteredData = value.filter((card)=>{
//         return card.rating === rating
//     })

//     filteredData.forEach((card)=>{
//         FilteredTestimonialHTML += 
//         `<div class="testimonial">
//             <img src="${card.image}" class="profile-testimonial" />
//             <p class="quote">"${card.quote}"</p>
//             <p class="author">- ${card.user}</p>
//             <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
//         </div>`
//     })

//     document.getElementById("testimonials").innerHTML = FilteredTestimonialHTML
// }).catch((err)=>{
//     console.log(err)
// })
// }

// function allTestimonial(){
    
//     promise.then((value)=>{
//     let FilteredTestimonialHTML = ""

//     value.forEach((card)=>{
//         FilteredTestimonialHTML += 
//         `<div class="testimonial">
//             <img src="${card.image}" class="profile-testimonial" />
//             <p class="quote">"${card.quote}"</p>
//             <p class="author">- ${card.user}</p>
//             <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
//         </div>`
//     })

//     document.getElementById("testimonials").innerHTML = FilteredTestimonialHTML
// }).catch((err)=>{
//     console.log(err)
// })
// }

//async-await
let TestimonialData = []
async function getData(){
    try{
        const response = await promise;
        TestimonialData = response
        allTestimonial()
    }catch(err){
        console.log(err)
    }
}

getData()

function allTestimonial() {
    let testimonialHTML = ""

    TestimonialData.forEach((card, index) => {
        testimonialHTML += 
        `<div class="testimonial">
        <img src="${card.image}" class="profile-testimonial" />
        <p class="quote">"${card.quote}"</p>
        <p class="author">- ${card.user}</p>
        <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
        </div>`
    })

    document.getElementById("testimonials").innerHTML = testimonialHTML
}

function filterTestimonial(rating) {
    let filteredTestimonialHTML = ""

    const filteredData = TestimonialData.filter((card) => {
        return card.rating === rating
    }) 

    filteredData.forEach((card) => {
        filteredTestimonialHTML += 
        `<div class="testimonial">
        <img src="${card.image}" class="profile-testimonial" />
        <p class="quote">"${card.quote}"</p>
        <p class="author">- ${card.user}</p>
        <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
        </div>
        `
    })

    document.getElementById("testimonials").innerHTML = filteredTestimonialHTML
}