document.getElementById("generateButton").addEventListener("click", function() {
    document.getElementById("RNA").textContent = getRandomInt(1, 10);
    document.getElementById("RNB").textContent = getRandomInt(1, 10);
    document.getElementById("RNC").textContent = getRandomInt(1, 10);
    document.getElementById("RND").textContent = getRandomInt(1, 10);
});

function getRandomInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}