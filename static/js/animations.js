function flexGrow(element, toVal, duration, callback) {
    element.animate({
        'flex-grow': toVal
    }, duration, function() {
        if (callback != null) {
            callback()
        }
    });
}

function fontShrink(element, toVal, duration) {
    element.animate({
        'font-size': toVal
    }, duration, function() {

    });
}

function fadeBackgroundColorTo(element, toVal, duration) {
    element.animate({backgroundColor: toVal}, duration)
}

function resizeSvg(element, widthToVal, heightToVal, duration) {
    element.animate({
        width: widthToVal,
        height: heightToVal
    }, duration)

}