var dataElement = document.querySelector("script[type='application/json']");
var data = JSON.parse(dataElement.textContent);

var map = L.map("map").setView([42.21, -8.7389], 12);

L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
	attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
}).addTo(map);

data.forEach((fuente) => {
	L.marker([fuente.Latitud, fuente.Longitud])
		.addTo(map)
		.bindPopup(`<b>${fuente.Titulo}<br><br><a href="/fuentes/${fuente.Id}">Más información</a>`);
});
