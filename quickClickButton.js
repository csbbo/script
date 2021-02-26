
var interval = setInterval(function(){ work() }, 10);
// var btn = document.getElementById('sign_up_btn')
var btn = document.getElementsByClassName('SearchBar-tool')[0].getElementsByTagName('button')[0].click()

var cnt = 0
function work() {
    btn.click()
    console.log(++cnt)

    var targetTime = '2021-02-04 10:54:00'
    var currentTime = new Date()
    targetTime = new Date(Date.parse(targetTime.replace(/-/g,"/")))
    if (currentTime > targetTime) {
        clearInterval(interval)
    }
}

