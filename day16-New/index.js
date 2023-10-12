const express = require('express') //import express
const app = express()
const PORT = 5000
const path = require('path')
const bcrypt = require('bcrypt')
const session = require('express-session')
const flash = require('express-flash')

//setup call hbs
app.set('view engine', 'hbs')
app.set("views", path.join(__dirname, "src/views"));



//set serving static file
app.use(express.static(path.join(__dirname, "src/assets")));

//parsing data 
app.use(express.urlencoded({ extended: false}))

//Sequelize init
const config = require('./src/config/config.json')
const { Sequelize, QueryTypes } = require('sequelize')//sequelize untuk menjalankan sequlizenya sedangkan querytype digunakan untuk menjalankan qwery
const { start } = require('repl')
const sequelize = new Sequelize(config.development)

//setup flash
app.use(flash())

//routing session
app.use(session({
    cookie: {
        httpOnly: true,
        secure: false,
        maxAge: 1000 * 60 * 60 * 2
    },
    store: new session.MemoryStore(),
    saveUninitialized: true,
    resave:false, //jika ada perubahan automatis save
    secret:'secretValue'//nilai rahasia
}))

//Routing
app.get('/', Home)
app.get('/contact', Contact)
app.get('/my-project', MyProject)
app.get('/project-detail/:id', ProjectDetail)
app.get('/delete-project/:id', DeleteProject)
app.post('/my-project', AddProject)
app.get('/edit-Project/:id', EditProject)
app.post('/edit-Project/:id', UpdateProject)


//Login & Registration
app.get('/form-login', FormLogin)
app.post('/form-login', Login)
app.get('/form-register', FormRegister)
app.post('/form-register', AddUser)


//local server
app.listen(PORT, () => {
    console.log("Example app listening on port 5000")
})

//index
function Home (req, res){
    res.render('index',{
        isLogin:req.session.isLogin,
        user:req.session.user
    }) 
}

//contact
function Contact (req, res){
    res.render('contact') 
}

//MyProject
async function MyProject(req, res){
    try{
        const query = `SELECT id, name, start_date, end_date, description, technologies, image, userid, "createdAt", "updatedAt" FROM projects ORDER BY "updatedAt" DESC;`
        let obj = await sequelize.query(query, {type: QueryTypes.SELECT})
        
        const data = obj.map(res => ({
        ...res,
        duration:getDurationPost(res.start_date,res.end_date),
        Nodejs:res.technologies[0],
        Golang:res.technologies[1],
        Reactjs:res.technologies[2],
        Javascript:res.technologies[3],
        postDate:convertdate(res.createdAt),
        postTime:getPostTime(res.updatedAt)
        }))
        console.log(data)
        res.render('myproject', {dataProject:data})
    }catch(err){

    }
}

//ProjectDetail
async function ProjectDetail(req, res){
    const {id} = req.params
    try{
        const query = `SELECT * FROM projects WHERE id=${id}`

        let obj = await sequelize.query(query, {type: QueryTypes.SELECT})
        
        const data = obj.map(res => ({
            ...res,
            duration:getDurationPost(res.start_date,res.end_date),
            StartDate:convertdateDetail(res.start_date),
            EndDate:convertdateDetail(res.end_date),
            Nodejs:res.technologies[0],
            Golang:res.technologies[1],
            Reactjs:res.technologies[2],
            Javascript:res.technologies[3],
            PostTime:getDurationPost(res.createdAt)

            }))

        res.render('project-detail',{data:data[0]})
    }catch(err){

    } 
}

async function EditProject(req, res){
    try{
        const {id} = req.params
        const query = `SELECT * FROM projects WHERE id=${id}`

        let obj = await sequelize.query(query, {type: QueryTypes.SELECT})
        
        const data = obj.map(res => ({
            ...res,
            StartDate:convertdateEdit(res.start_date),
            EndDate:convertdateEdit(res.end_date),
            Nodejs:res.technologies[0],
            Golang:res.technologies[1],
            Reactjs:res.technologies[2],
            Javascript:res.technologies[3],
            }))
            
            console.log(data)
            res.render('edit-project',{data: data[0]}) 
    }catch(err){

    } 
}

