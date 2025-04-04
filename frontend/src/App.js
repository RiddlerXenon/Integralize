import React, { useState } from 'react';
import './App.css';
import { MathJaxContext } from 'better-react-mathjax';
import CustomTabs from './components/Tabs';
import IntegralForm from './components/IntegralForm';
import DifferentialForm from './components/DifferentialForm';
import PredatorPreyForm from './components/PredatorPreyForm';
import Chart from './components/Chart';
import { toLatex } from './utils/mathUtils';

function App() {
    const [result, setResult] = useState(null);
    const [diffData, setDiffData] = useState(null);

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

    const handleIntegralSubmit = async (event, method, equation, lowerBound, upperBound, steps) => {
        event.preventDefault();
        const latexExpression = toLatex(equation);
        const requestBody = {
            equationType: method,
            expression: latexExpression,
            args: [parseFloat(lowerBound || -1), parseFloat(upperBound || 1), parseFloat(steps || 5)]
        };

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
            setResult(data.result);
        } catch (error) {
            console.error('Ошибка:', error);
        }
    };

    const handleDifferentialSubmit = async (event, method, equation, y0, x0) => {
        event.preventDefault();
        const latexExpression = toLatex(equation);
        const tMax = (parseFloat(x0) || 0) + 10;
        const h = tMax / 25;
        const requestBody = {
            equationType: method,
            expression: latexExpression,
            args: [parseFloat(y0), parseFloat(x0), tMax, h]
        };

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
            setDiffData(data);
        } catch (error) {
            console.error('Ошибка:', error);
        }
    };

    const handlePredatorPreySubmit = (event, preyEquation, predatorEquation) => {
        event.preventDefault();
        console.log(`Уравнение для жертвы: ${preyEquation}, Уравнение для хищника: ${predatorEquation}`);
        // Реализуйте запрос для модели Хищник-Жертва
    };

    return (
        <MathJaxContext>
            <div className="container">
                <h1>Математические Уравнения</h1>
                <CustomTabs
                    renderIntegralForm={() => (
                        <IntegralForm
                            integralMethods={integralMethods}
                            toLatex={toLatex}
                            handleSubmit={handleIntegralSubmit}
                            result={result}
                        />
                    )}
                    renderDifferentialForm={() => (
                        <DifferentialForm
                            differentialMethods={differentialMethods}
                            toLatex={toLatex}
                            handleSubmit={handleDifferentialSubmit}
                            diffData={diffData}
                            renderDiffChart={() => <Chart diffData={diffData} />}
                        />
                    )}
                    renderPredatorPreyForm={() => (
                        <PredatorPreyForm
                            handleSubmit={handlePredatorPreySubmit}
                        />
                    )}
                />
            </div>
        </MathJaxContext>
    );
}

export default App;
