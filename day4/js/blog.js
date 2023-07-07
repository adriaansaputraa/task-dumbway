let dataBlog = [];

function addBlog(event) {
    event.preventDefault();

    let projectName = document.getElementById("input-projectname").value;
    let startDate = new Date(document.getElementById("input-startdate").value);
    let endDate = new Date(document.getElementById("input-endDate").value);
    let description = document.getElementById("input-descripton").value;
    let nodejs = document.getElementById("input-nodejs").checked;
    let golang = document.getElementById("input-golang").checked;
    let reactjs = document.getElementById("input-reactjs").checked;
    let javascript = document.getElementById("input-javascript").checked;
    let file = document.getElementById("input-image").files;

    //DISTANCE DATE

    let distance = endDate - startDate;
    let miliSecond = 1000;
    let secondInHour = 3600; // convert to second
    let hourInDay = 24;
    let dayInWeek = 7;
    let dayInMonth = 30;
    let monthInYear = 12;

    let distanceInDay = Math.floor(distance/(miliSecond*secondInHour*hourInDay)) //Day
    let distanceInWeek = Math.floor(distance/(miliSecond*secondInHour*hourInDay*dayInWeek)) //Week
    let distanceInMonth = Math.floor(distance/(miliSecond*secondInHour*hourInDay*dayInMonth)) //Month
    let distanceInYear = Math.floor(distance/(miliSecond*secondInHour*hourInDay*dayInMonth*monthInYear)) //Year

    console.log(distanceInDay)
    console.log(distanceInWeek)
    console.log(distanceInMonth)
    console.log(distanceInYear)

    duration = "";


    if(distanceInDay < 8){
        duration = `${distanceInDay} day`
    }else if(distanceInWeek < 5){
        if(distanceInDay - (distanceInWeek*7) > 0){
            duration = `${distanceInWeek} Week ${distanceInDay-(distanceInWeek*7)} day`
        }else{
            duration = `${distanceInWeek} Week`
        }
    }else{
        duration = "24 jam"
    }


    // if(distanceInDay < 8 && distanceInDay > 0){
    //     duration = `${distanceInDay} day`
    // }
    // else if(distanceInDay >= 8 && distanceInWeek <= 4){
        
    //     if(distanceInDay-(distanceInWeek*7) >= 1){
    //         duration = `${distanceInWeek} Week ${distanceInDay-(distanceInWeek*7)} day`
    //     }
    //     else{
    //         duration = `${distanceInWeek} Week`
    //     }
    // }
    // else if(distanceInDay >= 8 && distanceInMonth < 12){
    //     duration = `${distanceInMonth} Month ${distanceInDay-30} day`
    // }
    // else if(distanceInMonth >= 12){
    //     if(duration = `${distanceInYear} year`){

    //     }

    // }



    // else{
    //     return alert("wrong input")
    // }


    // CHECK BOX FEATURES
    
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

