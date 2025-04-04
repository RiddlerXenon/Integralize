import React, { useState } from 'react';
import { MathJax } from 'better-react-mathjax';

const IntegralForm = ({ integralMethods, toLatex, handleSubmit, result }) => {
    const [integralMethod, setIntegralMethod] = useState('left rectangle');
    const [integralEquation, setIntegralEquation] = useState('');
    const [lowerBound, setLowerBound] = useState('');
    const [upperBound, setUpperBound] = useState('');
    const [steps, setSteps] = useState('');
    const [isEquationFocused, setIsEquationFocused] = useState(false);
    const [isLowerBoundFocused, setIsLowerBoundFocused] = useState(false);
    const [isUpperBoundFocused, setIsUpperBoundFocused] = useState(false);
    const [isStepsFocused, setIsStepsFocused] = useState(false);

    const renderIntegralLatexEquation = () => {
        return `\\int_{${lowerBound || -1}}^{${upperBound || 1}} ${toLatex(integralEquation)} \\, dx`;
    };

    return (
        <form onSubmit={(e) => handleSubmit(e, integralMethod, integralEquation, lowerBound, upperBound, steps)}>
            <label>Метод:</label>
            <select value={integralMethod} onChange={(e) => setIntegralMethod(e.target.value)}>
                {integralMethods.map((method) => (
                    <option key={method.value} value={method.value}>{method.label}</option>
                ))}
            </select>
            <label>Подынтегральное уравнение:</label>
            <input
                type="text"
                value={isEquationFocused ? integralEquation : (integralEquation || 'x')}
                onFocus={() => { setIsEquationFocused(true); if (!integralEquation) setIntegralEquation(''); }}
                onBlur={() => setIsEquationFocused(false)}
                onChange={(e) => setIntegralEquation(e.target.value)}
                required
            />
            <label>Нижняя граница:</label>
            <input
                type="text"
                value={isLowerBoundFocused ? lowerBound : (lowerBound || '-1')}
                onFocus={() => { setIsLowerBoundFocused(true); if (!lowerBound) setLowerBound(''); }}
                onBlur={() => setIsLowerBoundFocused(false)}
                onChange={(e) => setLowerBound(e.target.value)}
                required
            />
            <label>Верхняя граница:</label>
            <input
                type="text"
                value={isUpperBoundFocused ? upperBound : (upperBound || '1')}
                onFocus={() => { setIsUpperBoundFocused(true); if (!upperBound) setUpperBound(''); }}
                onBlur={() => setIsUpperBoundFocused(false)}
                onChange={(e) => setUpperBound(e.target.value)}
                required
            />
            <label>Количество шагов:</label>
            <input
                type="text"
                value={isStepsFocused ? steps : (steps || '5')}
                onFocus={() => { setIsStepsFocused(true); if (!steps) setSteps(''); }}
                onBlur={() => setIsStepsFocused(false)}
                onChange={(e) => setSteps(e.target.value)}
                required
            />
            <div className="latex-equation">
                <MathJax>{`$$${renderIntegralLatexEquation()}$$`}</MathJax>
            </div>
            <button type="submit">Рассчитать</button>
            {result !== null && (
                <div className="result">
                    <h3>Результат: {result}</h3>
                </div>
            )}
        </form>
    );
};

export default IntegralForm;
