var dataElement = document.querySelector("script[type='application/json']");
var data = JSON.parse(dataElement.textContent);

var map = L.map("map");

L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
	attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
}).addTo(map);

let markers = [];

data.forEach((fuente) => {
	marker = L.marker([fuente.Latitud, fuente.Longitud])
		.addTo(map)
		.bindPopup(`<b>${fuente.Titulo}<br><br><a href="/fuentes/${fuente.Id}">Más información</a>`);

	markers.push(marker);
});

let group = new L.featureGroup(markers);
map.fitBounds(group.getBounds().pad(0.035));
