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
        	console.log(evt.data);
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

});
