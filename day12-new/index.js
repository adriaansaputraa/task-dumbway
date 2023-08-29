const express = require('express') //import express
const app = express()
const PORT = 5000
const path = require('path')

//setup call hbs
app.set('view engine', 'hbs')

//set serving static file
app.use('/Assets', express.static('Assets'));

//parsing data 
app.use(express.urlencoded({ extended: false}))

const dataProject = [
    {
        ProjectName: "Project fake-1",
        StartDate: "2023-08-28",
        EndDate: "2023-08-30",
        Description: "Ini fake project untuk coba",
        Nodejs: true,
        Golang: false,
        Reactjs: true,
        Javascript: true,
        postDate: new Date()
    },
    {
        ProjectName: "Project fake-2",
        StartDate: "2023-08-28",
        EndDate: "2023-08-29",
        Description: "Ini fake project untuk coba kedua",
        Nodejs: true,
        Golang: true,
        Reactjs: false,
        Javascript: true,
        postDate: new Date()
    }
]

//Routing
app.get('/', Home)
app.get('/contact', Contact)
app.get('/my-project', MyProject)
app.get('/project-detail/:id', ProjectDetail)
app.get('/delete-project/:id', DeleteProject)
app.post('/my-project', AddProject)
app.get('/edit-Project/:id', EditProject)
app.post('/edit-Project/:id', UpdateProject)

//local server
app.listen(PORT, () => {
    console.log("Example app listening on port 5000")
})

//index
function Home (req, res){
    res.render('index') 
}

//contact
function Contact (req, res){
    res.render('contact') 
}

//MyProject
function MyProject(req, res){
    res.render('myproject', {dataProject}) 
}

//ProjectDetail
function ProjectDetail(req, res){
    const {id} = req.params

    res.render('project-detail',{data: dataProject[id]}) 
}

function EditProject(req, res){
    const {id} = req.params

    res.render('edit-project',{data: dataProject[id]}) 
}

function UpdateProject(req, res){
    const id = parseInt(req.params.id, 10);

    const ProjectIndex = dataProject.findIndex((project) => project.id === id);

    const{ProjectName, StartDate, EndDate, Description, Nodejs, Golang, Reactjs, Javascript} = req.body

    const duration = getDurationPost(StartDate, EndDate)
    
    const newData = {
        ProjectName,
        Description,
        StartDate,
        EndDate,
        Nodejs,
        Golang,
        Reactjs,
        Javascript,
        duration,
        postDate : convertdate(new Date())
    }

    console.log(req.body)
    
    if (ProjectIndex !== -1) {
    dataProject.splice(ProjectIndex, 1, newData);
} else {
    // Handle the error, for example:
    return res.status(404).send("Project not found");
}

    res.redirect('/my-project') 

}


function DeleteProject(req, res){
    const {id} = req.params
    dataProject.splice(id, 1)
    res.redirect('/my-project')
}

function AddProject(req, res){
    const{ProjectName, StartDate, EndDate, Description, Nodejs, Golang, Reactjs, Javascript} = req.body

    const duration = getDurationPost(StartDate, EndDate)
    
    const newData = {
        ProjectName,
        Description,
        StartDate,
        EndDate,
        Nodejs,
        Golang,
        Reactjs,
        Javascript,
        duration,
        postDate : convertdate(new Date())
    }

    

    dataProject.push(newData)
    console.log(req.body)

    res.redirect('/my-project') 

}

function getDurationPost(StartDate,EndDate){
    let start = new Date(StartDate);
    let end = new Date(EndDate);

    let distance = end - start;
    let hourInDay = 86400000; // convert milisecond -> 1 day
    let dayInWeek = 7;
    let dayInMonth = 30;
    let monthInYear = 12;

    let distanceInDay = Math.floor(distance/(hourInDay)) //Day
    let distanceInWeek = Math.floor(distance/(hourInDay*dayInWeek)) //Week
    let distanceInMonth = Math.floor(distance/(hourInDay*dayInMonth)) //Month
    let distanceInYear = Math.floor(distance/(hourInDay*dayInMonth*monthInYear)) //Year

    let duration = "";

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

    return duration
}

dataProject.forEach((data)=>{
    data.duration = getDurationPost(data.StartDate,data.EndDate)
})

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

dataProject.forEach((data)=>{
    data.postDate= convertdate(data.postDate)
})
