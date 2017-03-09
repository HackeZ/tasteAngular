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
    // id:0, sender: '767110505@qq.com', subject: 'Hi there, old friend',date: ''
}]