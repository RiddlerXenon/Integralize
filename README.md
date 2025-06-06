<div align="center">
  <img src="static/images/logo.svg" width="160" alt="Integralize Logo"/>
  <h1>Integralize</h1>
  <br>
  <b>Веб‑платформа для математического анализа</b>
  <br><br>
  <i>
    Проект разработан в Курском государственном университете<br>
    в рамках учебной деятельности
  </i>
  <br><br>
</div>

---

<p align="center">
  <img src="https://img.shields.io/badge/Go-%2347d9c4.svg?style=flat&logo=go&logoColor=white"/>
  <img src="https://img.shields.io/badge/HTML5-%23e34c26.svg?style=flat&logo=html5&logoColor=white"/>
  <img src="https://img.shields.io/badge/CSS3-%231572b6.svg?style=flat&logo=css3&logoColor=white"/>
  <img src="https://img.shields.io/badge/JavaScript-%23f7df1e.svg?style=flat&logo=javascript&logoColor=black"/>
</p>

---

## ✨ О проекте

**Integralize** — современная open-source платформа для студентов, преподавателей и исследователей.  
Позволяет удобно и наглядно решать задачи по математическому анализу:  
- вычислять интегралы различными методами,
- решать дифференциальные уравнения,
- строить графики и анализировать результаты.

Платформа реализована на Go с использованием HTML/CSS/JavaScript для фронтенда.

---

## 🚀 Основные возможности

- **Калькулятор интегралов**  
  Классические и современные численные методы: прямоугольников, Симпсона, трапеций, Монте-Карло, Гаусса-Лежандра, Чебышёва.
- **Решение дифференциальных уравнений**  
  Методы Эйлера, Рунге-Кутта, численное моделирование динамики (например, модель «хищник-жертва»).
- **Современный веб-интерфейс**  
  Быстрый ввод формул, мгновенный вывод результатов, наглядные графики.
- **REST API**  
  Гибкие маршруты для интеграции в ваши приложения и автоматизации расчётов.

---

## 🏁 Быстрый старт

```bash
git clone https://github.com/RiddlerXenon/Integralize.git
cd Integralize
go mod tidy
go run ./cmd/Integralize
```

Перейдите в браузере по адресу: [http://localhost:8080](http://localhost:8080)

---

## 📚 Примеры использования

- **Вычисление интегралов:**  
  Введите функцию, выберите метод — получите результат и график.
- **Решение дифференциальных уравнений:**  
  Введите параметры, выберите метод и получите численное решение.
- **Автоматизация через API:**  
  Используйте REST API для интеграции в свои сервисы (см. internal/routes/api.go).

---

## 🤝 Авторы и благодарности

- [RiddlerXenon](https://github.com/RiddlerXenon)
- [Tokarevi4](https://github.com/Tokarevi4)

Особая благодарность преподавателям и студентам Курского государственного университета за идеи, поддержку и тестирование!

---

## 💬 Вклад в проект

Pull requests и идеи приветствуются!  
Нашли ошибку или есть идея для улучшения? — Открывайте Issue или создавайте Pull Request.

---

<div align="center">
  <sub>С любовью к математике и открытым знаниям.<br>
  Курский государственный университет, 2025</sub>
</div>
