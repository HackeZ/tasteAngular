var app = angular.module('myApp', []);

app.controller('HelloController', function($scope) {
    $scope.greeting = { text: 'Hello' };    
});