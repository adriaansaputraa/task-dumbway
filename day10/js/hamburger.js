let hamburgerIsOpen = false
function openHamburger() {
    let hamburgerContainer = document.getElementById("hamburger-container")
    if (!hamburgerIsOpen) {
        hamburgerContainer.style.display = "flex"
        hamburgerIsOpen = true;

        console.log(hamburgerIsOpen)
    } else {
        hamburgerContainer.style.display = "none"
        hamburgerIsOpen = false;
    }
}