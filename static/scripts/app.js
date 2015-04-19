'use strict';

/**
* @ngdoc overview
* @name demoApp
* @description
* # demoApp
*
* Main module of the application.
*/
angular
.module('demoApp', [
  'ngRoute',
  'ngTouch',
  'ui.bootstrap',
  'ngNotify'
])

.config(function ($routeProvider) {
  $routeProvider
  .when('/incidents/visualize', {
    templateUrl: 'views/visualize-incidents.html',
    controller: 'VisualizeIncidentsCtrl'
  })
  .otherwise({
    redirectTo: '/incidents/visualize'
  });
})
