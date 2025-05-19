function setTheme(theme) {
  document.documentElement.setAttribute('data-theme', theme);
  localStorage.setItem('theme', theme);
  updateThemeIcon(theme);
  window.dispatchEvent(new Event('themechange'));
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

// Минималистичная плавная смена темы
function animateThemeFade(nextTheme) {
  document.body.classList.add('theme-fade');
  setTimeout(() => {
    setTheme(nextTheme);
    document.body.classList.remove('theme-fade');
  }, 300);
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
      animateThemeFade(next);
      setTimeout(() => btn.classList.remove('open'), 500);
    };
  });
});
