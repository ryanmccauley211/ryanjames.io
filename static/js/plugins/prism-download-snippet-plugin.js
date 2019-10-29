Prism.plugins.toolbar.registerButton('select-code', function (env) {
    var button = document.createElement('button');
    button.innerHTML = 'Select Code';

    button.addEventListener('click', function () {
        console.log("clicked")
    });
});