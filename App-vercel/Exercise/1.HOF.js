// function getcapture(camera){ //HOF
//     camera("adrian"); 
// }

// getcapture (function(nama){ //Anonymous function & Callback function
//     console.log("hallo bang"+nama)
// });

function returnFunction(){
    return function (nama){
        console.log("hello"+nama)
    }
}

const nilaifunction = returnFunction()

console.log(nilaifunction());
