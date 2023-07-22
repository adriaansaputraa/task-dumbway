const xhttp = new XMLHttpRequest();

xhttp.open('GET',"https://your-url",true)
//parameter 1 : is the method
//parameter 2 : place of data by url
//parameter 3 : true -> asynchronous, false -> synchronous




xhttp.onload = function(){

}//mengecek status

xhttp.onerror = function(){

}//menampilkan error ketika request

xhttp.send()