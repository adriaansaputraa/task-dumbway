let dataBlog = [];

function addBlog(event) {
    event.preventDefault();

    let projectName = document.getElementById("input-projectname").value;
    let startDate = document.getElementById("input-startdate").value;
    let endDate = document.getElementById("input-endDate").value;
    let description = document.getElementById("input-descripton").value;
    let nodejs = document.getElementById("input-nodejs").checked;
    let golang = document.getElementById("input-golang").checked;
    let reactjs = document.getElementById("input-reactjs").checked;
    let javascript = document.getElementById("input-javascript").checked;
    let file = document.getElementById("input-image").files;
    let getimage = document.getElementById("input-image").value;

    //DISTANCE DATE

    let start = new Date(startDate);
    let end = new Date(endDate);

    let distance = end - start;
    let hourInDay = 86400000; // conver milisecond -> 1 day
    let dayInWeek = 7;
    let dayInMonth = 30;
    let monthInYear = 12;

    let distanceInDay = Math.floor(distance/(hourInDay)) //Day
    let distanceInWeek = Math.floor(distance/(hourInDay*dayInWeek)) //Week
    let distanceInMonth = Math.floor(distance/(hourInDay*dayInMonth)) //Month
    let distanceInYear = Math.floor(distance/(hourInDay*dayInMonth*monthInYear)) //Year

    console.log(distanceInDay)
    console.log(distanceInWeek)
    console.log(distanceInMonth)
    console.log(distanceInYear)

    if(projectName == ""){
        return alert ("Project Name must be entered!")
    }else if(startDate == ""){
        return alert ("Start Date must be entered!")
    }else if(endDate == ""){
        return alert ("End Date must be entered!")
    }else if(description == ""){
        return alert ("Description must be entered!")
    }else if(getimage == ""){
        return alert ("Needs upload image")
    }



    duration = "";


    if(distanceInDay == 0){
        duration = "24 jam"

    }else if(distanceInDay < 0){
        return alert("wrong input")

    }else if(distanceInDay < 8){
        duration = `${distanceInDay} day`

    }else if(distanceInWeek < 5){
        if(distanceInDay - (distanceInWeek*7) > 0){
            duration = `${distanceInWeek} Week ${distanceInDay-(distanceInWeek*7)} day`
        }else{
            duration = `${distanceInWeek} Week`
        }

    }else if(distanceInMonth < 12){
        duration = `${distanceInMonth} Month ${distanceInDay-(distanceInMonth*30)} day`

    }else{
            duration = `${distanceInYear} year`
    }

    
    let technologies = [];
    
    if (nodejs) {
        technologies.push('<i class="fa-brands fa-js" id="javascript"></i>');
    }
    
    if (golang) {
        technologies.push('<i class="fa-brands fa-golang" id="golang"></i>');
    }

    if (reactjs) {
        technologies.push('<i class="fa-brands fa-react" id="reactsJS"></i>');
    }

    if (javascript) {
        technologies.push('<i class="fa-brands fa-java" id="java"></i>');
    }
    
    let technologiesHTML = technologies.join('');
    
    console.log(technologiesHTML);

  // TAKE THE VALUE OF IMAGES
    let image = URL.createObjectURL(file[0]);
    console.log(image);

  // PUSH BLOG TO DATABLOG
    let blog = {
        image,
        projectName,
        duration,
        description,
        technologiesHTML,
        durationPost : new Date()
    };
    
    console.log(blog);
    dataBlog.push(blog);
    
    renderBlog();
    
    console.log(dataBlog);
}

//SHOW THE RESULT OF THE FORM BLOG

function renderBlog() {
    document.getElementById("content").innerHTML = "";

    for (let i = 0; i < dataBlog.length; i++) {
        document.getElementById("content").innerHTML += 
        `<div class="container-card">
                    <a href="#"><img src= ${dataBlog[i].image} alt=""/></a>
                <h3>Dumbways Mobile App - 2023</h3>
                <span>Durasi : ${dataBlog[i].duration}</span>
                <div class="Post">
                <p>${convertdate(dataBlog[i].durationPost)}</p>
                <span>${getDurationPost(dataBlog[i].durationPost)}</span>
            </div>
            <hr>
            <p>
                ${dataBlog[i].description}
            </p>
            <div class="programming-language">
                ${dataBlog[i].technologiesHTML}
            </div>
                <div class="btn-group">
                    <button>Edit</button>
                    <button>Delete</button>
                </div>
        </div>`;
    }
}

// SHOW THE RESULT OF DURATION POST

function getDurationPost(time){
    let timenow = new Date();
    let timePost = time;

    let durationPost = timenow - timePost;

    let Seconds = Math.floor(durationPost/1000);
    let Minutes = Math.floor(Seconds/60);
    let Hours = Math.floor(Minutes/60)
    let day = Math.floor(Hours/24)
    let month = Math.floor(day/30)
    let year = Math.floor(month/12)

    // floor 1.5 -> 1
    // ceil 1.3 -> 2
    // round 1.3 -> 1   1.6 -> 2

    if(Seconds >= 60 && Minutes < 60){
        return `${Minutes} minute ago..`

    }else if(Minutes >= 60 && Hours < 60){
        return `${Hours} hour ago..`

    }else if(Hours >= 60 && day < 30){
        return `${day} day ago..`

    }else if(day >= 30 && month < 12){
        return `${month} Month ago..`
    }else if(month >= 12){
        return `${year} year ago..`
    }else{
        return `${Seconds} second ago..`
    }
}

function convertdate(date){
    const getdate = date.getDate();

    const listMonth = ["January","Febuary","March","April","May","June","July","August","September","October","November","December"]

    const getMonth = listMonth[date.getMonth()];

    const getYear = date.getFullYear();

    let getHours = date.getHours();

    let getMinute = date.getMinutes();

    if(getHours < 10){
        getHours = "0"+ getHours;
    }

    if(getMinute < 10){
        getMinute = "0"+ getMinute;
    }

    return `Post : ${getdate} ${getMonth} ${getYear} | ${getHours}:${getMinute}`
}

setInterval(function(){
    renderBlog()
}, 1000)
