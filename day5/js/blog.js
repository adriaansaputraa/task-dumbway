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
    
    let technologiesHTML = technologies.join();
    
    console.log(technologiesHTML);

  // WARNING INFORMATION IF VALUES ARE EMPTY

    if (projectName == "") {
        return alert("Input the value first");
    } else if (startDate == "") {
        return alert("Input the value first");
    } else if (endDate == "") {
        return alert("Input the value first");
    } else if (description == "") {
        return alert("Input the value first");
    } else if (file == "") {
        return alert("Input the value first");
    }

  // TAKE THE VALUE OF IMAGES
    file = URL.createObjectURL(file[0]);

  // PUSH dataBlog to blog
    let blog = {
        file,
        projectName,
        duration : "3 Bulan",
        description,
        technologiesHTML,
    };
    
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
                    <a href="detail-blog.html"><img src=${dataBlog[i].file}/></a>
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