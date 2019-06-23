document.addEventListener('DOMContentLoaded', () => {
    var path = document.querySelector('.movie');

    path.addEventListener('click', () => {
        console.log(document.querySelector('.movie').value);
    });

    console.log(document.querySelector('.test_go'));
});