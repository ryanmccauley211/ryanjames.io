$('.main-menu-item a').each(function () {
    $(this).on('click', function (event) {
        event.preventDefault();
        pageTarget = $(this).attr("href");
        transitionToPage(pageTarget);
    })
});

function transitionToPage(pageTarget) {
    if (pageTarget === "explore") {
        fadeBackgroundColorTo($(".page-half-partition-right"), 'white', 2000);
        fontShrink($(".main-menu-item a"), 0, 1000);
        resizeSvg($('#name-svg'), 0, 0, 1000);
        flexGrow($(".page-half-partition-left"), 0, 1000, function () {
            window.location = "explore";
            window.history.pushState(stateObj, 'Explore', '/explore');
        });
    }

    else if (pageTarget === "portfolio") {
        fadeBackgroundColorTo($(".page-half-partition-right"), 'white', 2000);
        fontShrink($(".main-menu-item a"), 0, 1000);
        resizeSvg($('#name-svg'), 0, 0, 1000);
        flexGrow($(".page-half-partition-left"), 0, 1000, function () {
            window.location = "portfolio";
            window.history.pushState(stateObj, 'Portfolio', '/portfolio');
        });
    }

    else if (pageTarget === "contact") {
        fadeBackgroundColorTo($(".page-half-partition-right"), 'white', 2000);
        fontShrink($(".main-menu-item a"), 0, 1000);
        resizeSvg($('#name-svg'), 0, 0, 1000);
        flexGrow($(".page-half-partition-left"), 0, 1000, function () {
            window.location = "contact";
            window.history.pushState(stateObj, 'Contact', '/contact');
        });
    }

}