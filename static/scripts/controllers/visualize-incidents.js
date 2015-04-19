'use strict';

/**
* @ngdoc function
* @name demoApp.controller:MainCtrl
* @description
* # MainCtrl
* Controller of the demoApp
*/
angular.module('demoApp')
.controller('VisualizeIncidentsCtrl', function ($scope, ngNotify) {
	
	if (window["WebSocket"]) {
		var wsPath = "ws://"+window.document.location.host+"/ws";
        var conn = new WebSocket(wsPath);

        conn.onclose = function(evt) {
        	console.log("Connection closed");
        }
        conn.onmessage = function(evt) {
        	var data = JSON.parse(evt.data)
        	var incident = data.d;
        	console.log(incident);
        	displayIncident(incident);
        }
    } else {
        console.log("Your browser does not support WebSockets")
    }

	var platform = new H.service.Platform({
        'app_id': '3km11WSzxkHOAgiPyMkD',
        'app_code': '6ca5fCTAZdajHHd8sQpHwA'
      });

	var defaultLayers = platform.createDefaultLayers();
	
	var map = new H.Map(
		document.getElementById('mapContainer'),
		defaultLayers.satellite.map,
		{
			zoom: 3,
			center: { lat: 35, lng: 5 }
		});

	// var behavior = new H.mapevents.Behavior(new H.mapevents.MapEvents(map));

	function displayIncident(incident){
		var icon = new H.map.Icon('images/'+incident.cause+'.marker.png');
		var marker = new H.map.Marker({ lat: incident.latitude, lng: incident.longitude }, { icon: icon });
		map.addObject(marker);
	}
});
