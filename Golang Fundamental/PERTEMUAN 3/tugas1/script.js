let judul = document.getElementById("judul");
let tombol = document.getElementById("tombol");
let klik = 0;

tombol.addEventListener('click', function() {
    klik++;
    switch (klik) {
        case 1:
            judul.innerHTML = "programmer handal";
            break;
        case 2:
            judul.innerHTML = "Aku pasti bisa menjadi programmer";
            break;
        default:
            judul.innerHTML = "Lumoshive Academy";
            klik = 0;
            break;
    }
});