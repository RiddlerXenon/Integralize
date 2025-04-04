import React, { useState } from 'react';
import { MathJax } from 'better-react-mathjax';

const DifferentialForm = ({ differentialMethods, toLatex, handleSubmit, diffData, renderDiffChart }) => {
    const [differentialMethod, setDifferentialMethod] = useState('euler');
    const [differentialEquation, setDifferentialEquation] = useState('');
    const [y0, setY0] = useState('');
    const [x0, setX0] = useState('');
    const [isEquationFocused, setIsEquationFocused] = useState(false);
    const [isY0Focused, setIsY0Focused] = useState(false);
    const [isX0Focused, setIsX0Focused] = useState(false);

    const renderDifferentialLatexEquation = () => {
        return `\\frac{dy}{dx} = ${toLatex(differentialEquation)}, \\ y(${x0 || 0}) = ${y0 || 0}`;
    };

    return (
        <form onSubmit={(e) => handleSubmit(e, differentialMethod, differentialEquation, y0, x0)}>
            <label>Метод:</label>
            <select value={differentialMethod} onChange={(e) => setDifferentialMethod(e.target.value)}>
                {differentialMethods.map((method) => (
                    <option key={method.value} value={method.value}>{method.label}</option>
                ))}
            </select>
            <label>Уравнение:</label>
            <input
                type="text"
                value={isEquationFocused ? differentialEquation : (differentialEquation || 'x')}
                onFocus={() => { setIsEquationFocused(true); if (!differentialEquation) setDifferentialEquation(''); }}
                onBlur={() => setIsEquationFocused(false)}
                onChange={(e) => setDifferentialEquation(e.target.value)}
                required
            />
            <label>y0 - начальное значение:</label>
            <input
                type="text"
                value={isY0Focused ? y0 : (y0 || '0')}
                onFocus={() => { setIsY0Focused(true); if (!y0) setY0(''); }}
                onBlur={() => setIsY0Focused(false)}
                onChange={(e) => setY0(e.target.value)}
                required
            />
            <label>x0 - начальное время:</label>
            <input
                type="text"
                value={isX0Focused ? x0 : (x0 || '0')}
                onFocus={() => { setIsX0Focused(true); if (!x0) setX0(''); }}
                onBlur={() => setIsX0Focused(false)}
                onChange={(e) => setX0(e.target.value)}
                required
            />
            <div className="latex-equation">
                <MathJax>{`$$${renderDifferentialLatexEquation()}$$`}</MathJax>
            </div>
            <button type="submit">Рассчитать</button>
            {diffData && renderDiffChart()}
        </form>
    );
};

export default DifferentialForm;
