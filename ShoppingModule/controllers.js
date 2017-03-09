var shoppingModule = angular.module('ShoppingModule', []);

shoppingModule.controller('ShoppingController', function($scope, Items) {
    $scope.items = Items.query();
});

shoppingModule.factory('Items', function() {
    var items = {};
    items.query = function() {
        // we will get data from http interface
        return [
            {title: 'Paint pots', quantity: 8, price: 3.95},
            {title: 'Polka dots', quantity: 17, price: 12.95},
            {title: 'Pebbles', quantity: 5, price:6.95}
        ];
    };
    return items;
});