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

//Routing
app.get('/', Home)
app.get('/contact', Contact)
app.get('/my-project', MyProject)
app.get('/project-detail/:id', ProjectDetail)
app.post('/my-project', AddProject)

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
    res.render('myproject') 
}

//ProjectDetail
function ProjectDetail(req, res){
    const {id} = req.params
    res.render('project-detail') 
}

function AddProject(req, res){
    const projectname = req.body.ProjectName
    const startdate = req.body.StartDate
    const enddate = req.body.EndDate
    const description = req.body.Descripton
    const nodejs = req.body.Nodejs
    const golang = req.body.Golang
    const reactjs = req.body.Reactjs
    const javascript = req.body.Javascript
    const InputImage = req.body.InputImage


    console.log(projectname)
    console.log(startdate)
    console.log(enddate)
    console.log(enddate)
    console.log(enddate)
    console.log(enddate)
    console.log(enddate)
    

}