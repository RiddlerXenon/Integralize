import React, { useState } from 'react';
import './App.css';
import { MathJax, MathJaxContext } from 'better-react-mathjax';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import { Line } from 'react-chartjs-2';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { parse } from 'mathjs';

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
);

function App() {
    const [integralMethod, setIntegralMethod] = useState('left rectangle');
    const [integralEquation, setIntegralEquation] = useState('');
    const [lowerBound, setLowerBound] = useState('');
    const [upperBound, setUpperBound] = useState('');
    const [steps, setSteps] = useState('');
    const [result, setResult] = useState(null);

    const [differentialMethod, setDifferentialMethod] = useState('euler');
    const [differentialEquation, setDifferentialEquation] = useState('');
    const [y0, setY0] = useState('');
    const [x0, setX0] = useState('');
    const [diffData, setDiffData] = useState(null);

    const [preyEquation, setPreyEquation] = useState('');
    const [predatorEquation, setPredatorEquation] = useState('');

    const integralMethods = [
        { value: "left rectangle", label: "Левый прямоугольник" },
        { value: "right rectangle", label: "Правый прямоугольник" },
        { value: "midpoint rectangle", label: "Средний прямоугольник" },
        { value: "trapezoidal", label: "Трапеция" },
        { value: "simpson", label: "Симпсон" },
        { value: "monte carlo", label: "Монте Карло" },
        { value: "gauss lejandre", label: "Гаусс Лежандр" },
        { value: "chebyshev", label: "Чебышев" }
    ];

    const differentialMethods = [
        { value: "euler", label: "Эйлер" },
        { value: "runge-kutta", label: "Рунге-Кутта" }
    ];

    const toLatex = (expression) => {
        try {
            const node = parse(expression);
            return node.toTex({ parenthesis: 'keep', implicit: 'show' });
        } catch (error) {
            console.error('Ошибка при преобразовании в LaTeX:', error);
            return expression;
        }
    };

    const handleIntegralSubmit = async (event) => {
        event.preventDefault();
        const latexExpression = toLatex(integralEquation);
        const requestBody = {
            equationType: integralMethod,
            expression: latexExpression,
            args: [parseFloat(lowerBound), parseFloat(upperBound), parseFloat(steps)]
        };

        console.log('Отправка запроса:', requestBody);

        try {
            const response = await fetch('http://127.0.0.1:8080/api/integral', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(requestBody)
            });
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            console.log('Ответ от сервера:', data);
            setResult(data.result);
        } catch (error) {
            console.error('Ошибка:', error);
        }
    };

    const handleDifferentialSubmit = async (event) => {
        event.preventDefault();
        const latexExpression = toLatex(differentialEquation);
        const tMax = 10;
        const h = tMax / 10;
        const requestBody = {
            equationType: differentialMethod,
            expression: latexExpression,
            args: [parseFloat(y0), parseFloat(x0), tMax, h]
        };

        console.log('Отправка запроса:', requestBody);

        try {
            const response = await fetch('http://127.0.0.1:8080/api/differential', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(requestBody)
            });
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            console.log('Ответ от сервера:', data);
            setDiffData(data);
        } catch (error) {
            console.error('Ошибка:', error);
        }
    };

    const handlePredatorPreySubmit = (event) => {
        event.preventDefault();
        console.log(`Уравнение для жертвы: ${preyEquation}, Уравнение для хищника: ${predatorEquation}`);
        // Реализуйте запрос для модели Хищник-Жертва
    };

    const renderIntegralLatexEquation = () => {
        return `\\int_{${lowerBound}}^{${upperBound}} ${toLatex(integralEquation)} \\, dx`;
    };

    const renderDifferentialLatexEquation = () => {
        return `${toLatex(differentialEquation)}, \\ y(${x0}) = ${y0}`;
    };

    const renderDiffChart = () => {
        if (!diffData) return null;

        const data = {
            labels: diffData.x,
            datasets: [
                {
                    label: 'Решение дифференциального уравнения',
                    data: diffData.y,
                    fill: false,
                    backgroundColor: 'rgba(75,192,192,0.2)',
                    borderColor: 'rgba(75,192,192,1)',
                },
            ],
        };

        const options = {
            scales: {
                x: {
                    title: {
                        display: true,
                        text: 'x',
                    },
                },
                y: {
                    title: {
                        display: true,
                        text: 'y',
                    },
                },
            },
        };

        return <Line data={data} options={options} />;
    };

    return (
        <MathJaxContext>
            <div className="container">
                <h1>Математические Уравнения</h1>
                <Tabs>
                    <TabList>
                        <Tab>Интегралы</Tab>
                        <Tab>Дифференциальные Уравнения</Tab>
                        <Tab>Модель Хищник-Жертва</Tab>
                    </TabList>

                    <TabPanel>
                        <h2>Интегральные Уравнения</h2>
                        <form onSubmit={handleIntegralSubmit}>
                            <label>Метод:</label>
                            <select value={integralMethod} onChange={(e) => setIntegralMethod(e.target.value)}>
                                {integralMethods.map((method) => (
                                    <option key={method.value} value={method.value}>{method.label}</option>
                                ))}
                            </select>
                            <label>Подынтегральное уравнение:</label>
                            <input type="text" value={integralEquation} onChange={(e) => setIntegralEquation(e.target.value)} required />
                            <label>Нижняя граница:</label>
                            <input type="number" value={lowerBound} onChange={(e) => setLowerBound(e.target.value)} required />
                            <label>Верхняя граница:</label>
                            <input type="number" value={upperBound} onChange={(e) => setUpperBound(e.target.value)} required />
                            <label>Количество шагов:</label>
                            <input type="number" value={steps} onChange={(e) => setSteps(e.target.value)} required />
                            <div className="latex-equation">
                                <MathJax>{`$$${renderIntegralLatexEquation()}$$`}</MathJax>
                            </div>
                            <button type="submit">Рассчитать</button>
                        </form>
                        {result !== null && (
                            <div className="result">
                                <h3>Результат: {result}</h3>
                            </div>
                        )}
                    </TabPanel>

                    <TabPanel>
                        <h2>Дифференциальные Уравнения</h2>
                        <form onSubmit={handleDifferentialSubmit}>
                            <label>Метод:</label>
                            <select value={differentialMethod} onChange={(e) => setDifferentialMethod(e.target.value)}>
                                {differentialMethods.map((method) => (
                                    <option key={method.value} value={method.value}>{method.label}</option>
                                ))}
                            </select>
                            <label>Уравнение:</label>
                            <input type="text" value={differentialEquation} onChange={(e) => setDifferentialEquation(e.target.value)} required />
                            <label>y0 - начальное значение:</label>
                            <input type="number" value={y0} onChange={(e) => setY0(e.target.value)} required />
                            <label>x0 - начальное время:</label>
                            <input type="number" value={x0} onChange={(e) => setX0(e.target.value)} required />
                            <div className="latex-equation">
                                <MathJax>{`$$${renderDifferentialLatexEquation()}$$`}</MathJax>
                            </div>
                            <button type="submit">Рассчитать</button>
                        </form>
                        {diffData && renderDiffChart()}
                    </TabPanel>

                    <TabPanel>
                        <h2>Модель Хищник-Жертва</h2>
                        <form onSubmit={handlePredatorPreySubmit}>
                            <label>Уравнение для жертвы:</label>
                            <input type="text" value={preyEquation} onChange={(e) => setPreyEquation(e.target.value)} required />
                            <label>Уравнение для хищника:</label>
                            <input type="text" value={predatorEquation} onChange={(e) => setPredatorEquation(e.target.value)} required />
                            <button type="submit">Рассчитать</button>
                        </form>
                    </TabPanel>
                </Tabs>
            </div>
        </MathJaxContext>
    );
}

export default App;
