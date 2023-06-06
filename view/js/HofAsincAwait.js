// const promis = new Promise((ada, gada) => {
//     const xhr = new XMLHttpRequest()

//     xhr.open('GET', 'https://api.npoint.io/34541eb1ead82da55d27', true)
//     xhr.onload = () => {
//         if(xhr.status === 200) {
//             ada(JSON.parse(xhr.response))
//         }else {
//             gada('error hehe')
//         }
//     };
//     xhr.onerror = () => {
//         gada('network error')
//     };
//     xhr.send();
// })


// const promis = new Promise((dataAda, dataGada) => {
//     const xh = new XMLHttpRequest();
//     xh.open('GET', 'https://api.npoint.io/34541eb1ead82da55d27', true);
//     xh.onload = () => {
//         if (xh.status === 200) {
//             dataAda(JSON.parse(xh.response));
//         } else {
//             dataGada('error')
//         }
//     };

//     xh.onerror = () => {
//         dataGada('net error')
//     }
//     xh.send();
//     // console.log('ini data', xh)
// });

// async function tampilkanData() {
//     const respon = await promis;
//     console.log('ini data', respon) 
// }

// tampilkanData()

const promis = new Promise((dataAda, dataGada) => {
    const xhr = new XMLHttpRequest();

    xhr.open("GET", "https://api.npoint.io/34541eb1ead82da55d27", true);
    xhr.onload = () => {
        if(xhr.status === 200) {
            dataAda(JSON.parse(xhr.response));
        }else{
            dataGada("URL Error")
        }
    }
    xhr.onerror = () => {
        dataGada('network error');
    }
    xhr.send();
})

    async function DataKontak(rating) {
        const response = await promis;
        console.log('ini data', response)

        let cardHtml = "";

        response.map((d) => {
            cardHtml += `<div class="card" style="width: 18rem;">
            <img src="${d.image}" class="card-img-top" alt="...">
            <div class="card-body">
              <h5 class="card-title">${d.name}</h5>
              <p class="card-text">${d.note}</p>
              <p class="card-text">${d.rating}</p>
            </div>
          </div>`
        })
        document.getElementById('hehe').innerHTML = cardHtml;

    }


    
    DataKontak()



