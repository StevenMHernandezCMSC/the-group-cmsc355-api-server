$.get('/slideshow/frames', function(frames) {
    var currentFrame = -1;

    var $projectInfo = $('#project-info');
    var $caption = $('#caption');

    $(document).keydown(function (e) {
        switch (e.which) {
            case 37: // left
                if (currentFrame > 0) {
                    currentFrame--;
                    changeVideo();
                }
                break;
            case 38: // up
                $projectInfo.show();
                $caption.hide();
                break;
            case 39: // right
                if (currentFrame < frames.length - 1) {
                    currentFrame++;
                    changeVideo();
                } else {
                    $projectInfo.show();
                    $caption.hide();
                }
                break;
            case 40: // down
                $projectInfo.hide();
                $caption.show();
                break;
            default:
                return;
        }
        e.preventDefault();
    });

    function changeVideo() {
        $projectInfo.hide();

        $('body').vide({
            mp4: 'http://snake-go-assets.shmah.com/' + frames[currentFrame]['Media'],
            poster: 'http://snake-go-assets.shmah.com/' + frames[currentFrame]['Media']
        }, {
            posterType: 'jpg'
        });

        $caption.html(frames[currentFrame]['Caption']);
        $caption.show();
    }
});