import React, { useState } from 'react';
import './App.css';
import { MathJax, MathJaxContext } from 'better-react-mathjax';

function App() {
    const [integralMethod, setIntegralMethod] = useState('left rectangle');
    const [integralEquation, setIntegralEquation] = useState('');
    const [lowerBound, setLowerBound] = useState('');
    const [upperBound, setUpperBound] = useState('');
    const [steps, setSteps] = useState('');
    const [result, setResult] = useState(null);

    const integralMethods = [
        "left rectangle",
        "right rectangle",
        "midpoint rectangle",
        "trapezoidal",
        "simpson",
        "monte carlo",
        "gauss lejandre",
        "chebyshev"
    ];

    const handleIntegralSubmit = async (event) => {
        event.preventDefault();
        const requestBody = {
            equationType: integralMethod,
            expression: integralEquation,
            args: [parseFloat(lowerBound), parseFloat(upperBound), parseFloat(steps)]
        };

        try {
            const response = await fetch('/calculate-integral', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(requestBody)
            });
            const data = await response.json();
            setResult(data.result);
        } catch (error) {
            console.error('Error:', error);
        }
    };

    const renderLatexEquation = () => {
        return `\\int_{${lowerBound}}^{${upperBound}} ${integralEquation} \\, dx`;
    };

    return (
        <MathJaxContext>
            <div className="container">
                <h1>Математические Уравнения</h1>

                <h2>Интегральные Уравнения</h2>
                <form onSubmit={handleIntegralSubmit}>
                    <label>Метод:</label>
                    <select value={integralMethod} onChange={(e) => setIntegralMethod(e.target.value)}>
                        {integralMethods.map((method) => (
                            <option key={method} value={method}>{method}</option>
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
                        <MathJax>{`$$${renderLatexEquation()}$$`}</MathJax>
                    </div>
                    <button type="submit">Рассчитать</button>
                </form>
                {result !== null && (
                    <div className="result">
                        <h3>Результат: {result}</h3>
                    </div>
                )}

                <h2>Дифференциальные Уравнения</h2>
                <form onSubmit={(event) => event.preventDefault()}>
                    <label>Уравнение:</label>
                    <input type="text" required />
                    <button type="submit">Рассчитать</button>
                </form>

                <h2>Модель Хищник-Жертва</h2>
                <form onSubmit={(event) => event.preventDefault()}>
                    <label>Уравнение для жертвы:</label>
                    <input type="text" required />
                    <label>Уравнение для хищника:</label>
                    <input type="text" required />
                    <button type="submit">Рассчитать</button>
                </form>
            </div>
        </MathJaxContext>
    );
}

export default App;