async function UpdateProject(req, res){
    try{
        const{id, ProjectName, StartDate, EndDate, Description, Nodejs, Golang, Reactjs, Javascript}  = req.body
        const image = "image.jpg"

        let technologies = [
            Nodejs === "nodejs",
            Golang === "golang",
            Reactjs === "reactjs",
            Javascript === "javascript"
        ];

        await sequelize.query(`UPDATE projects SET name='${ProjectName}', start_date='${StartDate}', end_date='${EndDate}', description='${Description}', technologies=ARRAY[${technologies}], image='${image}', "updatedAt"=NOW() WHERE id=${id};`)

        res.redirect('/my-project') 

    }catch(err){
        console.log(err)
    } 

}


async function DeleteProject(req, res){
    try{
        const {id} = req.params
        await sequelize.query(`DELETE FROM projects WHERE id=${id}`)
        res.redirect('/my-project')
    }catch(err){
        console.log(err)
    }
    
}

async function AddProject(req, res){
    try{
        const{ProjectName, StartDate, EndDate, Description, Nodejs, Golang, Reactjs, Javascript}  = req.body
        const image = "image.jpg"

        let technologies = [
            Nodejs === "nodejs",
            Golang === "golang",
            Reactjs === "reactjs",
            Javascript === "javascript"
        ];

        await sequelize.query(`INSERT INTO "projects" 
        (name, start_date, end_date, description, technologies, image, "createdAt", "updatedAt") 
        VALUES 
        ('${ProjectName}', '${StartDate}', '${EndDate}', '${Description}', ARRAY[${technologies}], '${image}', NOW(), NOW());
        `)

        res.redirect('/my-project') 

    }catch(err){
        console.log(err)
    }
}

function FormLogin(req, res){
    res.render('form-login') 
}

async function Login(req, res){
    try{
        const {email, password} = req.body
        const query = `SELECT * FROM users WHERE email = '${email}'`
        let obj = await sequelize.query(query, {type: QueryTypes.SELECT})

        //checking if email has not been registered
        if(!obj.length){
            req.flash('danger',"User has not been registered")
            res.redirect('/form-login') 
        }

        await bcrypt.compare(password,obj[0].password,(err,result)=>{
            if(!result){
                req.flash('danger','Password wrong')
                res.redirect('/form-login') 
            }else{
                req.session.isLogin = true
                req.session.user = obj[0].name
                req.flash('success','Login berhasil')
                res.redirect('/') 
                console.log(obj)
            }
        })
    }catch(err){
        console.log(err)
    }
}

function FormRegister(req, res){
    res.render('form-register') 
}


async function AddUser(req, res){
    try{
        const {username, email, password} = req.body
        const salt = 10

        if(password === ""){
        }

        await bcrypt.hash(password, salt, (err,hashPassword)=> {
            const query = `INSERT INTO users(name, email, password, "createdAt", "updatedAt")
            VALUES ('${username}', '${email}', '${hashPassword}', NOW(), NOW());`

            sequelize.query(query)
        })
        res.redirect('/form-login') 
    }catch(err){
        console.log(err)
    }
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

function convertdateDetail(date){
    const getdate = date.getDate();

    const listMonth = ["January","Febuary","March","April","May","June","July","August","September","October","November","December"]

    const getMonth = listMonth[date.getMonth()];

    const getYear = date.getFullYear();

    return `${getdate}/${getMonth}/${getYear}`
}

function convertdateEdit(date){
    const getdate = String(date.getDate()).padStart(2, '0');
    const getMonth = String(date.getMonth() + 1).padStart(2, '0'); // +1 because months are 0-indexed
    const getYear = date.getFullYear();

    return `${getYear}-${getMonth}-${getdate}`
}

function getPostTime(time){
    let timenow = new Date();
    let timePost = time;

    let durationPost = timenow - timePost;

    let Seconds = Math.floor(durationPost/1000);
    let Minutes = Math.floor(Seconds/60);
    let Hours = Math.floor(Minutes/60)
    let day = Math.floor(Hours/24)
    let month = Math.floor(day/30)
    let year = Math.floor(month/12)

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

