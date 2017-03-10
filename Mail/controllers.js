var aMailServices = angular.module('AMail', []);

function emailRouteConfig($routeProvider) {
    $routeProvider.
    when('/', {
        controller: ListController,
        templateUrl: 'list.html'
    }).
    when('/detail/:id', {
        controller: DetailController,
        templateUrl: 'detail.html'
    }).
    otherwise({
        redirectTo: '/'
    });
}

// 配置路由，以便 AMail 服务能够找到它
aMailServices.config(emailRouteConfig);

messages = [{
    id: 0,
    sender: '767110505@qq.com',
    subject: 'Hi there, old friend',
    date: 'Dec 7, 2013 12:32:00',
    recipients: ['greg@somecompany.com'],
    message: 'Hey, we should get together for lunch sometime and catch up. There are many things we should collaborate on this year.'
}, {
    id: 1,
    sender: 'hackerzgz@qq.com',
    subject: 'Where did you leave my laptop?',
    date: 'Dec 7, 2013 13:15:12',
    recipients: ['greg@somecompany.com'],
    message: 'I thought you were going to put it in my deck drawer. But it does not seem to be there'
}, {
    id: 2,
    sender: 'hackerzgz@gmail.com',
    subject: 'Lost Python',
    date: 'Dec 6, 2013 20:35:02',
    recipients: ['greg@somecompany.com'],
    message: 'Nobody panic, but my pet python is missing from her cage. She doesn\'t move to fast, so just call me if you see her.'
}, ];

function ListController($scope) {
    $scope.messages = messages;
}

function DetailController($scope, $routeParams) {
    $scope.message = messages[$routeParams.id];
}