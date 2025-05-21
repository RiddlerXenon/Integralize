// Улучшенный график с поддержкой темной темы

function getGraphThemeColors() {
  const root = document.documentElement;
  const isDark = root.getAttribute('data-theme') === 'dark';
  return isDark
    ? {
        gradTop: "rgba(107,178,253,0.18)",
        gradBottom: "rgba(107,178,253,0.01)",
        grid: "#2c3652",
        axis: "#6bb2fd",
        label: "#e4e7ef",
        curve: "#6bb2fd",
        point: "#4adedb",
        pointShadow: "#6bb2fd",
        pointBorder: "#23272a"
      }
    : {
        gradTop: "rgba(66,103,233,0.18)",
        gradBottom: "rgba(66,103,233,0.01)",
        grid: "#d5e0f7",
        axis: "#4267e9",
        label: "#4267e9",
        curve: "#4267e9",
        point: "#4adedb",
        pointShadow: "#4267e9",
        pointBorder: "#fff"
      };
}

function drawGraph(X, Y) {
  try {
    // console.log("[drawGraph] Входные массивы:", X, Y);
    const container = document.getElementById('diffeq-graph-container');
    const canvas = document.getElementById('diffeq-graph');
    const tooltip = document.getElementById('graph-tooltip');
    container.classList.add('active');
    canvas.style.display = 'block';

    const ctx = canvas.getContext('2d');
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    const colors = getGraphThemeColors();

    // Очистка и проверка
    let clean = [];
    for (let i = 0; i < X.length; i++) {
      if (isFinite(X[i]) && isFinite(Y[i])) clean.push([X[i], Y[i]]);
    }
    // console.log("[drawGraph] После очистки конечных точек:", clean);

    if (clean.length < 2) {
      // console.warn("[drawGraph] Недостаточно точек для графика!");
      return;
    }
    X = clean.map(p => p[0]);
    Y = clean.map(p => p[1]);

    // Масштаб
    const W = canvas.width, H = canvas.height;
    const YLIMIT = 1e6;
    let maxYabs = Math.max(...Y.map(y => Math.abs(y)));
    let filtered = Y;
    if (maxYabs > YLIMIT) {
      filtered = Y.map(y => Math.abs(y) > YLIMIT ? Math.sign(y) * YLIMIT : y);
      // console.log("[drawGraph] Обрезка больших значений Y");
    }
    let minX = Math.min(...X), maxX = Math.max(...X);
    let minY = Math.min(...filtered), maxY = Math.max(...filtered);
    let pad = 0.10;
    minX -= (maxX - minX) * pad; maxX += (maxX - minX) * pad;
    minY -= (maxY - minY) * pad; maxY += (maxY - minY) * pad;
    if (minY === maxY) { minY -= 1; maxY += 1; }
    if (minX === maxX) { minX -= 1; maxX += 1; }
    function sx(x) { return 55 + (x - minX) / (maxX - minX) * (W - 90); }
    function sy(y) { return H - 40 - (y - minY) / (maxY - minY) * (H - 70); }

    // Градиент под графиком
    ctx.save();
    ctx.beginPath();
    ctx.moveTo(sx(X[0]), sy(filtered[0]));
    for (let i = 1; i < X.length; i++) {
      ctx.lineTo(sx(X[i]), sy(filtered[i]));
    }
    ctx.lineTo(sx(X[X.length - 1]), H - 40);
    ctx.lineTo(sx(X[0]), H - 40);
    ctx.closePath();
    let grad = ctx.createLinearGradient(0, sy(Math.max(...filtered)), 0, H - 40);
    grad.addColorStop(0, colors.gradTop);
    grad.addColorStop(1, colors.gradBottom);
    ctx.fillStyle = grad;
    ctx.fill();
    ctx.restore();

    // Сетка
    ctx.save();
    ctx.strokeStyle = colors.grid;
    ctx.lineWidth = 1;
    ctx.setLineDash([2, 10]);
    let gridCountX = 5, gridCountY = 6;
    for (let i = 0; i <= gridCountX; i++) {
      let x = minX + (maxX - minX) * i / gridCountX;
      let px = sx(x);
      ctx.beginPath(); ctx.moveTo(px, 35); ctx.lineTo(px, H - 30); ctx.stroke();
    }
    for (let i = 0; i <= gridCountY; i++) {
      let y = minY + (maxY - minY) * i / gridCountY;
      let py = sy(y);
      ctx.beginPath(); ctx.moveTo(55, py); ctx.lineTo(W - 35, py); ctx.stroke();
    }
    ctx.setLineDash([]);
    ctx.restore();

    // Оси
    ctx.save();
    ctx.strokeStyle = colors.axis;
    ctx.lineWidth = 2.5;
    ctx.beginPath(); ctx.moveTo(sx(minX), sy(0)); ctx.lineTo(sx(maxX), sy(0)); ctx.stroke();
    ctx.beginPath(); ctx.moveTo(sx(0), sy(minY)); ctx.lineTo(sx(0), sy(maxY)); ctx.stroke();
    ctx.restore();

    // Подписи делений
    ctx.save();
    ctx.fillStyle = colors.label;
    ctx.font = "bold 14px 'Inter', Arial, sans-serif";
    ctx.textAlign = "center";
    ctx.textBaseline = "top";
    for (let i = 0; i <= gridCountX; i++) {
      let x = minX + (maxX - minX) * i / gridCountX;
      let px = sx(x);
      ctx.fillText(Number(x.toFixed(2)), px, H - 25);
    }
    ctx.textAlign = "right";
    ctx.textBaseline = "middle";
    for (let i = 0; i <= gridCountY; i++) {
      let y = minY + (maxY - minY) * i / gridCountY;
      let py = sy(y);
      ctx.fillText(Number(y.toFixed(2)), 50, py);
    }
    ctx.restore();

    // График кривой (плавный)
    ctx.save();
    ctx.strokeStyle = colors.curve;
    ctx.lineWidth = 3.3;
    ctx.beginPath();
    ctx.moveTo(sx(X[0]), sy(filtered[0]));
    for (let i = 1; i < X.length; i++) {
      const xc = (sx(X[i - 1]) + sx(X[i])) / 2;
      const yc = (sy(filtered[i - 1]) + sy(filtered[i])) / 2;
      ctx.quadraticCurveTo(sx(X[i - 1]), sy(filtered[i - 1]), xc, yc);
    }
    ctx.lineTo(sx(X[X.length - 1]), sy(filtered[X.length - 1]));
    ctx.stroke();
    ctx.restore();

    // Точки
    ctx.save();
    for (let i = 0; i < X.length; i++) {
      ctx.beginPath();
      ctx.arc(sx(X[i]), sy(filtered[i]), 6, 0, 2 * Math.PI);
      ctx.fillStyle = colors.point;
      ctx.shadowColor = colors.pointShadow;
      ctx.shadowBlur = 7;
      ctx.fill();
      ctx.shadowBlur = 0;
      ctx.strokeStyle = colors.pointBorder;
      ctx.lineWidth = 1.5;
      ctx.stroke();
    }
    ctx.restore();

    // Наведение на точки
    const points = X.map((x, i) => ({
      x: sx(x),
      y: sy(filtered[i]),
      origX: x,
      origY: filtered[i],
      idx: i
    }));

    canvas.onmousemove = function (e) {
      const rect = canvas.getBoundingClientRect();
      const mx = e.clientX - rect.left;
      const my = e.clientY - rect.top;
      let found = null;
      let minDist = 16;
      for (const pt of points) {
        const dx = pt.x - mx, dy = pt.y - my;
        const dist = Math.sqrt(dx * dx + dy * dy);
        if (dist < minDist) {
          minDist = dist;
          found = pt;
        }
      }
      if (found) {
        tooltip.style.opacity = 1;
        tooltip.textContent = `x = ${Number(found.origX.toFixed(4))}, y = ${Number(found.origY.toFixed(4))}`;
        const tooltipRect = tooltip.getBoundingClientRect();
        const canvasRect = canvas.getBoundingClientRect();
        let left = found.x + canvas.offsetLeft + 18;
        let top = found.y + canvas.offsetTop - 20;
        const tooltipWidth = tooltipRect.width || 120;
        const tooltipHeight = tooltipRect.height || 30;
        if (left + tooltipWidth > canvas.width) left = canvas.width - tooltipWidth - 5;
        if (left < 0) left = 5;
        if (top < 0) top = 5;
        if (top + tooltipHeight > canvas.height) top = canvas.height - tooltipHeight - 5;
        tooltip.style.left = left + "px";
        tooltip.style.top = top + "px";
      } else {
        tooltip.style.opacity = 0;
      }
    };
    canvas.onmouseleave = function () {
      tooltip.style.opacity = 0;
    };
    // console.log("[drawGraph] График успешно построен");
  } catch (err) {
    // console.error("[drawGraph] Ошибка при построении графика:", err);
    throw err; // для отлова в основном обработчике
  }
}

