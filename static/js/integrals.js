document.getElementById('integral-form').addEventListener('submit', async function(e) {
  e.preventDefault();
  const form = e.target;
  const resultBlock = document.getElementById('integral-result');
  resultBlock.classList.remove('active');
  resultBlock.innerHTML = '';
  resultBlock.style.display = 'none';

  resultBlock.classList.add('active');
  resultBlock.style.display = 'flex';
  resultBlock.innerHTML = '<div class="result loading">Вычисление...</div>';

  const method = form.method.value;
  const expression = form.expression.value;
  const a = parseFloat(form.a.value);
  const b = parseFloat(form.b.value);
  const n = parseFloat(form.n.value);

  const body = JSON.stringify({
    equationType: method,
    expression: expression,
    args: [a, b, n]
  });

  try {
    const response = await fetch('https://integralize.ru/api/integral', {
      method: 'POST',
      headers: {'Content-Type':'application/json'},
      body
    });
    const data = await response.json();
    if (data && data.result !== undefined) {
      resultBlock.classList.add('active');
      resultBlock.style.display = 'flex';

      let value = data.result;
      if (value === "+Inf" || value === "Inf" || value === "Infinity") value = "∞";
      if (value === "-Inf" || value === "-Infinity") value = "-∞";

      resultBlock.innerHTML = `<div class="result"><h2>Результат</h2><div class="value">${value}</div></div>`;
    } else if(data && data.error) {
      resultBlock.classList.add('active');
      resultBlock.style.display = 'flex';
      resultBlock.innerHTML = `<div class="result error">${data.error}</div>`;
    } else {
      resultBlock.classList.add('active');
      resultBlock.style.display = 'flex';
      resultBlock.innerHTML = '<div class="result error">Не удалось получить результат.</div>';
    }
  } catch (err) {
    resultBlock.classList.add('active');
    resultBlock.style.display = 'flex';
    resultBlock.innerHTML = '<div class="result error">Ошибка соединения с сервером.</div>';
  }
});
