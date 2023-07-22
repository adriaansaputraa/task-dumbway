//function untuk mengambil data contact
function submitData(event){
    event.preventDefault()

    let name = document.getElementById("input-name").value
    let email = document.getElementById("input-email").value
    let phone = document.getElementById("input-phone").value
    let subject = document.getElementById("input-subject").value
    let message = document.getElementById("input-message").value
    const formElement = document.getElementById("myForm");
    const emailReceiver = "adriaansaputraa@gmail.com";

    if (name === "") {
        return alert('Name harus diisi !')
    } else if (email === "") {
        return alert('Email harus diisi !')
    } else if (phone === "") {
        return alert('Phone harus diisi !')
    } else if (subject === "") {
        return alert('Subject harus diisi !')
    } else if (message === "") {
        return alert('Message harus diisi !')
    }

    let a = document.createElement('a')
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Hello my name ${name},\n${message}, Lets talk with me asap sent from ${email} or my contact ${phone}`
    a.click()

    formElement.reset();
}