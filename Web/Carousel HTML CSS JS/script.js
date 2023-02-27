const carousel = document.querySelector('.carousel');
let currdeg = 0;

document.querySelector('.next').addEventListener('click', function() {
    currdeg -= 60;
    carousel.style.transform = `rotateY(${currdeg}deg)`;
});

document.querySelector('.prev').addEventListener('click', function() {
    currdeg += 60;
    carousel.style.transform = `rotateY(${currdeg}deg)`;
});
