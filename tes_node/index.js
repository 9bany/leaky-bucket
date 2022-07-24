const axios = require('axios');

function getMovieTitles() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            let url = "http://localhost:8080/hello"
            axios.get(url)
                .then(function (response) {
                    resolve(response.data)
                }).catch(err => {
                    resolve(err.response.data)
                })
        }, 100)

    })

}

main()
async function main() {
    var index = 0
    while (true) {
        console.log(index)
        var data = await getMovieTitles()
        console.log(data)
        index++
    }
}