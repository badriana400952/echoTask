
const promise = new Promise((dataAda, dataTidakAda) => {
    const xx = new XMLHttpRequest();

    xx.open("GET", "https://api.npoint.io/ea470c91349b54c73573", true);

    xx.onload = () => {
        if (xx.status === 200) {
            dataAda(JSON.parse(xx.response));
        } else {
            dataTidakAda('error broh');
        }
    };

    xx.onerror = () => {
        dataTidakAda('network error');
    }
    xx.send();
    // console.log(xx)
})

async function dataMahasiswa() {
    const response = await promise;
    console.log('ini promisis', response)
    let dataHtml = "";
    response.map((d) => {
        dataHtml += ` <tr>
        <td>${d.id}</td>
        <td>${d.nama}</td>
        <td>${d.nim}</td>
        <td>${d.jurusan}</td>
    </tr>`
    });
    document.getElementById('table').innerHTML = dataHtml;
}

dataMahasiswa()
//buat function dari button search
//menggunakan async await dalam mendapatkan janji dari url
async function mapDataMahasiswa(nama) {
    const response = await promise;
    console.log(response)
// buat variabel input untuk mendapatkan input di html berdasarka id "search", dan ambil property value
    let input = document.getElementById('search');
    let value = input.value
    // lalu tampilkan apa yg di tulis dari input
    console.log("ini cari", value)
    //buat variabel filterSearching yang di dalam nya terdapat hasil respon dari async await. lalu filter dengan parameter yang akan di ulang
    const filterSearching = response.filter((d) => {
        //lalu jalankan parameter(d) jika sama dengan value
        //ingat value itu variabel yang berisi dom
        return d.nama === value;
    })
        // buat variabel kosng
    let dataHtml = "";
    // lalu hasil dari filter di pecah atau di ulang dengan parameter (d) 
    filterSearching.map(function (d) {
        // isikan data html yang tadi dengan tag atau data yang akan di ulang
        dataHtml += `
            <tr>
                <td>${d.id}</td>
                <td>${d.nama}</td>
                <td>${d.nim}</td>
                <td>${d.jurusan}</td>
            </tr>
        `
    });
    // lalu javascript carikan saya elemen berdasarkan id table, manipulasi dia dan isikan data html tadi
    document.getElementById('table').innerHTML = dataHtml;

}

[
    {
      name: "Abel Dustin",
      note: "Jagalah Kebersihan",
      image:
        "https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
      rating: 5,
    },
    {
      name: "Cintara Surya",
      note: "Keren cuys!!",
      image:
        "https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
      rating: 4,
    },
    {
      name: "Dandi Cuy",
      note: "The Best Pelayanannya!",
      image:
        "https://images.unsplash.com/photo-1564564321837-a57b7070ac4f?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8OHx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
      rating: 5,
    },
    {
      name: "bob",
      note: "lorem ipsums!",
      image:
        "https://images.unsplash.com/photo-1682687982502-1529b3b33f85?ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHw0NXx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
      rating: 2,
    },
  ];