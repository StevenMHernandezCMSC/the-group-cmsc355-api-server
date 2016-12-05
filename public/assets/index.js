var frames = [
    'startup-login',
    'redapples',
    'greenapples',
    'greenapples-fast',
    'applecloseups',
    'optionsmenu',
    'userstats',
    'highscores',
    'resumegame',
    'quitgame'
];

var currentFrame = -1;

$(document).keydown(function (e) {
    switch (e.which) {
        case 37: // left
            if (currentFrame > 0) {
                currentFrame--;
                changeVideo();
            }
            break;
        case 38: // up
            $("#project-info").show();
            break;
        case 39: // right
            if (currentFrame < frames.length - 1) {
                currentFrame++;
                changeVideo();
            }
            break;
        case 40: // down
            $("#project-info").hide();
            break;
        default:
            return;
    }
    e.preventDefault();
});

function changeVideo() {
    $("#project-info").hide();

    $('body').vide({
        mp4: 'http://snake-go-assets.shmah.com/' + frames[currentFrame]
    });
}