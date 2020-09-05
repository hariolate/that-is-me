<script>
    const logoStatic = require('./assets/logo_static.png');
    const sprite = {
        cat02: {
            down: [
                require('./assets/sprite/cat02/cat2_01.jpg'),
                require('./assets/sprite/cat02/cat2_02.jpg'),
                require('./assets/sprite/cat02/cat2_03.jpg'),
                require('./assets/sprite/cat02/cat2_04.jpg'),
            ],
            left: [
                require('./assets/sprite/cat02/cat2_05.jpg'),
                require('./assets/sprite/cat02/cat2_06.jpg'),
                require('./assets/sprite/cat02/cat2_07.jpg'),
                require('./assets/sprite/cat02/cat2_08.jpg'),
            ],
            up: [
                require('./assets/sprite/cat02/cat2_09.jpg'),
                require('./assets/sprite/cat02/cat2_10.jpg'),
                require('./assets/sprite/cat02/cat2_11.jpg'),
                require('./assets/sprite/cat02/cat2_12.jpg'),
            ],
            right: [
                require('./assets/sprite/cat02/cat2_13.jpg'),
                require('./assets/sprite/cat02/cat2_14.jpg'),
                require('./assets/sprite/cat02/cat2_15.jpg'),
                require('./assets/sprite/cat02/cat2_16.jpg'),
            ],
        },
    };
    export default {
        name: "common",
        logoStatic,
        sprite,
    }
</script>