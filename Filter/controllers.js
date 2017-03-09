app = angular.module('myApp', []);

app.controller('FilterController', function($scope) {
    $scope.price = 24.8;
    $scope.name = "hacker"
    $scope.title = "hello filter"
});

app.filter('titleCase', function() {
    var titleCaseFilter = function(input) {
        var words = input.split(' ');
        for (var i = 0; i < words.length; i++) {
            words[i] = words[i].charAt(0).toUpperCase() + words[i].slice(1);
        }
        return words.join(' ');
    };
    return titleCaseFilter;
});