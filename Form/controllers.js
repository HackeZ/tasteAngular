app = angular.module('myApp', []);

app.controller('FormController', function ($scope) {
    $scope.youCheckedIt = true;
});

app.controller('StartUpController', function ($scope) {
    $scope.funding = {
        startingEstimate: 0
    };

    computeNeeded = function () {
        $scope.funding.needed = $scope.funding.startingEstimate * 10;
    };

    $scope.$watch('funding.startingEstimate', computeNeeded);
});

app.controller('NewStartUpController', function($scope) {
    $scope.computeNeeded = function() {
        $scope.needed = $scope.startingEstimate * 10;
    };

    $scope.requestFund = function() {
        window.alert("Sorry, please get more customers first.")
    }
});