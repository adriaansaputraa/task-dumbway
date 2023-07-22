let condition = false;

let promise = new Promise((resolve,reject) =>{
    if(condition){
        resolve("Janji ditepati")
    }else{
        reject("Janji gugur!")
    }
})

promise.then((value)=> {
    console.log(value)
}).catch((err)=>{
    console.log(err)
}).finally(() => {
    console.log("Selesai")
})