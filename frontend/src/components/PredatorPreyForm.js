import React, { useState } from 'react';

const PredatorPreyForm = ({ handleSubmit }) => {
    const [preyEquation, setPreyEquation] = useState('');
    const [predatorEquation, setPredatorEquation] = useState('');

    return (
        <form onSubmit={(e) => handleSubmit(e, preyEquation, predatorEquation)}>
            <label>Уравнение для жертвы:</label>
            <input type="text" value={preyEquation} onChange={(e) => setPreyEquation(e.target.value)} required />
            <label>Уравнение для хищника:</label>
            <input type="text" value={predatorEquation} onChange={(e) => setPredatorEquation(e.target.value)} required />
            <button type="submit">Рассчитать</button>
        </form>
    );
};

export default PredatorPreyForm;
