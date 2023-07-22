let condition = true;

let promise = new Promise((resolve,reject) =>{
    if(condition){
        setTimeout(()=>{
            resolve("Janji ditepati")
        },3000)
    }else{
        reject("Janji gugur!")
    }
})

async function getData(){
    try{
        const response = await promise;
        console.log(response)
    } catch(err){
        console.log(err)
    }
    
}

getData()