// Поддержка смены темы "на лету"
if (!window._diffeq_theme_listener) {
  window._diffeq_theme_listener = true;
  window.addEventListener('themechange', () => {
    if (window._lastDiffeqX && window._lastDiffeqY) {
      drawGraph(window._lastDiffeqX, window._lastDiffeqY);
    }
  });
}

document.getElementById('diffeq-form').addEventListener('submit', async function (e) {
  e.preventDefault();

  const form = e.target;
  const x0 = parseFloat(form.x0.value);
  const y0 = parseFloat(form.y0.value);
  const tMax = parseFloat(form.tMax.value);
  const h = parseFloat(form.h.value);

  let error = "";

  if (isNaN(x0) || x0 < -100 || x0 > 100) {
    error += "x₀ должен быть в диапазоне от -100 до 100\n";
  }
  if (isNaN(y0) || y0 < -100 || y0 > 100) {
    error += "y₀ должен быть в диапазоне от -100 до 100\n";
  }
  if (isNaN(h) || h <= 0 || h > 10) {
    error += "Шаг должен быть строго больше 0 и не больше 10\n";
  }
  if (isNaN(tMax) || tMax < x0 + h) {
    error += "Максимальный x должен быть не меньше x₀ + шаг\n";
  }
  if (tMax > x0 + h + 100) {
    error += "Максимальный x не может быть больше x₀ + шаг + 100\n";
  }

  if (error) {
    alert(error);
    return;
  }

  const resultBlock = document.getElementById('diffeq-result');
  const container = document.getElementById('diffeq-graph-container');
  const canvas = document.getElementById('diffeq-graph');
  document.getElementById('graph-tooltip').style.opacity = 0;
  resultBlock.classList.remove('active');
  resultBlock.innerHTML = '';
  resultBlock.style.display = 'none';
  container.classList.remove('active');
  canvas.style.display = "none";

  const method = form.method.value;
  const expression = form.expression.value;

  const body = JSON.stringify({
    equationType: method,
    expression: expression,
    args: [y0, x0, tMax, h]
  });

  try {
    // console.log("[submit] Отправка запроса:", body);
      const response = await fetch('https://integralize.ru/api/differential', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body
      });

      if (!response.ok) {
        // console.error("[submit] HTTP ошибка:", response.status);
        throw new Error('Ошибка HTTP: ' + response.status);
      }

      const text = await response.text();
      // console.log("[submit] RAW ответ сервера:", text);

      let data;
      try {
        data = text ? JSON.parse(text) : null;
      } catch (err) {
        // console.error("[submit] Ошибка парсинга JSON:", err);
        resultBlock.classList.add('active');
        resultBlock.style.display = 'flex';
        resultBlock.innerHTML = '<div class="result error">Некорректный ответ от сервера (не JSON).</div>';
        return;
      }

      // дальше обычная обработка data...
      const X = data && (data.x || data.X);
      const Y = data && (data.y || data.Y);
    function parseNumber(val) {
      if (val === "+Inf" || val === "Inf" || val === "Infinity") return Infinity;
      if (val === "-Inf" || val === "-Infinity") return -Infinity;
      if (typeof val === "string" && !isNaN(Number(val))) return Number(val);
      if (typeof val === "number") return val;
      return NaN;
    }

    const Xnum = Array.isArray(X) ? X.map(parseNumber) : [];
    const Ynum = Array.isArray(Y) ? Y.map(parseNumber) : [];

    // Фильтруем пары (x, y), оставляя только конечные значения
    let cleanX = [];
    let cleanY = [];
    for (let i = 0; i < Xnum.length; i++) {
      if (isFinite(Xnum[i]) && isFinite(Ynum[i])) {
        cleanX.push(Xnum[i]);
        cleanY.push(Ynum[i]);
      }
    }
    // console.log("[submit] После фильтрации:", cleanX, cleanY);

    if (cleanX.length < 2) {
      resultBlock.classList.add('active');
      resultBlock.style.display = 'flex';
      resultBlock.innerHTML = '<div class="result error">Решение ушло в бесконечность или недостаточно корректных значений для построения графика.<br>Попробуйте уменьшить шаг или диапазон.</div>';
      // console.warn("[submit] Недостаточно точек для графика, график не строится");
      return;
    }

    resultBlock.classList.add('active');
    resultBlock.style.display = 'flex';
    resultBlock.innerHTML = '';

    window._lastDiffeqX = cleanX;
    window._lastDiffeqY = cleanY;

    try {
      drawGraph(cleanX, cleanY);
    } catch (err) {
      // Если drawGraph упал — логируем и показываем ошибку пользователю
      // console.error("[submit] Ошибка при вызове drawGraph:", err);
      resultBlock.innerHTML = '<div class="result error">Внутренняя ошибка при построении графика.</div>';
    }

  } catch (err) {
    // console.error("[submit] Ошибка основного блока:", err);
    resultBlock.classList.add('active');
    resultBlock.style.display = 'flex';
    resultBlock.innerHTML = '<div class="result error">Ошибка соединения с сервером или парсинга данных.</div>';
  }
});
