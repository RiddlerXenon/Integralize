import React, { useState } from 'react';

const PredatorPreyForm = ({  differentialMethods, handleSubmit, graphData, renderGraph }) => {
    const [differentialMethod, setDifferentialMethod] = useState('euler');
    const [alpha, setAlpha] = useState('');
    const [beta, setBeta] = useState('');
    const [delta, setDelta] = useState('');
    const [gamma, setGamma] = useState('');
    const [step, setStep] = useState('');
    const [steps, setSteps] = useState('');
    const [prey, setPrey] = useState('');
    const [pred, setPred] = useState('');

    const [isAlphaFocused, setIsAlphaFocused] = useState(false);
    const [isBetaFocused, setIsBetaFocused] = useState(false);  
    const [isDeltaFocused, setIsDeltaFocused] = useState(false);
    const [isGammaFocused, setIsGammaFocused] = useState(false);
    const [isStepFocused, setIsStepFocused] = useState(false);
    const [isStepsFocused, setIsStepsFocused] = useState(false);
    const [isPreyFocused, setIsPreyFocused] = useState(false);
    const [isPredFocused, setIsPredFocused] = useState(false);

    return (
        <form onSubmit={(e) => handleSubmit(e, differentialMethod, alpha, beta, delta, gamma, step, steps, prey, pred)}>
            <label>Метод:</label>
            <select value={differentialMethod} onChange={(e) => setDifferentialMethod(e.target.value)}>
                {differentialMethods.map((method) => (
                    <option key={method.value} value={method.value}>{method.label}</option>
                ))}
            </select>
            
            <label>Альфа:</label>
            <input
                type="text"
                value={isAlphaFocused ? alpha : (alpha || '0.1')}
                onChange={(e) => setAlpha(e.target.value)}
                onFocus={() => { setIsAlphaFocused(true); if (!alpha) setAlpha(''); }}
                onBlur={() => setIsAlphaFocused(false)}
                required
            />
            
            <label>Бета:</label>
            <input
                type="text"
                value={isBetaFocused ? beta : (beta || '0.02')}
                onChange={(e) => setBeta(e.target.value)}
                onFocus={() => { setIsBetaFocused(true); if (!beta) setBeta(''); }}
                onBlur={() => setIsBetaFocused(false)}
                required
            />
            <label>Гамма:</label>
            <input
                type="text"
                value={isGammaFocused ? gamma : (gamma || '0.3')}
                onChange={(e) => setGamma(e.target.value)}
                onFocus={() => { setIsGammaFocused(true); if (!gamma) setGamma(''); }}
                onBlur={() => setIsGammaFocused(false)}
                required
            />
            
            <label>Дельта:</label>
            <input 
                type="text"
                value={isDeltaFocused ? delta : (delta || '0.01')}
                onChange={(e) => setDelta(e.target.value)}
                onFocus={() => { setIsDeltaFocused(true); if (!delta) setDelta(''); }}
                onBlur={() => setIsDeltaFocused(false)}
                required
            />
            
            <label>Шаг:</label>
            <input
                type="text"
                value={isStepFocused ? step : (step || '0.1')}
                onChange={(e) => setStep(e.target.value)}
                onFocus={() => { setIsStepFocused(true); if (!step) setStep(''); }}
                onBlur={() => setIsStepFocused(false)}
                required
            />
            <label>Количество шагов:</label>
            <input
                type="text"
                value={isStepsFocused ? steps : (steps || '1000')}
                onChange={(e) => setSteps(e.target.value)}
                onFocus={() => { setIsStepsFocused(true); if (!steps) setSteps(''); }}
                onBlur={() => setIsStepsFocused(false)}
                required
            />
            <label>Начальное количество жертв:</label>
            <input
                type="text"
                value={isPreyFocused ? prey : (prey || '40')}
                onChange={(e) => setPrey(e.target.value)}
                onFocus={() => { setIsPreyFocused(true); if (!prey) setPrey(''); }}
                onBlur={() => setIsPreyFocused(false)}
                required
            />
            <label>Начальное количество хищников:</label>
            <input
                type="text"
                value={isPredFocused ? pred : (pred || '9')}
                onChange={(e) => setPred(e.target.value)}
                onFocus={() => { setIsPredFocused(true); if (!pred) setPred(''); }}
                onBlur={() => setIsPredFocused(false)}
                required
            />
            <button type="submit">Рассчитать</button>
            {graphData && renderGraph()}
        </form>
    );
};

export default PredatorPreyForm;
