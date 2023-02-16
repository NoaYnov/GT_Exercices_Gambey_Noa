let i = 0;

showSlide = (n) => {
    let ids = [0, 1, 2]
    let thisId = parseInt(i + n)
    if (thisId > 2){
        thisId = 0
    }else if (thisId < 0) {
        thisId = 2
    }
    let idsWithoutCurrent = ids.filter((id) => id != thisId)
    i = thisId
    const element = document.getElementById(`carousel-Badge-${thisId}`)
    element.classList.add("show")
    element.classList.add("fadein")
    idsWithoutCurrent.map((id) => {
        let thisElement = document.getElementById(`carousel-Badge-${id}`)
        thisElement.classList.remove("show")
        thisElement.classList.add("fadeout")
    })
}