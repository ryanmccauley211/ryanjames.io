$('.main-menu-item a').each(function () {
    $(this).on('click', function (event) {
        event.preventDefault();
        pageTarget = $(this).attr("href");
        transitionToPage(pageTarget);
    })
});

function transitionToPage(pageTarget) {
    if (pageTarget === "explore") {

        $('.svg-name-text').fadeOut(250, function () {
            fadeBackgroundColorTo($(".page-half-partition-right"), 'white', 1000);
            fontShrink($(".main-menu-item a"), 0, 1000);
            flexGrow($(".page-half-partition-left"), 0, 750, function () {
                window.location = "explore";
                window.history.pushState(stateObj, 'Explore', '/explore');
            });
        });
    }

    else if (pageTarget === "portfolio") {

        $('.svg-name-text').fadeOut(250, function () {
            fadeBackgroundColorTo($(".page-half-partition-right"), 'white', 1000);
            fontShrink($(".main-menu-item a"), 0, 1000);
            flexGrow($(".page-half-partition-left"), 0, 750, function () {
                window.location = "portfolio";
                window.history.pushState(stateObj, 'Portfolio', '/portfolio');
            });
        });
    }

    else if (pageTarget === "contact") {

        $('.svg-name-text').fadeOut(250, function () {
            fadeBackgroundColorTo($(".page-half-partition-right"), 'white', 1000);
            fontShrink($(".main-menu-item a"), 0, 1000);
            flexGrow($(".page-half-partition-left"), 0, 750, function () {
                window.location = "contact";
                window.history.pushState(stateObj, 'Contact', '/contact');
            });
        });
    }

}