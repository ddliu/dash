var Notification = require('node-notifier');

var notifier = new Notification();
notifier.notify({
    message: 'Hello World'
});