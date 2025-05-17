function setTheme(theme) {
  document.documentElement.setAttribute('data-theme', theme);
  localStorage.setItem('theme', theme);
  updateThemeIcon(theme);
}

function getTheme() {
  return localStorage.getItem('theme') ||
    (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light');
}

function updateThemeIcon(theme) {
  const icon = document.querySelector('.theme-fab .theme-icon');
  if (!icon) return;
  icon.classList.toggle('moon', theme === 'dark');
  icon.classList.toggle('sun', theme !== 'dark');
}

function getWaveBg(theme) {
  // Цвет волны — как фон новой темы
  return theme === 'dark' ? '#181b20' : '#fff';
}

function animateThemeWave(nextTheme) {
  // Центр волны — правый нижний угол
  const vw = Math.max(window.innerWidth, document.documentElement.clientWidth);
  const vh = Math.max(window.innerHeight, document.documentElement.clientHeight);

  // Максимальное расстояние до левого верхнего угла
  const maxRadius = Math.hypot(vw, vh);

  // Создаём волну
  let wave = document.createElement('div');
  wave.className = 'theme-wave';
  wave.style.setProperty('--wave-bg', getWaveBg(nextTheme));
  // scale, чтобы круг "дотянулся" до левого верхнего угла
  const scale = Math.ceil((maxRadius * 2) / 2);
  wave.style.setProperty('--wave-scale', scale);

  document.body.appendChild(wave);

  // Запустить анимацию расширения волны
  setTimeout(() => {
    wave.classList.add('expand');
  }, 16);

  // В момент, когда волна покрывает весь экран — меняем тему
  setTimeout(() => setTheme(nextTheme), 400);

  // Плавно убираем волну
  setTimeout(() => {
    wave.classList.add('fade');
  }, 780);

  // Удалить волну после завершения
  setTimeout(() => {
    wave.remove();
  }, 1100);
}

document.addEventListener('DOMContentLoaded', function() {
  // Добавить кнопку, если её нет
  if (!document.querySelector('.theme-fab')) {
    let btn = document.createElement('button');
    btn.className = 'theme-fab';
    btn.type = 'button';
    btn.innerHTML = '<span class="theme-icon"></span>';
    document.body.appendChild(btn);
  }
  setTheme(getTheme());
  document.querySelectorAll('.theme-fab').forEach(btn => {
    btn.onclick = function() {
      btn.classList.add('open');
      let current = getTheme();
      let next = current === 'dark' ? 'light' : 'dark';
      animateThemeWave(next);
      setTimeout(() => btn.classList.remove('open'), 700);
    };
  });
});
